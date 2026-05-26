package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type IamUserEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamUserEnumerator(repo repository.IAMRepository, factory resource.ResourceFactory) *IamUserEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamUserEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamUserEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
