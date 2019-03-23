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

func (this *LruCache) Add(key interface{}, value interface{}) {
	this.m.Lock()
	defer this.m.Unlock()

	this.cache.Add(key, value)
}

func (this *LruCache) Remove(key interface{}) {
	this.m.Lock()
	defer this.m.Unlock()

	this.cache.Remove(key)
}

func (this *LruCache) Clear() {
	this.m.Lock()
	defer this.m.Unlock()

	this.cache.Clear()
}

func (this *LruCache) Get(key interface{}) (value interface{}, ok bool) {
	this.m.Lock()
	defer this.m.Unlock()

	return this.cache.Get(key)
}

func (this *LruCache) Len() int {
	this.m.Lock()
	defer this.m.Unlock()

	return this.cache.Len()
}
