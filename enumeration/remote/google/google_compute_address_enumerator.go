package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeAddressEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeAddressEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeAddressEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeAddressEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeAddressEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Global addresses are handled as a dedicated resource
