package terraform

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type TerraformResourceFactory struct{}

func NewTerraformResourceFactory() *TerraformResourceFactory { _ = "STUB: not implemented"; return nil }

func (r *TerraformResourceFactory) CreateAbstractResource(ty, id string, data map[string]interface{}) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}
