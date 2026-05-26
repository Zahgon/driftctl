package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayResourceEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayResourceEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayResourceEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayResourceEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayResourceEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
