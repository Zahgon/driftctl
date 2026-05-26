package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

/**
  Fetching eip association from remote return every association but some of them are embedded in eip.
  This middleware will check for every eip_association that here is no corresponding association_id inside eip.
*/

type EipAssociationExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewEipAssociationExpander(resourceFactory resource.ResourceFactory) EipAssociationExpander {
	_ = "STUB: not implemented"
	return *new(EipAssociationExpander)
}

func (m EipAssociationExpander) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// This EIP have no association, check if we need to create one

func (m EipAssociationExpander) haveMatchingEipAssociation(cur *resource.Resource, stateRes *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
