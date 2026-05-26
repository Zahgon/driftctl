package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermLoadBalancerEnumerator struct {
	repository repository.NetworkRepository
	factory    resource.ResourceFactory
}

func NewAzurermLoadBalancerEnumerator(repo repository.NetworkRepository, factory resource.ResourceFactory) *AzurermLoadBalancerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermLoadBalancerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermLoadBalancerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
