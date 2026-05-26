package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeSubnetworkEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeSubnetworkEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeSubnetworkEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeSubnetworkEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeSubnetworkEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
