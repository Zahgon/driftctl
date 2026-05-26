package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"

	"github.com/snyk/driftctl/enumeration/resource"
)

type Route53RecordEnumerator struct {
	client  repository.Route53Repository
	factory resource.ResourceFactory
}

func NewRoute53RecordEnumerator(repo repository.Route53Repository, factory resource.ResourceFactory) *Route53RecordEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *Route53RecordEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *Route53RecordEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Route53RecordEnumerator) listRecordsForZone(zoneId string) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cleanRecordName
// Route 53 stores certain characters with the octal equivalent in ASCII format.
// This function converts all of these characters back into the original character.
// E.g. "*" is stored as "\\052" and "@" as "\\100"
func (e *Route53RecordEnumerator) cleanRecordName(name string) string {
	_ = "STUB: not implemented"
	return ""
}
