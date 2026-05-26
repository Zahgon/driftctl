package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Remote NS and SAO records from remote state if not managed by IAC
type Route53DefaultZoneRecordSanitizer struct{}

func NewRoute53DefaultZoneRecordSanitizer() Route53DefaultZoneRecordSanitizer {
	_ = "STUB: not implemented"
	return *new(Route53DefaultZoneRecordSanitizer)
}

func (m Route53DefaultZoneRecordSanitizer) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// We iterate on remote resource and adding them to a new slice except for default records
// added by aws in the zone at creation

// Ignore all resources other than route53 records

// Return true if the record is considered as default one added by aws
func isDefaultRecord(record *resource.Resource) bool { _ = "STUB: not implemented"; return false }
