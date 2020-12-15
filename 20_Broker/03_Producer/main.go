package main // needs refactoring with REST, Websockets and methods on request including marshalling

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const ip = "127.0.0.1"                       // from CLI - later
const producerCode = "P1"                    // from CLI - later
const defaultEventTTL = 600                  // seconds
const brokerAPIurl = "http://localhost:8080" // from CLI - later
const routeRegister = "/registerproducer"    // from CLI - later
const routeEvents = "/postevents"            // from CLI - later
const routeStatus = "/statusrequest"         // from CLI - later
const ramDisk = ""

func init() {
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

	evposted := false
	for {
		if !evposted {
			var ev request
			ev.code = producerCode
			ev.ttl = defaultEventTTL
			ev.id = 1
			ev.payload = []string{"https://golang.org", "https://jsonformatter.curiousconcept.com"}
			posted, errPost := postRequest(ev)

			go getStatus(ev)

			if errPost != nil {
				log.Println(errPost)
				time.Sleep(10 * time.Second)
			}
			log.Println("resp: ", posted)

			evposted = true
		}

		time.Sleep(10 * time.Second)
		log.Println("working...")
	}
}
