package resource

import (
	"reflect"
)

type Source interface {
	Source() string
	Namespace() string
	InternalName() string
}

type SerializableSource struct {
	S    string `json:"source"`
	Ns   string `json:"namespace"`
	Name string `json:"internal_name"`
}

type TerraformStateSource struct {
	State  string
	Module string
	Name   string
}

func NewTerraformStateSource(state, module, name string) *TerraformStateSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *TerraformStateSource) Source() string { _ = "STUB: not implemented"; return "" }

func (s *TerraformStateSource) Namespace() string { _ = "STUB: not implemented"; return "" }

func (s *TerraformStateSource) InternalName() string { _ = "STUB: not implemented"; return "" }

type Resource struct {
	Id     string
	Type   string
	Attrs  *Attributes
	Sch    *Schema `json:"-" diff:"-"`
	Source Source  `json:"-"`
}

func (r *Resource) Schema() *Schema { _ = "STUB: not implemented"; return nil }

func (r *Resource) ResourceId() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) ResourceType() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) Attributes() *Attributes { _ = "STUB: not implemented"; return nil }

func (r *Resource) Src() Source { _ = "STUB: not implemented"; return *new(Source) }

func (r *Resource) SourceString() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) Equal(res *Resource) bool { _ = "STUB: not implemented"; return false }

type ResourceFactory interface {
	CreateAbstractResource(ty, id string, data map[string]interface{}) *Resource
}

type SerializableResource struct {
	Id                 string              `json:"id"`
	Type               string              `json:"type"`
	ReadableAttributes map[string]string   `json:"human_readable_attributes,omitempty"`
	Source             *SerializableSource `json:"source,omitempty"`
}

func NewSerializableResource(res *Resource) *SerializableResource {
	_ = "STUB: not implemented"
	return nil
}

func formatReadableAttributes(res *Resource) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

type NormalizedResource interface {
	NormalizeForState() (Resource, error)
	NormalizeForProvider() (Resource, error)
}

func Sort(res []*Resource) []*Resource { _ = "STUB: not implemented"; return nil }

type Attributes map[string]interface{}

func (a *Attributes) Copy() *Attributes { _ = "STUB: not implemented"; return nil }

func (a *Attributes) Get(path string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *Attributes) GetSlice(path string) []interface{} { _ = "STUB: not implemented"; return nil }

func (a *Attributes) GetString(path string) *string { _ = "STUB: not implemented"; return nil }

func (a *Attributes) GetBool(path string) *bool { _ = "STUB: not implemented"; return nil }

func (a *Attributes) GetInt(path string) *int { _ = "STUB: not implemented"; return nil }

func (a *Attributes) GetFloat64(path string) *float64 { _ = "STUB: not implemented"; return nil }

func (a *Attributes) GetMap(path string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (a *Attributes) SafeDelete(path []string) { _ = "STUB: not implemented"; return }

func (a *Attributes) SafeSet(path []string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// should not happen ?

func (a *Attributes) DeleteIfDefault(path string) { _ = "STUB: not implemented"; return }

func concatenatePath(path, next string) string { _ = "STUB: not implemented"; return "" }

func (a *Attributes) SanitizeDefaults() { _ = "STUB: not implemented"; return }

func (a *Attributes) sanitize(path string, original, copy reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// Get rid of the wrapping interface

// Create a new object. Now new gives us a pointer, but we want the value it
// points to, so we have to call Elem() to unwrap it
