package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type IamAccessKeyEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamAccessKeyEnumerator(repository repository.IAMRepository, factory resource.ResourceFactory) *IamAccessKeyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamAccessKeyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamAccessKeyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
