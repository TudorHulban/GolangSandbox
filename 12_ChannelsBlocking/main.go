/*
	Example showing blocking on channels.
*/
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// schedule returns stop channel.
func schedule(what func(), interval time.Duration) chan bool {
	result := make(chan bool)

	go func() {
		for {
			what()

			// blocking. trigerring when there is send on channels.
			select {
			case <-time.After(interval):
				log.Println("interval finished, triggering")
			case v := <-result:
				log.Println("passed:", v)
				return
			}
		}
	}()
	return result
}

func cleanup() {
	log.Println("cleanup")
	time.Sleep(2 * time.Second) // delay exiting for catching previous sends on channel and logs.
}

func main() {
	// cancelling run channel
	cancelCh := make(chan os.Signal)
	signal.Notify(cancelCh, os.Interrupt, syscall.SIGTERM)

	// creating event and scheduling it
	event := func() {
		log.Println("EVENT Triggered")
	}
	stopSchedulingCh := schedule(event, 4*time.Second)

	go func() {
		<-cancelCh // this read from channel blocks until a value is placed on channel
		stopSchedulingCh <- true
		cleanup()
		os.Exit(1)
	}()

	log.Println("going to sleep now...")
	for {
		time.Sleep(5 * time.Second) // or runtime.Gosched()
		log.Println("awaken... going back to sleep...")
	}
}
