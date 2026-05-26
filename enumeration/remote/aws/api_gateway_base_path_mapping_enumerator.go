package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayBasePathMappingEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayBasePathMappingEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayBasePathMappingEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayBasePathMappingEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayBasePathMappingEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
