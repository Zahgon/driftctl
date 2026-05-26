package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2NetworkACLEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2NetworkACLEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2NetworkACLEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2NetworkACLEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2NetworkACLEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not handle default network acl since it is a dedicated resource
