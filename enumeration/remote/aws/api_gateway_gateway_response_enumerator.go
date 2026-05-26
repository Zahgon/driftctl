package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayGatewayResponseEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayGatewayResponseEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayGatewayResponseEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayGatewayResponseEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayGatewayResponseEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
