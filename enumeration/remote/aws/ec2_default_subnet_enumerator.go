package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2DefaultSubnetEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2DefaultSubnetEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2DefaultSubnetEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2DefaultSubnetEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2DefaultSubnetEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
