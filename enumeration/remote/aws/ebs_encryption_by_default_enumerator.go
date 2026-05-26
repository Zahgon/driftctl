package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type EC2EbsEncryptionByDefaultEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewEC2EbsEncryptionByDefaultEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *EC2EbsEncryptionByDefaultEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *EC2EbsEncryptionByDefaultEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *EC2EbsEncryptionByDefaultEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
