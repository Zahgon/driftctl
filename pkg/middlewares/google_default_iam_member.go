package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Some service accounts are created by default when activating APIs, this middleware will filter them unless they are managed.
type GoogleDefaultIAMMember struct{}

func NewGoogleDefaultIAMMember() *GoogleDefaultIAMMember { _ = "STUB: not implemented"; return nil }

func (m *GoogleDefaultIAMMember) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than BucketIamBinding

// Ignore all non service account member

// Ignore all service accounts that have project host

// Check if member is managed by IaC

// Include resource if it's managed by IaC

// Else, resource is not added to newRemoteResources slice, so it will be ignored
