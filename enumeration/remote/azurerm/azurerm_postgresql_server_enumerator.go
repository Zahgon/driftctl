package azurerm

import (
	"github.com/snyk/driftctl/enumeration/remote/azurerm/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type AzurermPostgresqlServerEnumerator struct {
	repository repository.PostgresqlRespository
	factory    resource.ResourceFactory
}

func NewAzurermPostgresqlServerEnumerator(repo repository.PostgresqlRespository, factory resource.ResourceFactory) *AzurermPostgresqlServerEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *AzurermPostgresqlServerEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *AzurermPostgresqlServerEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
