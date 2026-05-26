package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type LambdaEventSourceMappingEnumerator struct {
	repository repository.LambdaRepository
	factory    resource.ResourceFactory
}

func NewLambdaEventSourceMappingEnumerator(repo repository.LambdaRepository, factory resource.ResourceFactory) *LambdaEventSourceMappingEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *LambdaEventSourceMappingEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *LambdaEventSourceMappingEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
