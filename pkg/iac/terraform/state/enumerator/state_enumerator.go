package enumerator

import (
	"github.com/snyk/driftctl/pkg/iac/config"
	"github.com/snyk/driftctl/pkg/iac/terraform/state/backend"
)

type StateEnumerator interface {
	Origin() string
	Enumerate() ([]string, error)
}

func GetEnumerator(config config.SupplierConfig, opts *backend.Options) (StateEnumerator, error) {
	_ = "STUB: not implemented"
	return *new(StateEnumerator), nil
}
