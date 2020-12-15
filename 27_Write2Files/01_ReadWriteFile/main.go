package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const fileName = "log_data"
const noGRoutines = 20

func main() {
	content, errRead := ReadFile(fileName)
	if errRead != nil {
		log.Println("Error reading file.")
	}

	for _, row := range content {
		log.Println("row:", row)
	}

	startTime := time.Now()
	os.Remove(fileName)

	var wg sync.WaitGroup
	var invokedNo int

	//closure, references wg, invokedNo
	log2file := func(msg string) {
		defer wg.Done()

		invokedNo++
		isNow := strconv.FormatInt(time.Now().UnixNano(), 10)
		LogToFile(fileName, msg+" - "+strconv.Itoa(invokedNo), isNow+" - ")
	}

	wg.Add(noGRoutines)

	for i := 0; i < noGRoutines; i++ {
		go log2file("GRoutine: " + strconv.Itoa(i+1))
	}
	wg.Wait()
	log.Println(time.Since(startTime))
}
