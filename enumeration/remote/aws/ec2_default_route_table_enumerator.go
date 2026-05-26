package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2DefaultRouteTableEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2DefaultRouteTableEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2DefaultRouteTableEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2DefaultRouteTableEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2DefaultRouteTableEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
