package openshift

import (
	"context"
	"crypto/x509"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dexidp/dex/connector"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/auth/authproviders"
	"github.com/stackrox/rox/pkg/auth/authproviders/idputil"
	"github.com/stackrox/rox/pkg/auth/authproviders/openshift/internal/dexconnector"
	"github.com/stackrox/rox/pkg/auth/tokens"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/grpc/requestinfo"
	"github.com/stackrox/rox/pkg/httputil"
	"github.com/stackrox/rox/pkg/netutil"
	"github.com/stackrox/rox/pkg/satoken"
)

const (
	openshiftAPIUrl    = "https://openshift.default.svc"
	roxTokenExpiration = 5 * time.Minute
)

// This is the location for CA files which shall be used for certificate validation within
// openshift auth. In addition to the CA files here, the system's trusted root CAs will be used as well.
// The path may or may not exist depending on cluster state & configuration.
const (
	// serviceAccountCACertPath points to the secret of the service account, which within an OpenShift environment
	// also has the service-ca.crt which includes CA's for internal services and Ingress Controller certificates.
	serviceAccountCACertPath = "/run/secrets/kubernetes.io/serviceaccount/service-ca.crt"
)

var (
	defaultScopes = connector.Scopes{
		OfflineAccess: true,
		Groups:        true,
	}
)

type callbackAndRefreshConnector interface {
	connector.CallbackConnector
	connector.RefreshConnector
}

type backend struct {
	id                 string
	baseRedirectURL    url.URL
	openshiftConnector callbackAndRefreshConnector
}

type openShiftSettings struct {
	clientID        string
	clientSecret    string
	trustedCertPool *x509.CertPool
}

var _ authproviders.RefreshTokenEnabledBackend = (*backend)(nil)

func newBackend(id string, callbackURLPath string, _ map[string]string) (authproviders.Backend, error) {
	settings, err := getOpenShiftSettings()
	if err != nil {
		return nil, err
	}

	baseRedirectURL := url.URL{
		Scheme: "https",
		Path:   callbackURLPath,
	}

	dexCfg := dexconnector.Config{
		Issuer:          openshiftAPIUrl,
		ClientID:        settings.clientID,
		ClientSecret:    settings.clientSecret,
		TrustedCertPool: settings.trustedCertPool,
	}

	openshiftConnector, err := dexCfg.Open()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create dex openshiftConnector for OpenShift's OAuth Server")
	}

	b := &backend{
		id:                 id,
		baseRedirectURL:    baseRedirectURL,
		openshiftConnector: openshiftConnector,
	}

	return b, nil
}

// There is no config but static settings instead.
func (b *backend) Config() map[string]string {
	return nil
}

func (b *backend) LoginURL(clientState string, ri *requestinfo.RequestInfo) string {
	state := idputil.MakeState(b.id, clientState)

	// baseRedirectURL does not include the hostname, take it from the request.
	// Allow HTTP only if the client did not use TLS and the host is localhost.
	redirectURL := b.baseRedirectURL
	redirectURL.Host = ri.Hostname
	if !ri.ClientUsedTLS && netutil.IsLocalEndpoint(redirectURL.Host) {
		redirectURL.Scheme = "http"
	}

	loginURL, _ := b.openshiftConnector.LoginURL(defaultScopes, redirectURL.String(), state)
	return loginURL
}

func (b *backend) RefreshURL() string {
	return ""
}

func (b *backend) OnEnable(_ authproviders.Provider) {}

func (b *backend) OnDisable(_ authproviders.Provider) {}

func (b *backend) ProcessHTTPRequest(_ http.ResponseWriter, r *http.Request) (*authproviders.AuthResponse, error) {
	if r.URL.Path != b.baseRedirectURL.Path {
		return nil, httputil.Errorf(http.StatusNotFound, "path %q not found", r.URL.Path)
	}
	if r.Method != http.MethodGet {
		return nil, httputil.Errorf(http.StatusMethodNotAllowed, "unsupported method %q, only GET requests are allowed to this URL", r.Method)
	}
	id, err := b.openshiftConnector.HandleCallback(defaultScopes, r)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving user identity")
	}
	return b.idToAuthResponse(&id), nil
}

func (b *backend) idToAuthResponse(id *connector.Identity) *authproviders.AuthResponse {
	attributes := map[string][]string{
		"userid": {id.UserID},
		"name":   {id.Username},
		"groups": id.Groups,
	}
	if id.Email != "" {
		attributes["email"] = []string{id.Email}
	}

	return &authproviders.AuthResponse{
		Claims: &tokens.ExternalUserClaim{
			UserID:     id.Username,
			FullName:   id.Username,
			Email:      id.Email,
			Attributes: attributes,
		},
		Expiration: time.Now().Add(roxTokenExpiration),
		RefreshTokenData: authproviders.RefreshTokenData{
			RefreshToken: string(id.ConnectorData),
		},
	}
}

// RefreshAccessToken attempts to fetch user info and issue an updated auth
// status. If the refresh token has expired, error is returned.
func (b *backend) RefreshAccessToken(ctx context.Context, refreshTokenData authproviders.RefreshTokenData) (*authproviders.AuthResponse, error) {
	id, err := b.openshiftConnector.Refresh(ctx, defaultScopes, connector.Identity{
		ConnectorData: []byte(refreshTokenData.RefreshToken),
	})
	if err != nil {
		return nil, errors.Wrap(err, "retrieving user identity")
	}
	return b.idToAuthResponse(&id), nil
}

func (b *backend) RevokeRefreshToken(_ context.Context, _ authproviders.RefreshTokenData) error {
	return nil
}

func (b *backend) ExchangeToken(_ context.Context, _ string, _ string) (*authproviders.AuthResponse, string, error) {
	return nil, "", errors.New("not implemented")
}

func (b *backend) Validate(_ context.Context, _ *tokens.Claims) error {
	return nil
}

func getOpenShiftSettings() (openShiftSettings, error) {
	clientID := "system:serviceaccount:" + env.Namespace.Setting() + ":central"

	clientSecret, err := satoken.LoadTokenFromFile()
	if err != nil {
		return openShiftSettings{}, errors.Wrap(err, "reading service account token")
	}

	certPool, err := getSystemCertPoolWithAdditionalCA(serviceAccountCACertPath)
	if err != nil {
		return openShiftSettings{}, err
	}

	return openShiftSettings{
		clientID:        clientID,
		clientSecret:    clientSecret,
		trustedCertPool: certPool,
	}, nil
}

func getSystemCertPoolWithAdditionalCA(additionalCAPath string) (*x509.CertPool, error) {
	// Use the x509.SystemCertPool to include system's trusted CAs.
	sysCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "creating system cert pool")
	}

	rootCABytes, err := os.ReadFile(additionalCAPath)
	if errors.Is(err, os.ErrNotExist) {
		return sysCertPool, nil
	}

	if err != nil {
		return nil, errors.Wrapf(err, "reading CA at path %s", additionalCAPath)
	}

	if !sysCertPool.AppendCertsFromPEM(rootCABytes) {
		return nil, errors.Errorf("parsing root CA file from %s", additionalCAPath)
	}

	return sysCertPool, nil
}