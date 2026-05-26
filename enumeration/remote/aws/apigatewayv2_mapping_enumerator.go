package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayV2MappingEnumerator struct {
	repository   repository.ApiGatewayV2Repository
	repositoryV1 repository.ApiGatewayRepository
	factory      resource.ResourceFactory
}

func NewApiGatewayV2MappingEnumerator(repo repository.ApiGatewayV2Repository, repov1 repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayV2MappingEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayV2MappingEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayV2MappingEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
