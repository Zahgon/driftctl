package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2RouteTableAssociationEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2RouteTableAssociationEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2RouteTableAssociationEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2RouteTableAssociationEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2RouteTableAssociationEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EC2RouteTableAssociationEnumerator) shouldBeIgnored(assoc *ec2.RouteTableAssociation) bool {
	_ = "STUB: not implemented"
	// Ignore when nothing is associated
	return false
}

// Ignore when association is not associated
