package output

import (
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/pkg/analyser"
)

const FormatVersion = "0.1"
const PlanOutputType = "plan"
const PlanOutputExample = "plan://PATH/TO/FILE.json"

type plan struct {
	FormatVersion   string        `json:"format_version,omitempty"`
	PlannedValues   plannedValues `json:"planned_values,omitempty"`
	ResourceChanges []rscChange   `json:"resource_changes,omitempty"`
}

type plannedValues struct {
	RootModule module `json:"root_module,omitempty"`
}

type rscChange struct {
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Change  change `json:"change,omitempty"`
}

type change struct {
	Actions []string               `json:"actions,omitempty"`
	Before  map[string]interface{} `json:"before,omitempty"`
	After   map[string]interface{} `json:"after,omitempty"`
}

type module struct {
	Resources []rsc `json:"resources,omitempty"`
}

type rsc struct {
	Address         string                 `json:"address,omitempty"`
	Type            string                 `json:"type,omitempty"`
	Name            string                 `json:"name,omitempty"`
	AttributeValues map[string]interface{} `json:"values,omitempty"`
}

type Plan struct {
	path string
}

func NewPlan(path string) *Plan { _ = "STUB: not implemented"; return nil }

func (c *Plan) Write(analysis *analyser.Analysis) error { _ = "STUB: not implemented"; return nil }

func addPlannedValues(analysis *analyser.Analysis) module {
	_ = "STUB: not implemented"
	return *new(module)
}

func listRsc(resources []*resource.Resource) []rsc { _ = "STUB: not implemented"; return nil }

func addResourceChanges(analysis *analyser.Analysis) []rscChange {
	_ = "STUB: not implemented"
	return nil
}

func listRscChange(resources []*resource.Resource, action string) []rscChange {
	_ = "STUB: not implemented"
	return nil
}
