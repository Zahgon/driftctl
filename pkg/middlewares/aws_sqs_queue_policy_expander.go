package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

// Explodes policy found in aws_sqs_queue.policy from state resources to dedicated resources
type AwsSQSQueuePolicyExpander struct {
	resourceFactory          resource.ResourceFactory
	resourceSchemaRepository dctlresource.SchemaRepositoryInterface
}

func NewAwsSQSQueuePolicyExpander(resourceFactory resource.ResourceFactory, resourceSchemaRepository dctlresource.SchemaRepositoryInterface) AwsSQSQueuePolicyExpander {
	_ = "STUB: not implemented"
	return *new(AwsSQSQueuePolicyExpander)
}

func (m AwsSQSQueuePolicyExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than sqs_queue

func (m *AwsSQSQueuePolicyExpander) handlePolicy(queue *resource.Resource, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Return true if the sqs queue has a aws_sqs_queue_policy resource attached to itself.
// It is mandatory since it's possible to have a aws_sqs_queue with an inline policy
// AND a aws_sqs_queue_policy resource at the same time. At the end, on the AWS console,
// the aws_sqs_queue_policy will be used.
func (m *AwsSQSQueuePolicyExpander) hasPolicyAttached(queue *resource.Resource, resourcesFromState *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
