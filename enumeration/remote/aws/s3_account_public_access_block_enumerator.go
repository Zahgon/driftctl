package aws

import (
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type S3AccountPublicAccessBlockEnumerator struct {
	repository repository.S3ControlRepository
	factory    resource.ResourceFactory
	accountID  string
	alerter    alerter.AlerterInterface
}

func NewS3AccountPublicAccessBlockEnumerator(repo repository.S3ControlRepository, factory resource.ResourceFactory, accountId string, alerter alerter.AlerterInterface) *S3AccountPublicAccessBlockEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *S3AccountPublicAccessBlockEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *S3AccountPublicAccessBlockEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
