package resource

import (
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/zclconf/go-cty/cty"
)

type Deserializer struct {
	factory resource.ResourceFactory
}

func NewDeserializer(factory resource.ResourceFactory) *Deserializer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Deserializer) Deserialize(ty string, rawList []cty.Value) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Deserializer) DeserializeOne(ty string, value cty.Value) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marked values cannot be deserialized to JSON.
// For example, this ensures we can deserialize sensitive values too.
