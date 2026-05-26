package memstore

import (
	"sync"
)

type Store interface {
	Bucket(BucketName) Bucket
}

type store struct {
	m       *sync.Mutex
	buckets map[int]*bucket
}

func New() Store { _ = "STUB: not implemented"; return *new(Store) }

func (s store) Bucket(name BucketName) Bucket { _ = "STUB: not implemented"; return *new(Bucket) }
