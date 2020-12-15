package main

import (
	"log"
	"sync"

	singleresource "../pakresource"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		r1, _ := singleresource.GetInstance("1")
		r1.Comms <- "r1"
		log.Println("r1:", r1.Res1)
		wg.Done()
	}()

	go func() {
		r2, _ := singleresource.GetInstance("1")
		r2.Comms <- "r2"
		log.Println("r2:", r2.Res1)
		wg.Done()
	}()

	r3, _ := singleresource.GetInstance("2")
	r3.Comms <- "r3"
	log.Println("r3:", r3.Res1)

	wg.Wait()

	r3.ChDone <- struct{}{}
	log.Println("comms channel closed")
}
