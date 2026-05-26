package cache

import (
	"container/list"
	"sync"
)

type Cache interface {
	Put(string, interface{}) bool
	Get(string) interface{}
	GetAndLock(string) interface{}
	Unlock(string)
	Len() int
}

type LRUCache struct {
	cap     int
	mu      *sync.Mutex
	l       *list.List
	m       map[string]*list.Element
	lockMap *sync.Map
}

type pair struct {
	key   string
	value interface{}
}

func New(capacity int) Cache { _ = "STUB: not implemented"; return *new(Cache) }

func (c *LRUCache) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

// if the key exists, move to front

func (c *LRUCache) Put(key string, value interface{}) bool { _ = "STUB: not implemented"; return false }

// if the key already exists, move to front and update the value

// if the list is full, delete the last element

// initialize a new list node

func (c *LRUCache) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *LRUCache) GetAndLock(s string) interface{} { _ = "STUB: not implemented"; return nil }

func (c *LRUCache) Unlock(s string) { _ = "STUB: not implemented"; return }
