package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// The role of this middleware is to expand policy contained in `managed_policy_arns` to dedicated `aws_iam_policy_attachment`
// resources. Note that we do not use `aws_iam_role_policy_attachment` or `aws_iam_user_policy_attachment`
// Once theses resources created, we remove the old `managed_policy_arns` field to avoid false positive drifts

type AwsRoleManagedPolicyExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsRoleManagedPolicyExpander(resourceFactory resource.ResourceFactory) *AwsRoleManagedPolicyExpander {
	_ = "STUB: not implemented"
	return nil
}

func (a AwsRoleManagedPolicyExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than iam_role

// Ignore all resources other than iam_role

// if managed_policy_arns does not exist or is empty ignore resource

// Remove empty slices to match remote read results
