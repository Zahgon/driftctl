package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleCloudFunctionsFunctionEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleCloudFunctionsFunctionEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleCloudFunctionsFunctionEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleCloudFunctionsFunctionEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleCloudFunctionsFunctionEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
