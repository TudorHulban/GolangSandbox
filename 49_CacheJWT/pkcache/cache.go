// Package cache - creates a single cache of items of same type.
package pkcache

import (
	"errors"
	"log"
	"reflect"
	"sync"
	"time"
)

// Item is contained in cache. Cache is a map of items of same type. Using interface{} as workaround.
type Item struct {
	Value               interface{} // same type
	ExpirationTimestamp int64       // unix time, nano
}

// Cache is map of items. Default Time To Live in seconds.
type Cache struct {
	DefaultTTL int64 // time to live to be added when inserting in cache.
	Items      map[string]Item
	ItemType   reflect.Type
	Guard      sync.RWMutex
}

// New is constructor for cache. Parameter is Default Time To Live in seconds after now and item for defining cache item type.
func New(pTTL int64, pItem *Item) *Cache {
	instance := new(Cache)
	instance.DefaultTTL = pTTL
	instance.Guard = sync.RWMutex{}
	instance.Items = make(map[string]Item)
	instance.ItemType = reflect.TypeOf((*pItem).Value)
	return instance
}

// Add adds item to cache. Item passed by value for concurrency reasons.
func (c *Cache) Add(pKey string, pItem Item) error {
	if reflect.TypeOf(pItem.Value) != c.ItemType {
		return errors.New("item type different than cached items")
	}
	if !(pItem.ExpirationTimestamp > 0) {
		pItem.ExpirationTimestamp = time.Now().UnixNano() + (c.DefaultTTL * 1000000000)
	}
	c.Guard.Lock()
	c.Items[pKey] = pItem
	c.Guard.Unlock()
	return nil
}

// Get - returns the value for the key
func (c *Cache) Get(pStringToken string) (Item, bool) {
	i, isOK := c.Items[pStringToken]

	return i, isOK
}

// DeleteExpired deletes expired entries from cache.
func (c *Cache) DeleteExpired() {
	now := time.Now().UnixNano()
	log.Println("number of cache items: ", len(c.Items), " time: ", now)

	for k, v := range c.Items {
		log.Println("k: ", k)
		log.Println("v: ", v.ExpirationTimestamp, time.Unix(0, v.ExpirationTimestamp))
		log.Println("item expiration: ", v.ExpirationTimestamp)
		log.Println("delta delete: ", v.ExpirationTimestamp)

		if v.ExpirationTimestamp < now {
			c.Guard.Lock()
			log.Println("deleting: ", c.Items[k])
			delete(c.Items, k)
			c.Guard.Unlock()
		}
	}
}
