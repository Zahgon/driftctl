package resource

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type ResourceFactory interface {
	CreateAbstractResource(ty, id string, data map[string]interface{}) *resource.Resource
}

type DriftctlResourceFactory struct {
	resourceSchemaRepository SchemaRepositoryInterface
}

func NewDriftctlResourceFactory(resourceSchemaRepository SchemaRepositoryInterface) *DriftctlResourceFactory {
	_ = "STUB: not implemented"
	return nil
}

func (r *DriftctlResourceFactory) CreateAbstractResource(ty, id string, data map[string]interface{}) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}
