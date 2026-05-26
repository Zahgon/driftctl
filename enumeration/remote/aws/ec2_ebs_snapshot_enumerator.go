package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2EbsSnapshotEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2EbsSnapshotEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2EbsSnapshotEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2EbsSnapshotEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2EbsSnapshotEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
