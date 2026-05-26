package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayV2IntegrationEnumerator struct {
	repository repository.ApiGatewayV2Repository
	factory    resource.ResourceFactory
}

func NewApiGatewayV2IntegrationEnumerator(repo repository.ApiGatewayV2Repository, factory resource.ResourceFactory) *ApiGatewayV2IntegrationEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayV2IntegrationEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayV2IntegrationEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this is needed to discriminate in middleware. But it is nil when the type is mock...
