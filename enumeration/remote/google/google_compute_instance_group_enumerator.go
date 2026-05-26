package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeInstanceGroupEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeInstanceGroupEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeInstanceGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeInstanceGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeInstanceGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
