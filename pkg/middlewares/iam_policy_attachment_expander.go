package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Split Policy attachment when there is multiple user and groups and generate a repeatable id
type IamPolicyAttachmentExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewIamPolicyAttachmentExpander(resourceFactory resource.ResourceFactory) IamPolicyAttachmentExpander {
	_ = "STUB: not implemented"
	return *new(IamPolicyAttachmentExpander)
}

func (m IamPolicyAttachmentExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than policy attachment

// Ignore all resources other than policy attachment

func (m IamPolicyAttachmentExpander) expand(policyAttachment *resource.Resource) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

// we create one attachment per user

// we create one attachment per role

// we create one attachment per group
