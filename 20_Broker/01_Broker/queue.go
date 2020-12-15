package main

import (
	"log"
	"sort"
	"sync"
	"time"
)

type queue struct {
	events             []event
	ttl                int64 // default TTL
	lastID             int64
	intervalCheckAfter int64                // nanoseconds after which producer can check results, varies per load, inject in event creation
	producers          map[string]*producer // registered producers
	consumers          map[string]*consumer // registered consumers
	mux                sync.Mutex
}

func newQueue(defaultTTL int64) *queue {
	return &queue{
		ttl:       defaultTTL,
		producers: make(map[string]*producer),
		consumers: make(map[string]*consumer),
	}
}

// addEvent Method adds event to queue and returns the added event ID.
func (q *queue) addEvent(e event) int64 {
	// passed by copy event will use to add to queue
	e.ID = q.newID()

	q.events = append(q.events, e) // TODO: any race conditions?
	return e.ID
}

// getOldestEvent Method gets oldest event by creation time for consumer.
func (q *queue) getOldestEvent(c *consumer) *event {
	if len(q.events) == 0 {
		return nil
	}

	buffer := filterWith(filterAssignedToConsumer, q.events)
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].(event).producedTime < buffer[j].(event).producedTime
	})

	if len(buffer) == 0 {
		return nil
	}

	log.Println(buffer)

	result := q.getEventByID(buffer[0].(event).ID)
	result.consumedByCode = c.Code
	result.consumeStart = time.Now().UnixNano()
	return result
}

func filterWith(f func(e event) bool, theSlice interface{}) []interface{} {
	result := []interface{}{}

	for _, v := range theSlice.([]event) {
		if f(v) && !(v.consumeEnd > 0) {
			result = append(result, v)
		}
	}
	return result
}

func filterAssignedToConsumer(e event) bool {
	return len(e.consumedByCode) == 0
}

// getRequestStatus - TODO: in work
func (q *queue) getRequestStatus(reqID int64) bool {
	if len(q.events) == 0 {
		return true
	}

	buffer := filterWith(filterConsumed, q.events)
	if len(buffer) == 0 {
		return true
	}

	return false
}

func (q *queue) newID() int64 {
	var result int64

	q.mux.Lock()
	q.lastID++
	result = q.lastID
	q.mux.Unlock()

	return result
}

func (q *queue) clean() {
	q.sortEvents()

	var lastOldestEvent int64
	for k, v := range q.events {
		if v.ttl > time.Now().UnixNano() {
			break
		}
		lastOldestEvent = int64(k)
	}
	q.events = q.events[lastOldestEvent+1:]
}

// sortEvents - by TTL
func (q *queue) sortEvents() {
	sort.Slice(q.events, func(i, j int) bool {
		return q.events[i].ttl < q.events[j].ttl
	})
}

func (q *queue) getEventByID(id int64) *event {
	var ix int64

	for k, v := range q.events {
		if v.ID == id {
			ix = int64(k)
			break
		}
	}
	return &q.events[ix]
}

func (q *queue) addConsumer(c *consumer) {
	q.consumers[c.Code] = c
}

func (q *queue) addProducer(p *producer) {
	q.producers[p.Code] = p
}

func (q *queue) isRegisteredProducer(code string) (bool, *producer) {
	for _, v := range q.producers {
		if v.Code == code {
			return true, q.producers[code]
		}
	}
	return false, nil
}

func (q *queue) isRegisteredConsumer(code string) (bool, *consumer) {
	for _, v := range q.consumers {
		if v.Code == code {
			return true, q.consumers[code]
		}
	}
	return false, nil
}

func (q *queue) addPayload(req request) {
	for _, v := range req.Payload {
		q.addEvent(event{
			ID:             q.newID(),
			requestID:      req.ID,
			ttl:            req.TTL,
			producedByCode: req.ProducerCode,
			request:        v,
		})
	}
}

// deleteSliceID - not used
func deleteSliceID(slice *[]interface{}, id int64) {
	(*slice)[id] = (*slice)[len(*slice)-1] // order is not preserved
	(*slice) = (*slice)[:len((*slice))-1]
}

func filterConsumed(e event) bool {
	return !(e.consumeEnd > 0)
}
