package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Split security group rule if it needs to given its attributes
type VPCSecurityGroupRuleSanitizer struct {
	resourceFactory resource.ResourceFactory
}

func NewVPCSecurityGroupRuleSanitizer(resourceFactory resource.ResourceFactory) VPCSecurityGroupRuleSanitizer {
	_ = "STUB: not implemented"
	return *new(VPCSecurityGroupRuleSanitizer)
}

func (m VPCSecurityGroupRuleSanitizer) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than security group rule

func (m *VPCSecurityGroupRuleSanitizer) createRule(res *resource.Attributes) *resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

func shouldBeSplit(r *resource.Resource) bool { _ = "STUB: not implemented"; return false }
