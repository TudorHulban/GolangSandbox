package main

import (
	"fmt"
	"testing"
)

func TestConnect2RDBMS(t *testing.T) {

	h := new(Credentials)
	h.Host = "192.168.1.21"
	h.User = "postgres"
	h.Pass = "admin"
	h.Port = "5432"
	h.DatabaseName = "tests"
	h.DatabaseTable = ""
	h.RDBMS = "postgres"

	c, err := Connect2RDBMS(h)
	fmt.Println("connection error: ", err)

	content, err := QueryPosts(c, 2, 5)

	fmt.Println("content: ", content)

	if err != nil {
		t.Error(err)
	}

	howmany, err := HowManyPosts(c)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("how many: ", howmany)
}
