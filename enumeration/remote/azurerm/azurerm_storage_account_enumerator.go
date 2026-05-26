package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermStorageAccountEnumerator struct {
	repository repository.StorageRespository
	factory    resource.ResourceFactory
}

func NewAzurermStorageAccountEnumerator(repo repository.StorageRespository, factory resource.ResourceFactory) *AzurermStorageAccountEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermStorageAccountEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermStorageAccountEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
