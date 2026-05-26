package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayIntegrationEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayIntegrationEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayIntegrationEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayIntegrationEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayIntegrationEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
