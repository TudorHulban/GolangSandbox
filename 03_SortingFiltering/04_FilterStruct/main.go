package main

// file details filtering a slice of structs using a filtering function.

import (
	"log"
)

type person struct {
	age  int64
	name string
}

func main() {
	persons := []person{
		{name: "x", age: 4},
		{name: "y", age: 8},
		{name: "z", age: 11},
	}

	log.Println("Persons:", persons)
	log.Println("Maximum 7", filterPersons(persons, filter7max))
	log.Println("Minimum 7", filterPersons(persons, filter7min))
}

func filterPersons(persons []person, f func(i int64) bool) []person {
	result := []person{}

	for _, v := range persons {
		if f(v.age) {
			result = append(result, v)
		}
	}

	return result
}

func filter7max(i int64) bool {
	return i < 7
}

func filter7min(i int64) bool {
	return i >= 7
}
