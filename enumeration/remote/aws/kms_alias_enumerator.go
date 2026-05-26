package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type KMSAliasEnumerator struct {
	repository repository.KMSRepository
	factory    resource.ResourceFactory
}

func NewKMSAliasEnumerator(repo repository.KMSRepository, factory resource.ResourceFactory) *KMSAliasEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *KMSAliasEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *KMSAliasEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
