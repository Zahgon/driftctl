package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

/**
  When listing policy attachment from aws we retrieve only user_policy_attachment or role_policy_attachment thus making it
  impossible to compare with policy_attachment that could exist in terraform.
  We decided to transform all attachments to policy_attachment so we can find which attachments are managed.
*/

type IamPolicyAttachmentTransformer struct {
	resourceFactory resource.ResourceFactory
}

func NewIamPolicyAttachmentTransformer(resourceFactory resource.ResourceFactory) IamPolicyAttachmentTransformer {
	_ = "STUB: not implemented"
	return *new(IamPolicyAttachmentTransformer)
}

func (m IamPolicyAttachmentTransformer) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m IamPolicyAttachmentTransformer) transform(resources *[]*resource.Resource) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}
