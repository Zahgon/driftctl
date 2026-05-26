package terraform

import (
	"github.com/snyk/driftctl/enumeration/remote/aws"
	"github.com/snyk/driftctl/enumeration/remote/azurerm"
	"github.com/snyk/driftctl/enumeration/remote/github"
	"github.com/snyk/driftctl/enumeration/remote/google"
	"github.com/snyk/driftctl/enumeration/terraform"
)

func InitTestAwsProvider(providerLibrary *terraform.ProviderLibrary, version string) (*aws.AWSTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InitTestGithubProvider(providerLibrary *terraform.ProviderLibrary, version string) (*github.GithubTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InitTestGoogleProvider(providerLibrary *terraform.ProviderLibrary, version string) (*google.GCPTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InitTestAzureProvider(providerLibrary *terraform.ProviderLibrary, version string) (*azurerm.AzureTerraformProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
