package resource

import (
	"github.com/zclconf/go-cty/cty"
)

type Deserializer struct {
	factory ResourceFactory
}

func NewDeserializer(factory ResourceFactory) *Deserializer { _ = "STUB: not implemented"; return nil }

func (s *Deserializer) Deserialize(ty string, rawList []cty.Value) ([]*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Deserializer) DeserializeOne(ty string, value cty.Value) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marked values cannot be deserialized to JSON.
// For example, this ensures we can deserialize sensitive values too.
