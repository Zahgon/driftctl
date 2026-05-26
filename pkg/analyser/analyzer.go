package analyser

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/pkg/filter"

	"github.com/snyk/driftctl/enumeration/resource"
)

type UnmanagedSecurityGroupRulesAlert struct{}

func newUnmanagedSecurityGroupRulesAlert() *UnmanagedSecurityGroupRulesAlert {
	_ = "STUB: not implemented"
	return nil
}

func (u *UnmanagedSecurityGroupRulesAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (u *UnmanagedSecurityGroupRulesAlert) ShouldIgnoreResource() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UnmanagedSecurityGroupRulesAlert) Resource() *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

type ComputedDiffAlert struct{}

func NewComputedDiffAlert() *ComputedDiffAlert { _ = "STUB: not implemented"; return nil }

func (c *ComputedDiffAlert) Message() string { _ = "STUB: not implemented"; return "" }

func (c *ComputedDiffAlert) ShouldIgnoreResource() bool { _ = "STUB: not implemented"; return false }

func (c *ComputedDiffAlert) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

type Analyzer struct {
	alerter *alerter.Alerter
	filter  filter.Filter
}

func NewAnalyzer(alerter *alerter.Alerter, filter filter.Filter) *Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (a Analyzer) Analyze(remoteResources, resourcesFromState []*resource.Resource) (Analysis, error) {
	_ = "STUB: not implemented"
	return *

	// Iterate on remote resources and filter ignored resources
	new(Analysis), nil
}

// Remove managed resources, so it will remain only unmanaged ones

// Add remaining unmanaged resources

// Sort resources by Terraform Id
// The purpose is to have a predictable output

func findCorrespondingRes(resources []*resource.Resource, res *resource.Resource) (int, *resource.Resource, bool) {
	_ = "STUB: not implemented"
	return 0, nil, false
}

func removeResourceByIndex(i int, resources []*resource.Resource) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// hasUnmanagedSecurityGroupRules returns true if we find at least one unmanaged
// security group rule
func (a Analyzer) hasUnmanagedSecurityGroupRules(unmanagedResources []*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
