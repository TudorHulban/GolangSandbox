package main

import (
	"testing"
	"time"
)

func TestEventCreation(t *testing.T) {
	p1 := producer{IP: "127.0.0.1"}
	q1 := newQueue(100)
	e1 := newEvent(p1, "req 1", 1, 5, 0)

	q1.addEvent(e1)
	if len(q1.events) != 1 || q1.lastID != 1 {
		t.Error("event not created")
	}
}

func TestGetEvent(t *testing.T) {
	p1 := producer{Code: "P1", IP: "127.0.0.1"}
	c1 := consumer{Code: "C1", Socket: "127.0.0.2:8081"}
	q1 := newQueue(100)
	e1 := newEvent(p1, "req 1", 1, 5, 0)
	e2 := newEvent(p1, "req 2", 2, 5, 0)

	q1.addEvent(e1)
	q1.addEvent(e2)
	e3 := q1.getOldestEvent(&c1)
	if e3.request != e1.request {
		t.Error("wrong event fetched")
	}
	time.Sleep(10 * time.Millisecond)
	e4 := q1.getOldestEvent(&c1)
	if e4.request == e1.request {
		t.Error("event fetched again")
	}
}

func TestClean(t *testing.T) {
	p1 := producer{IP: "127.0.0.1"}
	c1 := consumer{Socket: "127.0.0.2:8081"}
	q1 := newQueue(100)
	e1 := newEvent(p1, "req 1", 1, 0, 0)
	e2 := newEvent(p1, "req 2", 2, 50, 0)
	e3 := newEvent(p1, "req 3", 3, 0, 0)
	e4 := newEvent(p1, "req 4", 4, 50, 0)

	go q1.addEvent(e1)
	go q1.addEvent(e2)
	go q1.addEvent(e3)
	go q1.addEvent(e4)
	time.Sleep(100 * time.Millisecond)
	q1.clean()

	e5 := q1.getOldestEvent(&c1)
	if e5.request != e2.request {
		t.Error("wrong event fetched")
	}
	if len(q1.events) != 2 {
		t.Error("queue not cleaned", len(q1.events))
	}
}
