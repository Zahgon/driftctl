package backend

import (
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend/options"
)

const BackendKeyAzureRM = "azurerm"

type AzureRMBackend struct {
	reader        io.ReadCloser
	storageClient azblob.BlockBlobClient
}

func NewAzureRMReader(path string, opts options.AzureRMBackendOptions) (*AzureRMBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *AzureRMBackend) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *AzureRMBackend) Close() error { _ = "STUB: not implemented"; return nil }
