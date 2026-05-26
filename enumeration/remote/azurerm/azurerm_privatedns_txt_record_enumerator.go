package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermPrivateDNSTXTRecordEnumerator struct {
	repository repository.PrivateDNSRepository
	factory    resource.ResourceFactory
}

func NewAzurermPrivateDNSTXTRecordEnumerator(repo repository.PrivateDNSRepository, factory resource.ResourceFactory) *AzurermPrivateDNSTXTRecordEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermPrivateDNSTXTRecordEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermPrivateDNSTXTRecordEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
