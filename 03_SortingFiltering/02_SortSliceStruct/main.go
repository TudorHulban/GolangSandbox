package main

import (
	"log"
	"sort"
)

type task struct {
	id   int
	name string
}

// no pointer needed due to slice mechanics
func sortAsc(tasks []task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].id < tasks[j].id
	})
}

func sortName(tasks []task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].name < tasks[j].name
	})
}

func main() {
	tasks := []task{{id: 11, name: "t11"}, {id: 2, name: "t2"}, {id: 3, name: "t4"}, {id: 3, name: "t3"}}
	log.Println("Raw:", tasks)

	sortAsc(tasks)
	sortName(tasks)
	log.Println("Processed:", tasks)
}
