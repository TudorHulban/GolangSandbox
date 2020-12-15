package main

import (
	"log"
	"os"

	stan "github.com/nats-io/stan.go"
)

func main() {
	natsConn, errCo := stan.Connect("test-cluster", "123", stan.NatsURL("nats://192.168.1.5:4222"))
	if errCo != nil {
		log.Println(errCo)
		os.Exit(1)
	}
	defer natsConn.Close()

	// Simple Synchronous Publisher
	errPub := natsConn.Publish("foo1", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
	log.Println("errPub: ", errPub)
}
