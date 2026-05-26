package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayIntegrationResponseEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayIntegrationResponseEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayIntegrationResponseEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayIntegrationResponseEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayIntegrationResponseEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
