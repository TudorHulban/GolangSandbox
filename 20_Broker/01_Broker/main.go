package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

type jsonAnswer struct {
	Response string `json:"response"`
}

const port = "8080"
const defaultQueueTTL = 600 // seconds
const urlRegisterProducer = "/registerproducer"
const urlRegisterConsumer = "/registerconsumer"
const urlPostEvents = "/postevents"
const urlConsumeEvent = "/consumevent" // switch to GET - later
const urlReadyEventID = "/readyevent"
const urlStatusRequestID = "/statusrequest" // switch to GET - later

var q queue = *newQueue(defaultQueueTTL) // one queue only for now

func main() {
	http.HandleFunc(urlRegisterProducer, handlerRegisterProducer)
	http.HandleFunc(urlRegisterConsumer, handlerRegisterConsumer)
	http.HandleFunc(urlPostEvents, handlerPostEvent)
	http.HandleFunc(urlConsumeEvent, handlerConsumeEvent)
	http.HandleFunc(urlReadyEventID, handlerReadyEvID)
	http.HandleFunc(urlStatusRequestID, handlerStatusRequestID)

	http.ListenAndServe((fmt.Sprintf(":%v", port)), nil)
}

func handlerRegisterProducer(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, registerProducer)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())
		return
	}
	var theStruct producer
	mapstructure.Decode(req, &theStruct)
	q.addProducer(&theStruct)
	log.Println("registered producers:", len(q.producers), q.producers)
	sendResponse(w, "ok")
}

func handlerRegisterConsumer(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, reqisterConsumer)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())
		return
	}
	var theStruct consumer
	mapstructure.Decode(req, &theStruct)
	q.addConsumer(&theStruct)
	log.Println("registered consumers:", len(q.consumers), q.consumers)
	sendResponse(w, "ok")
}

func handlerPostEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("-- post event")

	req, errDecode := decodeRequest(r, requestEvent)
	if errDecode != nil {
		log.Println("event decode error: ", errDecode)
		sendResponse(w, errDecode.Error())
		return
	}

	var payload request
	mapstructure.Decode(req, &payload)

	exists, _ := q.isRegisteredProducer(payload.ProducerCode)
	if !exists {
		sendResponse(w, "not registered")
		return
	}

	q.addPayload(payload)
	log.Println("registered events:", len(q.events))
	log.Println(q.events)
	sendResponse(w, "ok")
}

func handlerConsumeEvent(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, requestTask)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())
		return
	}

	var t task
	mapstructure.Decode(req, &t)
	log.Println("requesting work:", t.ConsumerCode)

	exists, consumer := q.isRegisteredConsumer(t.ConsumerCode)
	if !exists {
		sendResponse(w, "not registered")
		return
	}

	event := q.getOldestEvent(consumer)
	if event == nil {
		sendResponse(w, "no work")
		return
	}

	log.Println("consumed event:", event)
	log.Println(q.events)
	sendResponse(w, event.request)
}

// handlerReadyEvID - invoked by consumer
func handlerReadyEvID(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, readyEvent)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())
		return
	}
	var theStruct readyevent
	mapstructure.Decode(req, &theStruct)
	log.Println("ready event:", theStruct.EVentID)

	event := q.getEventByID(theStruct.EVentID)
	if event == nil {
		sendResponse(w, "event does not exist in queue")
		return
	}
	if event.consumedByCode != theStruct.ConsumerCode {
		sendResponse(w, "consumed wrong event id")
		return
	}
	event.consumeEnd = time.Now().UnixNano()
	log.Println("consumed event:", event)
	log.Println(q.events)
	sendResponse(w, event.request)
}

// handlerStatusRequestID - invoked by producer
func handlerStatusRequestID(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, requestStatus)
	if errDecode != nil {
		log.Println("event decode error: ", errDecode)
		sendResponse(w, errDecode.Error())
		return
	}
	var theStruct readyevent
	mapstructure.Decode(req, &theStruct)
	log.Println("ready event:", theStruct.EVentID)

	event := q.getEventByID(theStruct.EVentID)
	if event == nil {
		sendResponse(w, "event does not exist in queue")
		return
	}
	if event.consumedByCode != theStruct.ConsumerCode {
		sendResponse(w, "consumed wrong event id")
		return
	}
	event.consumeEnd = time.Now().UnixNano()
	log.Println("consumed event:", event)
	log.Println(q.events)
	sendResponse(w, event.request)
}

func sendResponse(w http.ResponseWriter, pResponse string) {
	response := jsonAnswer{Response: pResponse}
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)
}
