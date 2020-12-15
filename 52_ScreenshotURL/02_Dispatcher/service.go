package main

import (
	"log"
)

func (p *pool) AddWorker(pWorker worker) {
	p.workers = append(p.workers, pWorker)
}

func (p *pool) UpdateWorker(pWorker worker) {

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
