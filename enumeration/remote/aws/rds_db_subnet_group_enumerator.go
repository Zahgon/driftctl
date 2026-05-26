package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type RDSDBSubnetGroupEnumerator struct {
	repository repository.RDSRepository
	factory    resource.ResourceFactory
}

func NewRDSDBSubnetGroupEnumerator(repo repository.RDSRepository, factory resource.ResourceFactory) *RDSDBSubnetGroupEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *RDSDBSubnetGroupEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *RDSDBSubnetGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
