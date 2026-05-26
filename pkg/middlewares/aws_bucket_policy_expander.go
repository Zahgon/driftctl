package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Explodes policy found in aws_s3_bucket.policy from state resources to dedicated resources
type AwsBucketPolicyExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsBucketPolicyExpander(resourceFactory resource.ResourceFactory) AwsBucketPolicyExpander {
	_ = "STUB: not implemented"
	return *new(AwsBucketPolicyExpander)
}

func (m AwsBucketPolicyExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than s3_bucket

func (m *AwsBucketPolicyExpander) handlePolicy(bucket *resource.Resource, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Return true if the bucket has a aws_bucket_policy resource attached to itself.
// It is mandatory since it's possible to have a aws_bucket with an inline policy
// AND a aws_bucket_policy resource at the same time. At the end, on the AWS console,
// the aws_bucket_policy will be used.
func hasPolicyAttached(bucket string, resourcesFromState *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
