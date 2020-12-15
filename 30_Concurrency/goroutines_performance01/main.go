package main

import (
	"log"
	"time"
)

type memoZ func(int) int

func main() {
	log.Println("start")
	upTo := 18
	benchmark(upTo, fibo, "recursive")
	benchmark(upTo, fibogo, "goroutine")
	benchmark(upTo, fiboZ, "memoize")
	log.Println("\n")
	upTo = 19
	benchmark(upTo, fibo, "recursive")
	benchmark(upTo, fibogo, "goroutine")
	benchmark(upTo, fiboZ, "memoize")
}

func benchmark(pValue int, f func(x int) int, pNote string) {
	start := time.Now()
	result := f(pValue)
	elapsed := time.Since(start)
	log.Println("value: ", pValue, "fibo:", result, "elapsed", elapsed, "("+pNote+")")
}

func workermemo(pZ memoZ) memoZ {
	cache := make(map[int]int)

	return func(pKey int) int {
		value, isFound := cache[pKey]
		if isFound {
			return value
		}
		cache[pKey] = pZ(pKey)
		return cache[pKey]
	}
}

func fiboZ(x int) int {
	return workermemo(fibo)(x)
}

func fibo(x int) int {
	switch x {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		return fibo(x-2) + fibo(x-1)
	}
}

func worker(pChValues chan int, pCounter int) {
	n1 := 0
	n2 := 1
	for i := 0; i < pCounter; i++ {
		pChValues <- n1
		x := n1
		n1 = n2
		n2 = x + n2
	}
	close(pChValues)
}

func fibogo(x int) int {
	x = x + 1
	chValues := make(chan int)

	go worker(chValues, x)
	var result int

	for num := range chValues { // used only for picking values from channel
		result = num
	}
	return result
}
