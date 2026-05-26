package schemas

import (
	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/hashicorp/terraform/providers"
	"github.com/snyk/driftctl/enumeration/resource"
)

type SchemaRepository struct {
	schemas map[string]*resource.Schema
}

func NewSchemaRepository() *SchemaRepository { _ = "STUB: not implemented"; return nil }

func (r *SchemaRepository) GetSchema(resourceType string) (*resource.Schema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *SchemaRepository) fetchNestedBlocks(root string, metadata map[string]resource.AttributeSchema, block map[string]*configschema.NestedBlock) {
	_ = "STUB: not implemented"
	return
}

func (r *SchemaRepository) Init(providerName, providerVersion string, schema map[string]providers.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (r SchemaRepository) SetFlags(typ string, flags ...resource.Flags) {
	_ = "STUB: not implemented"
	return
}

func (r *SchemaRepository) UpdateSchema(typ string, schemasMutators map[string]func(attributeSchema *resource.AttributeSchema)) {
	_ = "STUB: not implemented"
	return
}

func (r *SchemaRepository) SetNormalizeFunc(typ string, normalizeFunc func(res *resource.Resource)) {
	_ = "STUB: not implemented"
	return
}

func (r *SchemaRepository) SetHumanReadableAttributesFunc(typ string, humanReadableAttributesFunc func(res *resource.Resource) map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (r *SchemaRepository) SetDiscriminantFunc(typ string, fn func(self, res *resource.Resource) bool) {
	_ = "STUB: not implemented"
	return
}
