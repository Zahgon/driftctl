package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2EipAssociationEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2EipAssociationEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2EipAssociationEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2EipAssociationEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2EipAssociationEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
