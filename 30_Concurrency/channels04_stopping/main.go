package main

import (
	"log"
	"sync"
	"time"
)

var poolWorkers int
var mutex = &sync.Mutex{}

type workTask struct {
	begin int
	end   int
}

func main() {
	chIntegers := make(chan int)
	chPause := make(chan bool)
	chWork := make(chan bool)
	chQuit := make(chan bool)

	for {
		ev, isOpen := <-chIntegers
		if !isOpen {
			break
		}
		log.Println("ev:", ev)
	}
}

// https://stackoverflow.com/questions/38798863/golang-pause-a-loop-in-a-goroutine-with-channels
func doControlledWork(pChResults chan int, pChPause, pChWork, pChQuit chan bool, pTask workTask) {
	for {
		select {
		case <-pChPause:
			log.Println("Pause")
		case ev := <-pChQuit:
			{
				if !ev {
					log.Println("Quit")
					return
				}
			}
		default:
			{
				//do work
			}
		}
	}
}

func doQuit(pChAction chan bool, pAfterSeconds int) {
	time.Sleep(time.Duration(pAfterSeconds) * time.Second)
	pChAction <- false
}
