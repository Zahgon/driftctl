package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermStorageContainerEnumerator struct {
	repository repository.StorageRespository
	factory    resource.ResourceFactory
}

func NewAzurermStorageContainerEnumerator(repo repository.StorageRespository, factory resource.ResourceFactory) *AzurermStorageContainerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermStorageContainerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermStorageContainerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
