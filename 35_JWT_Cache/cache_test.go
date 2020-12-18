package cache

import (
	"fmt"
	"testing"
	"time"
)

const (
	cacheExpiration  = 2
	janitorRetrigger = 5
)

func TestCache(t *testing.T) {
	c := NEWCache(cacheExpiration)

	j := NEWJanitor(c, janitorRetrigger)
	j.Clean()

	c.Add("10", &Item{Value: "x", Expiration: 0})
	c.Add("11", &Item{Value: "y", Expiration: 0})
	c.Add("12", &Item{Value: "z", Expiration: 0})

	v1, isFound := c.Get("10")
	fmt.Println("get: ", v1, isFound)

	v2, isFound := c.Get("11")
	fmt.Println("get: ", v2, isFound)

	v3, isFound := c.Get("12")
	fmt.Println("get: ", v3, isFound)

	time.Sleep(time.Second * 6)
	fmt.Println("------------------")

	v1, isFound = c.Get("10")
	fmt.Println("get: ", v1, isFound)

	v2, isFound = c.Get("11")
	fmt.Println("get: ", v2, isFound)

	v3, isFound = c.Get("12")
	fmt.Println("get: ", v3, isFound)

	j.Stop()
	time.Sleep(time.Second * 3)
}
