package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermVirtualNetworkEnumerator struct {
	repository repository.NetworkRepository
	factory    resource.ResourceFactory
}

func NewAzurermVirtualNetworkEnumerator(repo repository.NetworkRepository, factory resource.ResourceFactory) *AzurermVirtualNetworkEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermVirtualNetworkEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermVirtualNetworkEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
