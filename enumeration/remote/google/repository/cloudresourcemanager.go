package repository

import (
	"github.com/snyk/driftctl/enumeration/remote/cache"
	"github.com/snyk/driftctl/enumeration/remote/google/config"
	"google.golang.org/api/cloudresourcemanager/v1"
)

type CloudResourceManagerRepository interface {
	ListProjectsBindings() (map[string]map[string][]string, error)
}

type cloudResourceManagerRepository struct {
	service *cloudresourcemanager.Service
	config  config.GCPTerraformConfig
	cache   cache.Cache
}

func NewCloudResourceManagerRepository(service *cloudresourcemanager.Service, config config.GCPTerraformConfig, cache cache.Cache) CloudResourceManagerRepository {
	_ = "STUB: not implemented"
	return *new(CloudResourceManagerRepository)
}

func (s *cloudResourceManagerRepository) ListProjectsBindings() (map[string]map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
