package main

import (
	"log"
	"math"
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
	timeStart := time.Now()
	dispatchWork(chIntegers, 0, 1000, 20)
	for {
		ev, isOpen := <-chIntegers
		if !isOpen {
			break
		}
		log.Println("ev:", ev)
	}
	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timeStart)
	log.Println("Elapsed", elapsed)
}

func dispatchWork(pChResults chan int, pFrom, pSteps, pNoWorkers int) {
	work := distributeWork(pFrom, pSteps, pNoWorkers)

	for _, task := range work {
		go doWork(pChResults, task)
		poolWorkers++
	}
}

func distributeWork(pFrom, pSteps, pNoWorkers int) []workTask {
	var remain, lastStep int
	load := int(math.Floor(float64(pSteps) / float64(pNoWorkers)))

	result := []workTask{}
	for i := 0; i < pNoWorkers; i++ {
		task := workTask{}
		switch i {
		case 0:
			{
				task.begin = 0
				lastStep = load
			}
		default:
			{
				task.begin = lastStep + 1
				lastStep = lastStep + load
			}
		}
		task.end = lastStep
		remain = remain + load
		result = append(result, task)
	}
	result[pNoWorkers-1].end = result[pNoWorkers-1].end + (pSteps - remain)

	for k, v := range result {
		log.Println("Task ", k+1, " details:", v)
	}
	return result
}

func doWork(pChResults chan int, pTask workTask) {
	for i := pTask.begin; i <= pTask.end; i++ {
		pChResults <- i
	}
	mutex.Lock()
	poolWorkers--

	if poolWorkers < 1 {
		log.Println("Closing work.")
		close(pChResults)
	}
	mutex.Unlock()
}
