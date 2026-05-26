package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type ApiGatewayMethodSettingsEnumerator struct {
	repository repository.ApiGatewayRepository
	factory    resource.ResourceFactory
}

func NewApiGatewayMethodSettingsEnumerator(repo repository.ApiGatewayRepository, factory resource.ResourceFactory) *ApiGatewayMethodSettingsEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ApiGatewayMethodSettingsEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ApiGatewayMethodSettingsEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
