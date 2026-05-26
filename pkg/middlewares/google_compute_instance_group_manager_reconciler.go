package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

type GoogleComputeInstanceGroupManagerReconciler struct{}

// NewGoogleComputeInstanceGroupManagerReconciler imports remote instance groups when they're managed by a managed instance group manager.
// Creating a "google_compute_instance_group_manager" resource via Terraform leads to having several unmanaged instance groups.
// This middleware adds remote instance groups to the state by matching them with managed instance group managers.
func NewGoogleComputeInstanceGroupManagerReconciler() *GoogleComputeInstanceGroupManagerReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (a GoogleComputeInstanceGroupManagerReconciler) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than google_compute_instance_group

// Ignore all resources other than google_compute_instance_group_manager

// Import instance group in the state
