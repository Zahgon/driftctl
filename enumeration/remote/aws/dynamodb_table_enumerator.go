package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type DynamoDBTableEnumerator struct {
	repository repository.DynamoDBRepository
	factory    resource.ResourceFactory
}

func NewDynamoDBTableEnumerator(repository repository.DynamoDBRepository, factory resource.ResourceFactory) *DynamoDBTableEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *DynamoDBTableEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *DynamoDBTableEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
