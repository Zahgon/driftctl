package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermRouteEnumerator struct {
	repository repository.NetworkRepository
	factory    resource.ResourceFactory
}

func NewAzurermRouteEnumerator(repo repository.NetworkRepository, factory resource.ResourceFactory) *AzurermRouteEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermRouteEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermRouteEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
