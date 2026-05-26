package middlewares

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

// Remove root_block_device from aws_instance resources and create dedicated aws_ebs_volume resources
type AwsInstanceBlockDeviceResourceMapper struct {
	resourceFactory resource.ResourceFactory
}

func NewAwsInstanceBlockDeviceResourceMapper(resourceFactory resource.ResourceFactory) AwsInstanceBlockDeviceResourceMapper {
	_ = "STUB: not implemented"
	return *new(AwsInstanceBlockDeviceResourceMapper)
}

func (a AwsInstanceBlockDeviceResourceMapper) Execute(remoteResources, resourcesFromState *[]*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore all resources other than aws_instance

func (a AwsInstanceBlockDeviceResourceMapper) volumeTags(instance *resource.Resource, blockDevice map[string]interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (a AwsInstanceBlockDeviceResourceMapper) hasBlockDevice(blockDevice map[string]interface{}, resourcesFromState *[]*resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}
