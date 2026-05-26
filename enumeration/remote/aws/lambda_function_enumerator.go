package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type LambdaFunctionEnumerator struct {
	repository repository.LambdaRepository
	factory    resource.ResourceFactory
}

func NewLambdaFunctionEnumerator(repo repository.LambdaRepository, factory resource.ResourceFactory) *LambdaFunctionEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *LambdaFunctionEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *LambdaFunctionEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
