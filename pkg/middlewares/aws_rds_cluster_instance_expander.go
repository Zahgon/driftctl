package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// AwsRDSClusterInstanceExpander search for cluster instances from state to import corresponding remote db instances.
// RDS cluster instance does not represent an actual AWS resource, so shouldn't be used for comparison.
type AwsRDSClusterInstanceExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewRDSClusterInstanceExpander(resourceFactory resource.ResourceFactory) AwsRDSClusterInstanceExpander {
	_ = "STUB: not implemented"
	return *new(AwsRDSClusterInstanceExpander)
}

func (m AwsRDSClusterInstanceExpander) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than rds_cluster_instance

// If the db instance's id matches the rds cluster instance's id, import it in the state

// If we don't manage to find a db instance corresponding to this RDS cluster instance, simply add it back to the state.
