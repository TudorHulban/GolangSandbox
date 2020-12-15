package main

import (
	"log"
)

func main() {
	var y = func(x string) {
		log.Println(x)
	}
	y("xxx")

	var z = func(x int) int {
		return x + 1
	}
	log.Println(z(1))

	s := []string{"a", "b"}

	var superSlice = func(x []string) {
		x = append(x, "c")
		log.Println(x)
	}
	superSlice(s)   // copy slice and do something with it
	log.Println(s)
}
