package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const socket = "127.0.0.1:8081"              // from CLI - later
const consumerCode = "C1"                    // from CLI - later
const brokerAPIurl = "http://localhost:8080" // from CLI - later
const routeRegister = "/registerconsumer"    // from CLI - later
const routeConsume = "/consumevent"          // from CLI - later
const ramDisk = ""
const cmd = "firefox -headless --screenshot "

func init() {
	errInit := runCmd("bash", "cleanpng.sh")
	if errInit != nil && errInit.Error() != "exit status 1" {
		log.Println("clean images error:", errInit)
	}
	registered, errRegister := registerWithBroker()
	log.Println("broker: ", registered, errRegister)
}

func cleanup() {
	log.Println("cleanup")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		event, errConsume := consumeEvent()
		if errConsume != nil {
			log.Println("error consume: ", errConsume)
			time.Sleep(10 * time.Second)
		}
		log.Println("consume: ", event)
		if event != "no work" {
			doWork(event)
		} else {
			time.Sleep(10 * time.Second)
		}
		log.Println("working...")
	}
}
