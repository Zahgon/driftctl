package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleCloudRunServiceEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleCloudRunServiceEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleCloudRunServiceEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleCloudRunServiceEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleCloudRunServiceEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
