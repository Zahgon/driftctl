package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type LoadBalancerEnumerator struct {
	repository repository.ELBV2Repository
	factory    resource.ResourceFactory
}

func NewLoadBalancerEnumerator(repo repository.ELBV2Repository, factory resource.ResourceFactory) *LoadBalancerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *LoadBalancerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *LoadBalancerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
