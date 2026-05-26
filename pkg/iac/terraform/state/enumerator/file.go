package enumerator

import (
	"github.com/snyk/driftctl/pkg/iac/config"
)

type FileEnumeratorConfig struct {
	Bucket *string
	Prefix *string
}

type FileEnumerator struct {
	config config.SupplierConfig
}

func NewFileEnumerator(config config.SupplierConfig) *FileEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (s *FileEnumerator) Origin() string { _ = "STUB: not implemented"; return "" }

func (s *FileEnumerator) Enumerate() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// if we got a symlink, use its destination
