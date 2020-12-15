package pkcache

import (
	"log"
	"reflect"
	"testing"
	"time"

	"../pkjanitor"
)

// Manual testing with go test -race.
func TestCacheAdd(t *testing.T) {
	i1 := Item{Value: "x", ExpirationTimestamp: 0}
	i2 := Item{Value: 1, ExpirationTimestamp: 0}

	c := New(2, &i1)
	log.Println("cache items type:", c.ItemType)

	e1 := c.Add("1", i1)
	log.Println("e1:", e1)
	e2 := c.Add("2", i2)
	log.Println("e2:", e2)

	if len(c.Items) != 1 {
		t.Error("cache items were not added correctly")
	}
	for k, v := range c.Items {
		log.Println("got: ", k, v, reflect.TypeOf(v.Value))
	}
}

// Manual testing with go test -race.
func TestJanitor(t *testing.T) {
	i1 := Item{Value: "x", ExpirationTimestamp: 0}
	i2 := Item{Value: "y", ExpirationTimestamp: 0}
	i3 := Item{Value: "z", ExpirationTimestamp: 0}

	c := New(2, &i1)

	j := pkjanitor.New(2)
	j.Clean(c.DeleteExpired)

	c.Add("1", i1)
	c.Add("2", i2)
	c.Add("3", i3)

	log.Println("items in cache:", len(c.Items))

	for k, v := range c.Items {
		log.Println("got: ", k, v, reflect.TypeOf(v.Value))
	}
	if len(c.Items) != 3 {
		t.Error("cache items were not added correctly")
	}
	time.Sleep(time.Second * 2)

	j.Stop()
	time.Sleep(time.Second * 1)

	if len(c.Items) != 0 {
		t.Error("cache items were not fully cleaned by janitor")
	}
}
