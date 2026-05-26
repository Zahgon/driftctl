package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type IamRolePolicyAttachmentEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamRolePolicyAttachmentEnumerator(repository repository.IAMRepository, factory resource.ResourceFactory) *IamRolePolicyAttachmentEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamRolePolicyAttachmentEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamRolePolicyAttachmentEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
