package localscanner

import (
	"crypto/x509"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestHandler(t *testing.T) {
	suite.Run(t, new(operatorSuite))
}

type operatorSuite struct {
	suite.Suite
}

func (s *operatorSuite) TestGetScannerSecretDurationFromCertificate() {
	now := time.Now()
	afterOffset := 2 * 24 * time.Hour
	scannerCert := &x509.Certificate{
		NotBefore: now,
		NotAfter:  now.Add(afterOffset),
	}
	certDuration := getScannerSecretDurationFromCertificate(scannerCert)
	s.Assert().LessOrEqual(certDuration, afterOffset/2)
}