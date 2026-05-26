package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermContainerRegistryEnumerator struct {
	repository repository.ContainerRegistryRepository
	factory    resource.ResourceFactory
}

func NewAzurermContainerRegistryEnumerator(repo repository.ContainerRegistryRepository, factory resource.ResourceFactory) *AzurermContainerRegistryEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermContainerRegistryEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermContainerRegistryEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
