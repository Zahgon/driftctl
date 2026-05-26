package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type IamGroupPolicyEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamGroupPolicyEnumerator(repo repository.IAMRepository, factory resource.ResourceFactory) *IamGroupPolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamGroupPolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamGroupPolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
