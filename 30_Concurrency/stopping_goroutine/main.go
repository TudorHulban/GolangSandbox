package main

import (
	"sync"
	"time"
)
import "fmt"

func main() {
	cStop := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)
	go work(&wg, cStop)

	time.Sleep(2 * time.Second)
	fmt.Println("-----------------")
	cStop <- false
	cStop <- true
	wg.Wait()

	fmt.Println("stopping")
}

func work(pWG *sync.WaitGroup, cSignal chan bool) {

	for {
		fmt.Println("routine work", time.Now().Unix())
		time.Sleep(100 * time.Millisecond)

		select {
		case event := <-cSignal:
			{
				if event {
					fmt.Println("STOP message received")
					pWG.Done()
					return
				} else {
					fmt.Println("event received: ", event)
				}

			}
		default:
			fmt.Println("no message received")
		}
	}
}
