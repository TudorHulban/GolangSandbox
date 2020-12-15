// Package janitor creates a janitor. Basically function that triggers periodically the run of another function.
package pkjanitor

import (
	"log"
	"time"
)

// Janitor is structure used for cleaning process.
type Janitor struct {
	RunInterval time.Duration //in seconds
	stop        chan bool
}

// NEWJanitor is constructor for janitor.
func New(pRunSecondsInterval int) *Janitor {
	instance := new(Janitor)
	instance.RunInterval = time.Duration(pRunSecondsInterval) * time.Second
	instance.stop = make(chan bool)
	return instance
}

// Clean - to be invoked for cleaning the cache. runs in separate goroutine
func (j *Janitor) Clean(pClean func()) {
	ticker := time.NewTicker(j.RunInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Janitor Triggered")
				pClean()
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
	log.Println("Janitor stopping...")
	j.stop <- true
}
