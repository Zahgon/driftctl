package enumerator

import (
	"cloud.google.com/go/storage"

	"github.com/snyk/driftctl/pkg/iac/config"
)

type GSEnumerator struct {
	config config.SupplierConfig
	client storage.Client
}

func NewGSEnumerator(config config.SupplierConfig) (*GSEnumerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GSEnumerator) Origin() string { _ = "STUB: not implemented"; return "" }

func (s *GSEnumerator) Enumerate() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// prefix should contains everything that does not have a glob pattern
// Pattern should be the glob matcher string

// We combine the prefix and pattern to match file names against.
