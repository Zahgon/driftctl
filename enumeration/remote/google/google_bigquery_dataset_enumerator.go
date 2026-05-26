package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleBigqueryDatasetEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleBigqueryDatasetEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleBigqueryDatasetEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleBigqueryDatasetEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleBigqueryDatasetEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
