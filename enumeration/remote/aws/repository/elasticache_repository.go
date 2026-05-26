package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type ElastiCacheRepository interface {
	ListAllCacheClusters() ([]*elasticache.CacheCluster, error)
}

type elasticacheRepository struct {
	client elasticacheiface.ElastiCacheAPI
	cache  cache.Cache
}

func NewElastiCacheRepository(session *session.Session, c cache.Cache) *elasticacheRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *elasticacheRepository) ListAllCacheClusters() ([]*elasticache.CacheCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
