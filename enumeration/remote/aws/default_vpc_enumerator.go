package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type DefaultVPCEnumerator struct {
	repo    repository.EC2Repository
	factory resource.ResourceFactory
}

func NewDefaultVPCEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *DefaultVPCEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *DefaultVPCEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *DefaultVPCEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
