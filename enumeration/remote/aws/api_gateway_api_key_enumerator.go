package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayApiKeyEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayApiKeyEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayApiKeyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayApiKeyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayApiKeyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
