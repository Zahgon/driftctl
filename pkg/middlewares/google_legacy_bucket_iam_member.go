package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Creating buckets add legacy role bindings, this middleware will filter them unless they are managed.
type GoogleLegacyBucketIAMMember struct{}

func NewGoogleLegacyBucketIAMMember() *GoogleLegacyBucketIAMMember {
	_ = "STUB: not implemented"
	return nil
}

func (m *GoogleLegacyBucketIAMMember) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than BucketIamBinding

// Ignore all non-legacy member

// Check if member is managed by IaC

// Include resource if it's managed in IaC

// Else, resource is not added to newRemoteResources slice, so it will be ignored
