package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayV2DomainNameEnumerator struct {
	// AWS SDK list domain names endpoint from API Gateway v2 returns the
	// same results as the v1 one, thus let's re-use the method from
	// the API Gateway v1
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayV2DomainNameEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayV2DomainNameEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayV2DomainNameEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayV2DomainNameEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
