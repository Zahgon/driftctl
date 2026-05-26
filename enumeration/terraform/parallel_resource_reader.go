package terraform

import (
	"github.com/snyk/driftctl/enumeration/parallel"
	"github.com/zclconf/go-cty/cty"
)

type ParallelResourceReader struct {
	runner *parallel.ParallelRunner
}

func NewParallelResourceReader(runner *parallel.ParallelRunner) *ParallelResourceReader {
	_ = "STUB: not implemented"
	return nil
}

func (p *ParallelResourceReader) Wait() ([]cty.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ParallelResourceReader) Run(runnable func() (cty.Value, error)) {
	_ = "STUB: not implemented"
	return
}
