package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// This middelware goal is to explode aws_network_acl ingress and egress block into a set of aws_network_acl_rule
type AwsNetworkACLExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsNetworkACLExpander(resourceFactory resource.ResourceFactory) AwsNetworkACLExpander {
	_ = "STUB: not implemented"
	return *new(AwsNetworkACLExpander)
}

func (m AwsNetworkACLExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than network acl

// Then we need to remove ingress and egress block from remote resource too

func (e *AwsNetworkACLExpander) expandBlock(resourcesFromState *[]*resource.Resource, networkAclId string, egress bool, ruleBlock []interface{}) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}
