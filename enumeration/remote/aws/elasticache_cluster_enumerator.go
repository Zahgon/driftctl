package aws

import (
	"github.com/snyk/driftctl/enumeration/remote/aws/repository"
	"github.com/snyk/driftctl/enumeration/resource"
)

type ElastiCacheClusterEnumerator struct {
	repository repository.ElastiCacheRepository
	factory    resource.ResourceFactory
}

func NewElastiCacheClusterEnumerator(repo repository.ElastiCacheRepository, factory resource.ResourceFactory) *ElastiCacheClusterEnumerator {
	_ = "STUB: not implemented"
	return nil
}

func (e *ElastiCacheClusterEnumerator) SupportedType() resource.ResourceType {
	_ = "STUB: not implemented"
	return *new(resource.ResourceType)
}

func (e *ElastiCacheClusterEnumerator) Enumerate() ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
