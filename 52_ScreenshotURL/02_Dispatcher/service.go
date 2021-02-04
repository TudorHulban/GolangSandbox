package main

import (
	"log"
)

func (p *pool) AddWorker(w worker) {
	p.workers = append(p.workers, w)
}

func (p *pool) UpdateWorker(w worker) {

}

func (p *pool) RequestWorker() *worker {
	// needs round robin as load is the same (1)

	result := &p.workers[0]
	log.Println("selected: ", result)
	return result
}

func (p *pool) QuitWork() {
	log.Println("...quit service")
}
