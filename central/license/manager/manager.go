package manager

import (
	"github.com/stackrox/rox/central/deploymentenvs"
	"github.com/stackrox/rox/central/license/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	licenseproto "github.com/stackrox/rox/generated/shared/license"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/license/validator"
)

// LicenseManager is responsible for managing product licenses.
type LicenseManager interface {

	// Initialize starts the license manager and returns the active license, if any. The listener is registered
	// synchronously and will deliver any license event *after* the selection of an initially active license.
	Initialize() (*licenseproto.License, error)
	Stop() concurrency.Waitable

	GetActiveLicenseKey() string
	GetActiveLicense() *licenseproto.License
	GetAllLicenses() []*v1.LicenseInfo

	AddLicenseKey(licenseKey string, activate bool) (*v1.LicenseInfo, error)
	SelectLicense(licenseID string) (*v1.LicenseInfo, error)

	GetLicenseStatus() v1.Metadata_LicenseStatus

	SignWithLicenseKeyHash(licenseID string, payload []byte) ([]byte, error)
}

// New creates and returns a new license manager, using the given license key store and validator.
func New(dataStore datastore.DataStore, validator validator.Validator, deploymentEnvsMgr deploymentenvs.Manager) LicenseManager {
	return newManager(dataStore, validator, deploymentEnvsMgr)
}
