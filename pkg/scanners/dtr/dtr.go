package dtr

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/scanners/types"
	"github.com/stackrox/rox/pkg/urlfmt"
)

const (
	requestTimeout = 30 * time.Second
)

var (
	log = logging.LoggerForModule()
)

// Creator provides the type an scanners.Creator to add to the scanners Registry.
func Creator() (string, func(integration *v1.ImageIntegration) (types.ImageScanner, error)) {
	return "dtr", func(integration *v1.ImageIntegration) (types.ImageScanner, error) {
		scan, err := newScanner(integration)
		return scan, err
	}
}

type dtr struct {
	client *http.Client

	conf     config
	registry string

	protoImageIntegration *v1.ImageIntegration
}

type config v1.DTRConfig

func (c config) validate() error {
	errorList := errorhelpers.NewErrorList("Validation")
	if c.Username == "" {
		errorList.AddString("username parameter must be defined for DTR")
	}
	if c.Password == "" {
		errorList.AddString("password parameter must be defined for DTR")
	}
	if c.Endpoint == "" {
		errorList.AddString("endpoint parameter must be defined for DTR")
	}
	return errorList.ToError()
}

func newScanner(protoImageIntegration *v1.ImageIntegration) (*dtr, error) {
	dtrConfig, ok := protoImageIntegration.IntegrationConfig.(*v1.ImageIntegration_Dtr)
	if !ok {
		return nil, fmt.Errorf("DTR configuration required")
	}
	conf := config(*dtrConfig.Dtr)
	if err := conf.validate(); err != nil {
		return nil, err
	}

	// Trim any trailing slashes as the expectation will be that the input is in the form
	// https://12.12.12.12:8080 or https://dtr.com
	var err error
	conf.Endpoint, err = urlfmt.FormatURL(conf.Endpoint, urlfmt.HTTPS, urlfmt.NoTrailingSlash)
	if err != nil {
		return nil, err
	}
	registry := urlfmt.GetServerFromURL(conf.Endpoint)
	client := &http.Client{
		Timeout: requestTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: conf.Insecure,
			},
		},
	}

	scanner := &dtr{
		client:                client,
		registry:              registry,
		conf:                  conf,
		protoImageIntegration: protoImageIntegration,
	}

	return scanner, nil
}

func (d *dtr) sendRequest(client *http.Client, method, urlPrefix string) ([]byte, error) {
	req, err := http.NewRequest(method, d.conf.Endpoint+urlPrefix, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(d.conf.Username, d.conf.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := errorFromStatusCode(resp.StatusCode); err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

// GetScan takes in an id and returns the image scan for that id if applicable
func (d *dtr) GetScans(image *v1.Image) ([]*v1.ImageScan, error) {
	if image == nil || image.GetName().GetRemote() == "" || image.GetName().GetTag() == "" {
		return nil, nil
	}
	getScanURL := fmt.Sprintf("/api/v0/imagescan/repositories/%v/%v?detailed=true", image.GetName().GetRemote(), image.GetName().GetTag())
	body, err := d.sendRequest(d.client, "GET", getScanURL)
	if err != nil {
		return nil, err
	}
	scans, err := parseDTRImageScans(body)
	if err != nil {
		scanErrors, err := parseDTRImageScanErrors(body)
		if err != nil {
			return nil, err
		}
		var errMsg string
		for _, scanErr := range scanErrors.Errors {
			errMsg += scanErr.Message + "\n"
		}
		return nil, errors.New(errMsg)
	}
	if len(scans) == 0 {
		return nil, fmt.Errorf("expected to receive at least one scan for %v", image.String())
	}
	// After should sort in descending order based on completion
	sort.SliceStable(scans, func(i, j int) bool { return scans[i].CheckCompletedAt.After(scans[j].CheckCompletedAt) })
	return convertTagScanSummariesToImageScans(d.conf.Endpoint, scans), nil
}

//GET /api/v0/imagescan/repositories/{namespace}/{reponame}/{tag}?detailed=true
// Scan initiates a scan of the passed id
func (d *dtr) Scan(image *v1.Image) error {
	_, err := d.sendRequest(d.client, "POST", fmt.Sprintf("/api/v0/imagescan/scan/%v/%v/linux/amd64", image.GetName().GetRemote(), image.GetName().GetTag()))
	return err
}

func errorFromStatusCode(status int) error {
	switch status {
	case 400:
		return fmt.Errorf("HTTP 400: Scanning is not enabled")
	case 401:
		return fmt.Errorf("HTTP 401: The client is not authenticated")
	case 405:
		return fmt.Errorf("HTTP 405: Method Not Allowed")
	case 406:
		return fmt.Errorf("HTTP 406: Not Acceptable")
	case 415:
		return fmt.Errorf("HTTP 415: Unsupported Media Type")
	case 200:
	default:
		return nil
	}
	return nil
}

// Test initiates a test of the DTR which verifies that we have the proper scan permissions
func (d *dtr) Test() error {
	client := &http.Client{
		Timeout: requestTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: d.conf.Insecure,
			},
		},
	}
	_, err := d.sendRequest(client, "GET", "/api/v0/imagescan/status")
	return err
}

// GetLastScan retrieves the most recent scan
func (d *dtr) GetLastScan(image *v1.Image) (*v1.ImageScan, error) {
	log.Infof("Getting latest scan for image %v", image.GetName().GetFullName())
	imageScans, err := d.GetScans(image)
	if err != nil {
		return nil, err
	}
	if len(imageScans) == 0 {
		return nil, fmt.Errorf("no scans were found for image %v", image.GetName().GetFullName())
	}
	return imageScans[0], nil
}

// Match decides if the image is contained within this registry
func (d *dtr) Match(image *v1.Image) bool {
	return d.registry == image.GetName().GetRegistry()
}

func (d *dtr) Global() bool {
	return len(d.protoImageIntegration.GetClusters()) == 0
}
