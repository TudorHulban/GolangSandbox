package main

import (
	"log"
	"math"
	"strconv"
)
import "time"

const noWorkers = 30
const noTasks = 100

type task struct {
	request string
}

// worker - primitive for performing tasks
func worker(pWorkerID int, pChJobs <-chan task, pChResults chan<- string) {
	log.Println("creating worker")
	for j := range pChJobs {
		log.Println("worker "+strconv.Itoa(pWorkerID)+"started  job", j)
		time.Sleep(1 * time.Millisecond)
		log.Println("worker "+strconv.Itoa(pWorkerID)+"finished job", j)
		pChResults <- "finished: " + strconv.FormatInt(time.Now().UnixNano(), 10) + " " + j.request
		log.Println("work sent by: ", pWorkerID, j.request)
	}
	log.Println("closing worker ", pWorkerID)
}

func main() {
	timeStart := time.Now()
	// buffered channel, but buffer is smaller in order to handle congestion, optim around 1/3 of number of workers?
	noBuffer := int(math.Min(float64(noTasks), 10.0))
	chJobs := make(chan task, noBuffer)
	chResults := make(chan string)
	results := []string{"\n"}

	// create workers
	for w := 0; w < noWorkers; w++ {
		go worker(w, chJobs, chResults)
		log.Println("created worker:", w)
	}

	go createWork(noTasks, chJobs, noBuffer)

	// read results
	for r := 0; r < noTasks; r++ {
		results = append(results, <-chResults+"\n")
	}
	close(chResults)
	log.Println(results)
	elapsed := time.Since(timeStart)
	log.Println("elapsed: ", elapsed)
}

func createWork(pNoTasks int, pChJobs chan task, pMaxConcTasks int) {
	var load int
	for pNoTasks > 0 {
		load = len(pChJobs)
		log.Println("creating work, remaining tasks: ", pNoTasks, "load:", load)
		if load <= pMaxConcTasks {
			log.Println("---------- channel load: ", load)
			req := "R" + strconv.Itoa(pNoTasks)
			pChJobs <- task{request: req}
			log.Println("request: ", req)
			pNoTasks--
		}
	}
	log.Println("no more work, remaining tasks: ", pNoTasks)
	close(pChJobs)
}
