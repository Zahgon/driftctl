package repository

import (
	"sync"

	"cloud.google.com/go/storage"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type StorageRepository interface {
	ListAllBindings(bucketName string) (map[string][]string, error)
}

type storageRepository struct {
	client *storage.Client
	cache  cache.Cache
	lock   sync.Locker
}

func NewStorageRepository(client *storage.Client, cache cache.Cache) *storageRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s storageRepository) ListAllBindings(bucketName string) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
