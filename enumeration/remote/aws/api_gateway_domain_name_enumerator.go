package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayDomainNameEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayDomainNameEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayDomainNameEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayDomainNameEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayDomainNameEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
