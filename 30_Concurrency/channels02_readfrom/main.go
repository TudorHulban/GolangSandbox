package main

import (
	"log"
	"sync"
)

var poolWorkers int
var mutex = &sync.Mutex{}

func main() {
	chIntegers := make(chan int)

	go doWork(chIntegers, 0, 10)
	poolWorkers++
	go doWork(chIntegers, 11, 10)
	poolWorkers++

	for {
		ev, isOpen := <-chIntegers
		if !isOpen {
			break
		}
		log.Println("ev:", ev)
	}
}

func doWork(pCh chan int, pFrom int, pSteps int) {
	for i := pFrom; i <= pFrom+pSteps; i++ {
		pCh <- i
	}
	mutex.Lock()
	poolWorkers--

	if poolWorkers < 1 {
		log.Println("Closing work.")
		close(pCh)
	}
	mutex.Unlock()
}
