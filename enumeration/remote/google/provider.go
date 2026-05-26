package google

import (
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/remote/google/config"
	"github.com/snyk/driftctl/enumeration/remote/terraform"
)

type GCPTerraformProvider struct {
	*terraform.TerraformProvider
	name    string
	version string
}

func NewGCPTerraformProvider(version string, progress enumeration.ProgressCounter, configDir string) (*GCPTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *GCPTerraformProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *GCPTerraformProvider) Version() string { _ = "STUB: not implemented"; return "" }

func (p *GCPTerraformProvider) GetConfig() config.GCPTerraformConfig {
	_ = "STUB: not implemented"
	return *new(config.GCPTerraformConfig)
}

func (p *GCPTerraformProvider) CheckCredentialsExist() error { _ = "STUB: not implemented"; return nil }
