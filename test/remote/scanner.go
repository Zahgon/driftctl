package remote

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type SortableScanner struct {
	Scanner resource.Supplier
}

func NewSortableScanner(scanner resource.Supplier) *SortableScanner {
	_ = "STUB: not implemented"
	return nil
}

func (s *SortableScanner) Resources() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
