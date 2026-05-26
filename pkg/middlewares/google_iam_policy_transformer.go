package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// GoogleStorageBucketIAMPolicyTransformer Transforms Bucket IAM policy in bucket iam binding to ease comparison.
type GoogleStorageBucketIAMPolicyTransformer struct {
	resourceFactory resource.ResourceFactory
	resFieldByType  map[string]string // map of the field to add to resource attribute for all supported type
}

func NewGoogleIAMPolicyTransformer(resourceFactory resource.ResourceFactory) *GoogleStorageBucketIAMPolicyTransformer {
	_ = "STUB: not implemented"
	return nil
}

func (m *GoogleStorageBucketIAMPolicyTransformer) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources with type not in resFieldByType map

type policyDataType struct {
	Bindings []map[string]interface{}
}
