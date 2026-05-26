package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermNetworkSecurityGroupEnumerator struct {
	repository repository.NetworkRepository
	factory    resource.ResourceFactory
}

func NewAzurermNetworkSecurityGroupEnumerator(repo repository.NetworkRepository, factory resource.ResourceFactory) *AzurermNetworkSecurityGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermNetworkSecurityGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermNetworkSecurityGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
