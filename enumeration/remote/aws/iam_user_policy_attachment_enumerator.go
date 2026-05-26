package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type IamUserPolicyAttachmentEnumerator struct {
	repository repository.IAMRepository
	factory    resource.ResourceFactory
}

func NewIamUserPolicyAttachmentEnumerator(repository repository.IAMRepository, factory resource.ResourceFactory) *IamUserPolicyAttachmentEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *IamUserPolicyAttachmentEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *IamUserPolicyAttachmentEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
