package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermResourceGroupEnumerator struct {
	repository repository.ResourcesRepository
	factory    resource.ResourceFactory
}

func NewAzurermResourceGroupEnumerator(repo repository.ResourcesRepository, factory resource.ResourceFactory) *AzurermResourceGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermResourceGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermResourceGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
