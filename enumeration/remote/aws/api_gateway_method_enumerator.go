package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayMethodEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayMethodEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayMethodEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayMethodEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayMethodEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
