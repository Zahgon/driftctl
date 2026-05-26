package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayV2ApiEnumerator struct {
	repository repository.ApiGatewayV2Repository
	factory    resource.ResourceFactory
}

func NewApiGatewayV2ApiEnumerator(repo repository.ApiGatewayV2Repository, factory resource.ResourceFactory) *ApiGatewayV2ApiEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayV2ApiEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayV2ApiEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
