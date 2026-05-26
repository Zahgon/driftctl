package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2InternetGatewayEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2InternetGatewayEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2InternetGatewayEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2InternetGatewayEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2InternetGatewayEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
