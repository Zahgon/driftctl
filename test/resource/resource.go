package resource_test

import (
	"github.com/snyk/driftctl/pkg/resource"
)

func InitFakeSchemaRepository(provider, version string) resource.SchemaRepositoryInterface {
	_ = "STUB: not implemented"
	return *new(resource.SchemaRepositoryInterface)
}

// TODO HANDLER ERROR PROPERLY
