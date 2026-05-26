package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeHealthCheckEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeHealthCheckEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeHealthCheckEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleComputeHealthCheckEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleComputeHealthCheckEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
