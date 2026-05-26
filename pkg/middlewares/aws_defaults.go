package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

const defaultIamRolePathPrefix = "/aws-service-role/"

// AwsDefaults represents service-linked AWS resources
// When scanning a AWS account, some users may see irrelevant results about default AWS roles or role policies.
// We ignore these resources by default when strict mode is disabled.
type AwsDefaults struct{}

func NewAwsDefaults() AwsDefaults { _ = "STUB: not implemented"; return *new(AwsDefaults) }

func (m AwsDefaults) awsIamRoleDefaults(remoteResources []*resource.Resource) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than iam role

func (m AwsDefaults) awsIamRolePolicyDefaults(remoteResources []*resource.Resource) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than role policy

func (m AwsDefaults) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}
