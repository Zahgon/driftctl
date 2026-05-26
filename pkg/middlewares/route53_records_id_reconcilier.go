package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Since AWS returns the FQDN as the name of the remote record, we must change the Id of the
// state record to be equivalent (ZoneId_FQDN_Type_SetIdentifier)
// For a TXT record toto for zone example.com with Id 1234
// From AWS provider, we retrieve: 1234_toto.example.com_TXT
// From Terraform state, we retrieve: 1234_toto_TXT
type Route53RecordIDReconcilier struct{}

func NewRoute53RecordIDReconcilier() Route53RecordIDReconcilier {
	_ = "STUB: not implemented"
	return *new(Route53RecordIDReconcilier)
}

func (m Route53RecordIDReconcilier) Execute(_, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}
