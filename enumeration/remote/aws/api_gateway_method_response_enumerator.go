package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayMethodResponseEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayMethodResponseEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayMethodResponseEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayMethodResponseEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayMethodResponseEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
