package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type Route53HealthCheckEnumerator struct {
	repository repository.Route53Repository
	factory    resource.ResourceFactory
}

func NewRoute53HealthCheckEnumerator(repo repository.Route53Repository, factory resource.ResourceFactory) *Route53HealthCheckEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *Route53HealthCheckEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *Route53HealthCheckEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
