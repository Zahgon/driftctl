package remote

import (
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/parallel"
	"github.com/snyk/driftctl/enumeration/remote/common"
	"github.com/snyk/driftctl/enumeration/resource"
)

type Scanner struct {
	enumeratorRunner *parallel.ParallelRunner
	remoteLibrary    *common.RemoteLibrary
	alerter          alerter.AlerterInterface
	filter           enumeration.Filter
}

func NewScanner(remoteLibrary *common.RemoteLibrary, alerter alerter.AlerterInterface, filter enumeration.Filter) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) retrieveRunnerResults(runner *parallel.ParallelRunner) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scanner) scan() ([]*resource.Resource, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Scanner) Resources() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scanner) Stop() { _ = "STUB: not implemented"; return }
