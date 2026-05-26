package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeNetworkEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeNetworkEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeNetworkEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeNetworkEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeNetworkEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
