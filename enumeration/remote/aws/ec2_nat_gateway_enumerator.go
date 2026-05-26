package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2NatGatewayEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2NatGatewayEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2NatGatewayEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2NatGatewayEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2NatGatewayEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
