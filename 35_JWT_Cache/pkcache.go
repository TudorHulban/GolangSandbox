// Package cache - creates a single cache. Uses singleton for this.
package cache

import (
	"log"
	"sync"
	"time"
)

// Item is part of cache. Cache is a map of items.
type Item struct {
	Value      interface{}
	Expiration int64
}

// Cache contains the items.
type Cache struct {
	defaultExpirationSeconds int64           //amount of time to be added to insert in cache time
	items                    map[string]Item // TODO: switch to sync map?
	guard                    sync.RWMutex
}

// Janitor is used to delete expired items from cache.
type Janitor struct {
	ToClean     *Cache
	RunInterval time.Duration // in seconds
	stop        chan bool     // channel used to signalize exit
}

var onceCache sync.Once
var onceJanitor sync.Once

// NEWJanitor - constructor for janitor
func NEWJanitor(toClean *Cache, runSecondsInterval int) *Janitor {
	instance := new(Janitor)

	onceJanitor.Do(func() {
		instance.ToClean = toClean
		instance.RunInterval = time.Duration(runSecondsInterval) * time.Second
		instance.stop = make(chan bool)
	})
	return instance
}

// Clean - to be invoked for cleaning the cache. runs in separate goroutine
func (j *Janitor) Clean() {
	ticker := time.NewTicker(j.RunInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("deleting")
				j.ToClean.DeleteExpired()
			case <-j.stop:
				ticker.Stop()
				log.Println("ticker stopped")
				return
			}
		}
	}()
}

// Stop - stops janitor goroutine
func (j *Janitor) Stop() {
	j.stop <- true
}

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
	i := *item

	if item.Expiration > 0 {
		i.Expiration = item.Expiration
	} else {
		i.Expiration = time.Now().UnixNano() + (cache.defaultExpirationSeconds * 1000 * 1000 * 1000)
	}

	cache.guard.Lock()
	cache.items[key] = i
	cache.guard.Unlock()
}

// Get - returns the value for the key
func (cache *Cache) Get(key string) (Item, bool) {
	return cache.items[key]
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
