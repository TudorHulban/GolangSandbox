package main

import (
	"log"
	"reflect"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func main() {
	ptr := &User{
		Name: "John",
		Age:  32,
	}

	rtype := reflect.TypeOf(ptr)
	isPointer := strings.HasPrefix(rtype.String(), "*")
	log.Println(rtype, isPointer)
}
