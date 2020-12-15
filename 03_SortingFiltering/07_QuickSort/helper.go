package main

import (
	"math/rand"
	"time"
)

// generateSlice Generates a slice of passed size and filled with random numbers.
func generateSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())

	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(99) - rand.Intn(99)
	}
	return result
}
