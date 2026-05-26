package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermLoadBalancerRuleEnumerator struct {
	repository repository.NetworkRepository
	factory    resource.ResourceFactory
}

func NewAzurermLoadBalancerRuleEnumerator(repo repository.NetworkRepository, factory resource.ResourceFactory) *AzurermLoadBalancerRuleEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermLoadBalancerRuleEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermLoadBalancerRuleEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
