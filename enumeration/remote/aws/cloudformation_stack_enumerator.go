package aws

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hashicorp/terraform/flatmap"
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type CloudformationStackEnumerator struct {
	repository repository.CloudformationRepository
	factory    resource.ResourceFactory
}

func NewCloudformationStackEnumerator(repo repository.CloudformationRepository, factory resource.ResourceFactory) *CloudformationStackEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *CloudformationStackEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *CloudformationStackEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func flattenParameters(parameters []*cloudformation.Parameter) flatmap.Map {
	_ = "STUB: not implemented"
	return *new(flatmap.Map)
}
