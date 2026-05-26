package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// AwsALBTransformer is a simple middleware to turn all aws_alb resources into aws_lb ones
// Both types provide the same functionality, but we can't know which one was used to provision cloud resources.
// So we use aws_lb as the common type.
type AwsALBTransformer struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsALBTransformer(resourceFactory resource.ResourceFactory) AwsALBTransformer {
	_ = "STUB: not implemented"
	return *new(AwsALBTransformer)
}

func (m AwsALBTransformer) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}
