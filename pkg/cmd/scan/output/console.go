package output

import (
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/pkg/analyser"
)

const ConsoleOutputType = "console"
const ConsoleOutputExample = "console://"

type Console struct {
	summary string
}

func NewConsole() *Console { _ = "STUB: not implemented"; return nil }

func (c *Console) Write(analysis *analyser.Analysis) error { _ = "STUB: not implemented"; return nil }

func (c Console) writeSummary(analysis *analyser.Analysis) { _ = "STUB: not implemented"; return }

func groupByType(resources []*resource.Resource) (map[string][]*resource.Resource, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formatResourceAttributes(res *resource.Resource) string { _ = "STUB: not implemented"; return "" }

// sort attributes

// retrieve stringer
