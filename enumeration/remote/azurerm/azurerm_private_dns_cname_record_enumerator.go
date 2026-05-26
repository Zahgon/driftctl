package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermPrivateDNSCNameRecordEnumerator struct {
	repository repository.PrivateDNSRepository
	factory    resource.ResourceFactory
}

func NewAzurermPrivateDNSCNameRecordEnumerator(repo repository.PrivateDNSRepository, factory resource.ResourceFactory) *AzurermPrivateDNSCNameRecordEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermPrivateDNSCNameRecordEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermPrivateDNSCNameRecordEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
