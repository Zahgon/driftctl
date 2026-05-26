package azurerm

import (
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/remote/azurerm/common"
	"github.com/snyk/driftctl/enumeration/remote/terraform"
)

type AzureTerraformProvider struct {
	*terraform.TerraformProvider
	name    string
	version string
}

func NewAzureTerraformProvider(version string, progress enumeration.ProgressCounter, configDir string) (*AzureTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Just pass your version and name

// Use TerraformProviderInstaller to retrieve the provider if needed

func (p *AzureTerraformProvider) GetConfig() common.AzureProviderConfig {
	_ = "STUB: not implemented"
	return *new(common.AzureProviderConfig)
}

func (p *AzureTerraformProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *AzureTerraformProvider) Version() string { _ = "STUB: not implemented"; return "" }

func (p *AzureTerraformProvider) CheckCredentialsExist() error {
	_ = "STUB: not implemented"
	return nil
}
