package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermPrivateDNSSRVRecordEnumerator struct {
	repository repository.PrivateDNSRepository
	factory    resource.ResourceFactory
}

func NewAzurermPrivateDNSSRVRecordEnumerator(repo repository.PrivateDNSRepository, factory resource.ResourceFactory) *AzurermPrivateDNSSRVRecordEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermPrivateDNSSRVRecordEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermPrivateDNSSRVRecordEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
