package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// AwsALBListenerTransformer is a simple middleware to turn all aws_alb_listener resources into aws_lb_listener ones
// Both types provide the same functionality, but we can't know which one was used to provision cloud resources.
// So we use aws_lb_listener as the common type.
type AwsALBListenerTransformer struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsALBListenerTransformer(resourceFactory resource.ResourceFactory) AwsALBListenerTransformer {
	_ = "STUB: not implemented"
	return *new(AwsALBListenerTransformer)
}

func (m AwsALBListenerTransformer) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}
