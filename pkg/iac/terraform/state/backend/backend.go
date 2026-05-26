package backend

import (
	"io"

	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend/options"
)

var supportedBackends = []string{
	BackendKeyFile,
	BackendKeyS3,
	BackendKeyHTTP,
	BackendKeyHTTPS,
	BackendKeyTFCloud,
	BackendKeyGS,
	BackendKeyAzureRM,
}

type Backend io.ReadCloser

type Options struct {
	Headers         map[string]string
	TFCloudToken    string
	TFCloudEndpoint string
	options.AzureRMBackendOptions
}

func IsSupported(backend string) bool { _ = "STUB: not implemented"; return false }

func GetBackend(config config.SupplierConfig, opts *Options) (Backend, error) {
	_ = "STUB: not implemented"
	return *new(Backend), nil
}

func GetSupportedBackends() []string { _ = "STUB: not implemented"; return nil }
