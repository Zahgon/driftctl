package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeInstanceGroupManagerEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeInstanceGroupManagerEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeInstanceGroupManagerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeInstanceGroupManagerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeInstanceGroupManagerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
