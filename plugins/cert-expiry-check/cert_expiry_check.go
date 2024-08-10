package cert_expiry_check

import "github.com/StandardRunbook/hypothecary/plugins"

type CertExpiryCheck struct {
}

func NewCertExpiryCheck() plugins.IPlugin {
	return &CertExpiryCheck{}
}

func (c *CertExpiryCheck) Name() string {
	//TODO implement me
	panic("implement me")
}

func (c *CertExpiryCheck) Run() string {
	//TODO implement me
	panic("implement me")
}

func (c *CertExpiryCheck) ParseOutput() string {
	//TODO implement me
	panic("implement me")
}
