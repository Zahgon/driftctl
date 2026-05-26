package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermImageEnumerator struct {
	repository repository.ComputeRepository
	factory    resource.ResourceFactory
}

func NewAzurermImageEnumerator(repo repository.ComputeRepository, factory resource.ResourceFactory) *AzurermImageEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermImageEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermImageEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Here we turn the resource group into lowercase because for some reason the API returns it in uppercase.
