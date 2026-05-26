package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type VPCEnumerator struct {
	repo    repository.EC2Repository
	factory resource.ResourceFactory
}

func NewVPCEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *VPCEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *VPCEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *VPCEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
