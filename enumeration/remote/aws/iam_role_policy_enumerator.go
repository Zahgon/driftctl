package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type IamRolePolicyEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamRolePolicyEnumerator(repository repository.IAMRepository, factory resource.ResourceFactory) *IamRolePolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamRolePolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamRolePolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
