package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeGlobalAddressEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeGlobalAddressEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeGlobalAddressEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeGlobalAddressEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeGlobalAddressEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
