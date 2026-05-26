package resource

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform/configs/configschema"
)

type AttributeSchema struct {
	ConfigSchema configschema.Attribute
	JsonString   bool
}

type Flags uint32

func (f Flags) HasFlag(flag Flags) bool { _ = "STUB: not implemented"; return false }

func (f *Flags) AddFlag(flag Flags) { _ = "STUB: not implemented"; return }

type Schema struct {
	ProviderVersion             *version.Version
	Flags                       Flags
	SchemaVersion               int64
	Attributes                  map[string]AttributeSchema
	NormalizeFunc               func(res *Resource)
	HumanReadableAttributesFunc func(res *Resource) map[string]string
	DiscriminantFunc            func(*Resource, *Resource) bool
}

func (s *Schema) IsComputedField(path []string) bool { _ = "STUB: not implemented"; return false }

func (s *Schema) IsJsonStringField(path []string) bool { _ = "STUB: not implemented"; return false }
