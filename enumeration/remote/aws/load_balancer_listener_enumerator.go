package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type LoadBalancerListenerEnumerator struct {
	repository repository.ELBV2Repository
	factory    resource.ResourceFactory
}

func NewLoadBalancerListenerEnumerator(repo repository.ELBV2Repository, factory resource.ResourceFactory) *LoadBalancerListenerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *LoadBalancerListenerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *LoadBalancerListenerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
