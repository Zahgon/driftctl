package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type KMSKeyEnumerator struct {
	repository repository.KMSRepository
	factory    resource.ResourceFactory
}

func NewKMSKeyEnumerator(repo repository.KMSRepository, factory resource.ResourceFactory) *KMSKeyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *KMSKeyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *KMSKeyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
