package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type VPCDefaultSecurityGroupEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewVPCDefaultSecurityGroupEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *VPCDefaultSecurityGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *VPCDefaultSecurityGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *VPCDefaultSecurityGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
