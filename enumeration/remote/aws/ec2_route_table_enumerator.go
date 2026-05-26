package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2RouteTableEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2RouteTableEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2RouteTableEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2RouteTableEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2RouteTableEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isMainRouteTable(routeTable *ec2.RouteTable) bool { _ = "STUB: not implemented"; return false }
