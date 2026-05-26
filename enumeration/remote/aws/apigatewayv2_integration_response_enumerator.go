package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayV2IntegrationResponseEnumerator struct {
	repository repository.ApiGatewayV2Repository
	factory    resource.ResourceFactory
}

func NewApiGatewayV2IntegrationResponseEnumerator(repo repository.ApiGatewayV2Repository, factory resource.ResourceFactory) *ApiGatewayV2IntegrationResponseEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayV2IntegrationResponseEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayV2IntegrationResponseEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
