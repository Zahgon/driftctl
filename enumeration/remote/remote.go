package remote

import (
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/remote/common"
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/enumeration/terraform"
)

var supportedRemotes = []string{
	common.RemoteAWSTerraform,
	common.RemoteGithubTerraform,
	common.RemoteGoogleTerraform,
	common.RemoteAzureTerraform,
}

func IsSupported(remote string) bool { _ = "STUB: not implemented"; return false }

func Activate(remote, version string, alerter alerter.AlerterInterface, providerLibrary *terraform.ProviderLibrary, remoteLibrary *common.RemoteLibrary, progress enumeration.ProgressCounter, factory resource.ResourceFactory, configDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSupportedRemotes() []string { _ = "STUB: not implemented"; return nil }
