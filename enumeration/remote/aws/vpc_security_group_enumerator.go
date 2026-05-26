package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type VPCSecurityGroupEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewVPCSecurityGroupEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *VPCSecurityGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *VPCSecurityGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *VPCSecurityGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
