package filter

import (
	"github.com/jmespath/go-jmespath"
	"github.com/snyk/driftctl/enumeration/resource"
)

type FilterEngine struct {
	expr *jmespath.JMESPath
}

func NewFilterEngine(expr *jmespath.JMESPath) *FilterEngine { _ = "STUB: not implemented"; return nil }

type filtrableResource struct {
	Attr     interface{}
	Res      *resource.Resource
	Type, Id string
}

func (e *FilterEngine) Run(resources []*resource.Resource) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We convert a list of resource in a list of DTO to run JMESPath on

// We need to serialize all attributes to untyped interface from JMESPath to work
// map[string]string and map[string]SomeThing will not work without it
// https://github.com/jmespath/go-jmespath/issues/22

// Do the filter

// Convert back filtered results into a resource list
