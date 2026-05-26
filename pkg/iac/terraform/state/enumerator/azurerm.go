package enumerator

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend/options"
)

type AzureRMEnumerator struct {
	containerName, objectPath string
	containerClient           azblob.ContainerClient
	origin                    string
}

func NewAzureRMEnumerator(config config.SupplierConfig, opts options.AzureRMBackendOptions) (*AzureRMEnumerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *AzureRMEnumerator) Origin() string { _ = "STUB: not implemented"; return "" }

func (s *AzureRMEnumerator) Enumerate() ([]string, error) {
	_ = "STUB: not implemented"
	// prefix should contains everything that does not have a glob pattern
	// Pattern should be the glob matcher string
	return nil, nil
}

// We combine the prefix and pattern to match file names against.
