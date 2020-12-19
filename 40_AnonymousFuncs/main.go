package main

import (
	"log"
)

func main() {
	a1 := func(x string) {
		log.Println(x)
	}
	a1("xxx")

	a2 := func(x int) int {
		return x + 1
	}
	log.Println(a2(1))

	s := []string{"a", "b"}

	a3 := func(x []string) {
		x = append(x, "c")
		log.Println(x)
	}
	a3(s)
}
