package singleresource

import (
	"log"
	"sync"
)

// Resource is generic type of resource.
type Resource struct {
	Res1   string
	Comms  chan string
	ChDone chan struct{}
}

var instance Resource
var instanceCreated bool
var mtx sync.RWMutex

// GetInstance provides singleton resource.
func GetInstance(pSalt string) (*Resource, error) {
	mtx.Lock()
	if !instanceCreated {
		instanceCreated = true
		newResource("xxxxx", pSalt)
	}
	mtx.Unlock()
	return &instance, nil
}

func newResource(pProp1, pProp2 string) {
	instance.Res1 = pProp1 + pProp2
	instance.Comms = make(chan string)
	instance.ChDone = make(chan struct{})

	go func() {
		defer close(instance.ChDone)
		defer close(instance.Comms)
		for {
			select {
			case msg := <-instance.Comms:
				{
					log.Println(msg)
				}
			case <-instance.ChDone:
				{
					log.Println("closing logger.")
					return
				}
			}
		}
	}()
}
