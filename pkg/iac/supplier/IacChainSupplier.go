package supplier

import (
	"github.com/snyk/driftctl/enumeration/parallel"
	resource2 "github.com/snyk/driftctl/pkg/resource"

	"github.com/snyk/driftctl/enumeration/resource"
)

type IacChainSupplier struct {
	suppliers []resource2.IaCSupplier
	runner    *parallel.ParallelRunner
}

func NewIacChainSupplier() *IacChainSupplier { _ = "STUB: not implemented"; return nil }

func (r *IacChainSupplier) SourceCount() uint { _ = "STUB: not implemented"; return 0 }

func (r *IacChainSupplier) AddSupplier(supplier resource2.IaCSupplier) {
	_ = "STUB: not implemented"
	return
}

func (r *IacChainSupplier) Resources() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Type cannot be invalid as return type is enforced
// in run function on top

// only fail if all suppliers failed

type result struct {
	err error
	res []*resource.Resource
}
