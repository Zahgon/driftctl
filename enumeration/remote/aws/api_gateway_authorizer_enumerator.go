package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayAuthorizerEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayAuthorizerEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayAuthorizerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayAuthorizerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayAuthorizerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
