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

	// Simple Async Subscriber
	subscriber, errSub := natsConn.Subscribe("foo1", func(m *stan.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	log.Println("errSub: ", errSub)

	// Unsubscribe
	subscriber.Unsubscribe()
}
