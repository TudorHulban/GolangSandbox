package main

import (
	"log"
	"time"
)

func getStatus(pRequest request) {
	wait := time.NewTimer(20 * time.Second)
	<-wait.C

	log.Println("ready request: ", pRequest.id, readyRequest(pRequest))
}
