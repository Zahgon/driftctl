package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayRequestValidatorEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayRequestValidatorEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayRequestValidatorEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayRequestValidatorEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayRequestValidatorEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
