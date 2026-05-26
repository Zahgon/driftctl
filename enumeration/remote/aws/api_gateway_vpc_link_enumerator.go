package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayVpcLinkEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayVpcLinkEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayVpcLinkEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayVpcLinkEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayVpcLinkEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
