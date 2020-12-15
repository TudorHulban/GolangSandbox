package main

import (
	"fmt"
	"log"
)

func main() {
	a := []int{1, 2}
	log.Println("array:", a)

	b, _ := removeElem(a, 1)
	log.Println("array:", b)
}

func removeElem(a []int, index int) ([]int, error) {
	if index > len(a) {
		return []int{}, fmt.Errorf("maximum index is %v, passed value is %v", len(a), index)
	}

	return append(a[:index], a[index+1:]...), nil
}
