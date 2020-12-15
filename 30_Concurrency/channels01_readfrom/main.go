package main

import (
	"log"
)

func main() {
	n := 11
	chIntegers := make(chan int)

	go func() {
		defer close(chIntegers)

		for i := 0; i < n; i++ {
			chIntegers <- i
		}
	}()

	for {
		ev, isOpen := <-chIntegers
		if !isOpen {
			break
		}
		log.Println("ev:", ev)
	}
}
