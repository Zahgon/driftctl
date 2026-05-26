package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleBigtableTableEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleBigtableTableEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleBigtableTableEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleBigtableTableEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleBigtableTableEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
