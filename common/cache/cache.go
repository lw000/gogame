package ggcache

import (
	"sync"

	"github.com/golang/groupcache/lru"
)

type LruCache struct {
	cache *lru.Cache
	m     sync.RWMutex
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		cache: lru.New(size),
		m:     sync.RWMutex{},
	}
}

func (lru *LruCache) Add(key interface{}, value interface{}) {
	lru.m.Lock()
	defer lru.m.Unlock()

	lru.cache.Add(key, value)
}

func (lru *LruCache) Remove(key interface{}) {
	lru.m.Lock()
	defer lru.m.Unlock()

	lru.cache.Remove(key)
}

func (lru *LruCache) Clear() {
	lru.m.Lock()
	defer lru.m.Unlock()

	lru.cache.Clear()
}

func (lru *LruCache) Get(key interface{}) (value interface{}, ok bool) {
	lru.m.Lock()
	defer lru.m.Unlock()

	return lru.cache.Get(key)
}

func (lru *LruCache) Len() int {
	lru.m.Lock()
	defer lru.m.Unlock()

	return lru.cache.Len()
}
