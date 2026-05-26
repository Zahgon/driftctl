package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// GoogleIAMBindingTransformer Transforms Bucket IAM binding in bucket iam member to ease comparison.
type GoogleIAMBindingTransformer struct {
	resourceFactory resource.ResourceFactory
	resFieldByType  map[string]string // map of the field to add to resource attribute for all supported type
}

func NewGoogleIAMBindingTransformer(resourceFactory resource.ResourceFactory) *GoogleIAMBindingTransformer {
	_ = "STUB: not implemented"
	return nil
}

func (m *GoogleIAMBindingTransformer) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than IamBinding
