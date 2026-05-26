package memstore

import (
	"sync"
)

type Bucket interface {
	Set(string, interface{})
	Get(string) interface{}
	Values() map[string]interface{}
}

type bucket struct {
	m      *sync.RWMutex
	values map[string]interface{}
}

func (b bucket) Set(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (b bucket) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (b bucket) Values() map[string]interface{} { _ = "STUB: not implemented"; return nil }
