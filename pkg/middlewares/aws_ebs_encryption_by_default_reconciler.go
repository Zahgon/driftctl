package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// AwsEbsEncryptionByDefaultReconciler is a middleware that either creates an 'aws_ebs_encryption_by_default' resource
// based on its equivalent state one just for the purpose of getting the Terraform custom Id, or removes the resource
// from our list of remote resources if it is not managed and is disabled.
type AwsEbsEncryptionByDefaultReconciler struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsEbsEncryptionByDefaultReconciler(resourceFactory resource.ResourceFactory) AwsEbsEncryptionByDefaultReconciler {
	_ = "STUB: not implemented"
	return *new(AwsEbsEncryptionByDefaultReconciler)
}

func (m AwsEbsEncryptionByDefaultReconciler) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_ebs_encryption_by_default

// We can encounter this case when we don't have permission to get this setting from AWS.

// Ignore all resources other than aws_ebs_encryption_by_default

// Create a new remote resource that will be similar to the state resource but with the 'enabled' attribute of the remote one.
// The reason why is that the id is a random string created by Terraform that we need to compare two resources.
