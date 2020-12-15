package cache

import (
	"fmt"
	"testing"
	"time"
)

const cacheExpiration = 2
const janitorRetrigger = 5

var c *Cache
var j *Janitor

func init() {
	c = NEWCache(cacheExpiration)

	j = NEWJanitor(c, janitorRetrigger)
	j.Clean()
}

func TestCache(t *testing.T) {
	i1 := Item{Value: "x", Expiration: 0}
	i2 := Item{Value: "y", Expiration: 0}
	i3 := Item{Value: "z", Expiration: 0}

	c.Add("10", &i1)
	c.Add("11", &i2)
	c.Add("12", &i3)

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
