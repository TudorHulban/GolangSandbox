package main

import (
	"log"
	"reflect"
)

func main() {
	log.Println(reflect.TypeOf(retFunc(1, 2)))
	log.Println(reflect.TypeOf(retFunc(1, 2)(3, 4)))
	log.Println(retFunc(1, 2)(3, 4))
}

func retFunc(p1, p2 int) func(p3, p4 int) int {
	return func(p3, p4 int) int {
		return p1 + p2 + p3 + p4
	}
}
