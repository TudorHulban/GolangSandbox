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
func doControlledWork(chResults chan int, chPause, chWork, chQuit chan bool, task workTask) {
	for {
		select {
		case <-chPause:
			log.Println("Pause")

		case ev := <-chQuit:
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

func doQuit(chAction chan bool, afterSeconds int) {
	time.Sleep(time.Duration(afterSeconds) * time.Second)

	chAction <- false
}
