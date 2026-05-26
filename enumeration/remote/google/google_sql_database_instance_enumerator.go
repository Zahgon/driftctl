package google

import (
	"github.com/snyk/driftctl/enumeration/remote/google/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleSQLDatabaseInstanceEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleSQLDatabaseInstanceEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleSQLDatabaseInstanceEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *GoogleSQLDatabaseInstanceEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *GoogleSQLDatabaseInstanceEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
