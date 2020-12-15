package main

// file detailing singleton definition and use.

import (
	"log"
	"sync"
)

type resource struct {
	resourceLocation string
}

var instance *resource

func getInstance(salt string) *resource {
	newResource := func(resourceLoc string) *resource {
		return &resource{
			resourceLocation: resourceLoc,
		}
	}

	if instance == nil {
		instance = newResource("xxxxx")
	}

	return &resource{
		resourceLocation: instance.resourceLocation + salt,
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		r1 := getInstance("1")
		log.Println("r1:", *r1)
		wg.Done()
	}()

	r2 := getInstance("2")
	log.Println("r2:", *r2)
	wg.Wait()
}
