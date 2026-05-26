package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type Route53ZoneSupplier struct {
	client  repository.Route53Repository
	factory resource.ResourceFactory
}

func NewRoute53ZoneEnumerator(repo repository.Route53Repository, factory resource.ResourceFactory) *Route53ZoneSupplier {
	_ = "STUB: not implemented"
	return nil
}

func (e *Route53ZoneSupplier) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *Route53ZoneSupplier) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
