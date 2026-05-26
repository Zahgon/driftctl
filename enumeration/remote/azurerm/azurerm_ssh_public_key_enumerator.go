package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermSSHPublicKeyEnumerator struct {
	repository repository.ComputeRepository
	factory    resource.ResourceFactory
}

func NewAzurermSSHPublicKeyEnumerator(repo repository.ComputeRepository, factory resource.ResourceFactory) *AzurermSSHPublicKeyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermSSHPublicKeyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermSSHPublicKeyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
