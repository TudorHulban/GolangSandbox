package cache

import (
	"log"
	"sync"
	"time"
)

// Janitor is used to delete expired items from cache.
type Janitor struct {
	ToClean     *Cache
	RunInterval time.Duration // in seconds
	stop        chan bool     // channel used to signalize exit
}

var onceJanitor sync.Once

// NEWJanitor - constructor for janitor
// TODO: refactor to flyweight pattern.
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
