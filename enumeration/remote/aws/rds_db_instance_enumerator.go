package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type RDSDBInstanceEnumerator struct {
	repository repository.RDSRepository
	factory    resource.ResourceFactory
}

func NewRDSDBInstanceEnumerator(repo repository.RDSRepository, factory resource.ResourceFactory) *RDSDBInstanceEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *RDSDBInstanceEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *RDSDBInstanceEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
