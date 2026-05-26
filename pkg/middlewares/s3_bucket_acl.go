package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Remove grant field on remote resources when acl field != private in state
type S3BucketAcl struct{}

func NewS3BucketAcl() S3BucketAcl { _ = "STUB: not implemented"; return *new(S3BucketAcl) }

func (m S3BucketAcl) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than s3 buckets
