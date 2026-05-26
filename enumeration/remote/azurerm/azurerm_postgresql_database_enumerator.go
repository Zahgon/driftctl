package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermPostgresqlDatabaseEnumerator struct {
	repository repository.PostgresqlRespository
	factory    resource.ResourceFactory
}

func NewAzurermPostgresqlDatabaseEnumerator(repo repository.PostgresqlRespository, factory resource.ResourceFactory) *AzurermPostgresqlDatabaseEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermPostgresqlDatabaseEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermPostgresqlDatabaseEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
