package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Remove default security group rules of the default security group from remote resources
type AwsDefaultSecurityGroupRule struct{}

func NewAwsDefaultSecurityGroupRule() AwsDefaultSecurityGroupRule {
	_ = "STUB: not implemented"
	return *new(AwsDefaultSecurityGroupRule)
}

func (m AwsDefaultSecurityGroupRule) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than security group rules

// Ignore if it's not the default ingress or egress rule

func isDefaultIngress(rule *resource.Resource, remoteResources *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func isDefaultEgress(rule *resource.Resource, remoteResources *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func isFromDefaultSecurityGroup(sgId *string, remoteResources *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
