package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayRestApiPolicyEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayRestApiPolicyEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayRestApiPolicyEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayRestApiPolicyEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayRestApiPolicyEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
