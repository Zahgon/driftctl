package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2DefaultNetworkACLEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2DefaultNetworkACLEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2DefaultNetworkACLEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2DefaultNetworkACLEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2DefaultNetworkACLEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not handle non-default network acl since it is a dedicated resource
