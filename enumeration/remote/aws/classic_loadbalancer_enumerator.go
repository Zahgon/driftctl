package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ClassicLoadBalancerEnumerator struct {
	repository repository.ELBRepository
	factory    resource.ResourceFactory
}

func NewClassicLoadBalancerEnumerator(repo repository.ELBRepository, factory resource.ResourceFactory) *ClassicLoadBalancerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClassicLoadBalancerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ClassicLoadBalancerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
