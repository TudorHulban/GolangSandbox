package main

import (
	"encoding/json"
	"net/http"
)

// producer - quota not implemented
type producer struct {
	Code          string `json:"code"`
	IP            string `json:"ip"`
	quotaInterval int64
	quotaAmount   int64
	quotaStart    int64
}

type consumer struct {
	Code   string `json:"code"`
	Socket string `json:"socket"`
}

type request struct {
	ProducerCode string   `json:"code"`
	ID           int64    `json:"id"`
	TTL          int64    `json:"ttl"` // seconds
	Payload      []string `json:"payload"`
}

type readyevent struct {
	ConsumerCode string `json:"code"`
	EVentID      int64  `json:"id"`
}

type requeststatus struct {
	ProducerCode string `json:"code"`
	RequestID    int64  `json:"id"`
}

type task struct {
	ConsumerCode string `json:"code"`
}

func decodeRequest(r *http.Request, f func() interface{}) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	ev := f()

	errDecode := decoder.Decode(&ev)
	if errDecode != nil {
		return nil, errDecode
	}
	return ev, nil
}

func registerProducer() interface{} {
	instance := producer{}
	return interface{}(instance) // refactor, remove interim variable?
}

func reqisterConsumer() interface{} {
	instance := consumer{}
	return interface{}(instance)
}

func requestEvent() interface{} {
	return new(request)
}

func requestTask() interface{} {
	return new(task)
}

func readyEvent() interface{} {
	return new(readyevent)
}

func requestStatus() interface{} {
	return new(requeststatus)
}
