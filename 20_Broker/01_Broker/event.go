package main

import (
	"time"
)

const defaultCheckAfterSeconds = 60

type event struct {
	ID             int64
	requestID      int64
	producedTime   int64
	ttl            int64 // unix time
	reassignTime   int64 // unix time after which event can be reassigned to consumer
	consumeStart   int64
	consumeEnd     int64
	checkAfter     int64 // unix time after to check result
	producedByCode string
	consumedByCode string
	request        string
	response       string
}

// newEvent - TTL in seconds
func newEvent(theProducer producer, theRequest string, theRequestID, ttl, checkAfterSecs int64) event {
	now := time.Now().UnixNano()

	return event{
		requestID:      theRequestID,
		producedTime:   now,
		ttl:            now + (time.Duration(ttl) * time.Second).Nanoseconds(),
		checkAfter:     now + (time.Duration(checkAfterSecs) * time.Second).Nanoseconds(),
		producedByCode: theProducer.Code,
		request:        theRequest,
	}
}
