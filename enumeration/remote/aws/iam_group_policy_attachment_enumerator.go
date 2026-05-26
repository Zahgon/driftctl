package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type IamGroupPolicyAttachmentEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamGroupPolicyAttachmentEnumerator(repository repository.IAMRepository, factory resource.ResourceFactory) *IamGroupPolicyAttachmentEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamGroupPolicyAttachmentEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamGroupPolicyAttachmentEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
