package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type IamGroupEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamGroupEnumerator(repo repository.IAMRepository, factory resource.ResourceFactory) *IamGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
