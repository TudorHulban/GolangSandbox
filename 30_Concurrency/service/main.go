package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

type request struct {
	id int64
}

type resource struct {
	id int64
}

type service struct {
	mutex     sync.RWMutex
	cache     map[request]resource
	chResults chan int64
}

const noWorkers = 10

func main() {
	chIntegers := make(chan int64)
	serv := newService(chIntegers)

	/*
		req1 := request{id: 1}
		r1 := serv.createResource(&req1)
		req2 := request{id: 2}
		serv.createResource(&req2)
		req3 := request{id: 3}

		serv.listResources()
		//serv.cleanResource(r1)
		//serv.cleanCache()
		serv.listResources()

		r3, _ := serv.provideResource(&req1)
		log.Println("r1: ", *r1)
		log.Println("r3: ", *r3)

		r4, err := serv.provideResource(&req3)
		if err != nil {
			r4 = serv.createResource(&req3)
		}
		log.Println("r4: ", *r4)
	*/

	go dispatchWork(noWorkers, chIntegers, *serv)

	for {
		ev, isOpen := <-chIntegers
		if !isOpen {
			break
		}
		log.Println("ev:", ev)
	}
}

func dispatchWork(pNoWorkers int, pChComm chan int64, pService service) {
	var wg sync.WaitGroup

	for i := 0; i < pNoWorkers; i++ {
		wg.Add(1)
		go doWork(pChComm, &pService, &wg)
	}
	wg.Wait()
	close(pService.chResults)
}

func doWork(pChResults chan int64, pService *service, wg *sync.WaitGroup) {
	req := request{id: time.Now().UnixNano()}
	res := pService.createResource(&req)

	pChResults <- (*res).id
	wg.Done()
}

func newService(pChResults chan int64) *service {
	instance := new(service)
	instance.cache = make(map[request]resource, 0)
	instance.chResults = pChResults
	return instance
}

func (s *service) createResource(pRequest *request) *resource {
	instance := new(resource)
	instance.id = time.Now().UnixNano()

	s.mutex.Lock()
	s.cache[*pRequest] = *instance
	s.mutex.Unlock()
	log.Println("res: ", instance.id)
	return instance
}

func (s *service) provideResource(pRequest *request) (*resource, error) {
	result := new(resource)
	err := errors.New("resource not found")
	for k, v := range s.cache {
		if k == *pRequest {
			result = &v
			err = nil
			log.Println("resource identified: ", v.id)
			break
		}
	}
	log.Println("----------- Provide resource:", err)
	return result, err
}

func (s *service) cleanResource(pResource *resource) {
	pResource.Close()

	for k, v := range s.cache {
		if v == *pResource {
			delete(s.cache, k)
			log.Println("resource deleted: ", pResource.id)
			break
		}
	}
}

func (s *service) cleanCache() {
	for k, res := range s.cache {
		res.Close()
		delete(s.cache, k)
		log.Println("resource deleted: ", res)
	}
}

func (s *service) listResources() {
	log.Println("number cached resources: ", len(s.cache))
	for k, res := range s.cache {
		log.Println("List resources: ", k, res.id)
	}
}

func newResource() *resource {
	instance := new(resource)
	return instance
}

func (r *resource) Close() error {
	log.Println("resource closed: ", *r)
	return nil
}
