package output

import (
	"github.com/snyk/driftctl/pkg/analyser"
)

const JSONOutputType = "json"
const JSONOutputExample = "json://PATH/TO/FILE.json"

type JSON struct {
	path string
}

func NewJSON(path string) *JSON { _ = "STUB: not implemented"; return nil }

func (c *JSON) Write(analysis *analyser.Analysis) error { _ = "STUB: not implemented"; return nil }
