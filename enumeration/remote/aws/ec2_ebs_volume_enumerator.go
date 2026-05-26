package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2EbsVolumeEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2EbsVolumeEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2EbsVolumeEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2EbsVolumeEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2EbsVolumeEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
