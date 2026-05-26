package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type RDSRepository interface {
	ListAllDBInstances() ([]*rds.DBInstance, error)
	ListAllDBSubnetGroups() ([]*rds.DBSubnetGroup, error)
	ListAllDBClusters() ([]*rds.DBCluster, error)
}

type rdsRepository struct {
	client rdsiface.RDSAPI
	cache  cache.Cache
}

func NewRDSRepository(session *session.Session, c cache.Cache) *rdsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *rdsRepository) ListAllDBInstances() ([]*rds.DBInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *rdsRepository) ListAllDBSubnetGroups() ([]*rds.DBSubnetGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *rdsRepository) ListAllDBClusters() ([]*rds.DBCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
