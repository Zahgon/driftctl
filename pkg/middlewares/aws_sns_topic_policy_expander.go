package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

// Explodes policy found in aws_sns_topic from state resources to aws_sns_topic_policy resources
type AwsSNSTopicPolicyExpander struct {
	resourceFactory          resource.ResourceFactory
	resourceSchemaRepository dctlresource.SchemaRepositoryInterface
}

func NewAwsSNSTopicPolicyExpander(resourceFactory resource.ResourceFactory, resourceSchemaRepository dctlresource.SchemaRepositoryInterface) AwsSNSTopicPolicyExpander {
	_ = "STUB: not implemented"
	return *new(AwsSNSTopicPolicyExpander)
}

func (m AwsSNSTopicPolicyExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than sns_topic

func (m *AwsSNSTopicPolicyExpander) splitPolicy(topic *resource.Resource, results *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *AwsSNSTopicPolicyExpander) hasPolicyAttached(topic *resource.Resource, resourcesFromState *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
