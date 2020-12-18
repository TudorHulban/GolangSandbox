package cache

import (
	"sync"
	"time"
)

// Cache contains the items.
type Cache struct {
	defaultExpirationSeconds int64           //amount of time to be added to insert in cache time
	items                    map[string]Item // TODO: switch to sync map?
	guard                    sync.RWMutex
}

// Item is part of cache. Cache is a map of items.
type Item struct {
	Value      interface{}
	Expiration int64
}

var onceCache sync.Once

// NEWCache is constructor for cache. Expiration in seconds.
func NEWCache(ttlSecs int64) *Cache {
	instance := new(Cache)

	onceCache.Do(func() {
		instance.defaultExpirationSeconds = ttlSecs
		instance.guard = sync.RWMutex{}
		instance.items = make(map[string]Item)
	})
	return instance
}

// Add - adds item to cache
func (cache *Cache) Add(key string, item *Item) {
	i := *item // creating local copy to minimize critical path length

	if item.Expiration > 0 {
		i.Expiration = item.Expiration
	} else {
		i.Expiration = time.Now().UnixNano() + (cache.defaultExpirationSeconds * 1000 * 1000 * 1000)
	}

	cache.guard.Lock()
	defer cache.guard.Unlock()

	cache.items[key] = i
}

// Get Method returns the value for the key
func (cache *Cache) Get(key string) (Item, bool) {
	return cache.items[key], true
}

// DeleteExpired - deletes expired entries from cache
func (cache *Cache) DeleteExpired() {
	now := time.Now().UnixNano()

	for k, v := range cache.items {
		if v.Expiration < now {
			cache.guard.Lock()
			delete(cache.items, k)
			cache.guard.Unlock()
		}
	}
}
