package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleBigqueryTableEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleBigqueryTableEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleBigqueryTableEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleBigqueryTableEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleBigqueryTableEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
