package main

import "math/rand"

// With reccurence.
// Due to slice mechanics the passed slice is updated within the function.
func quicksort(sliceToSort []int) []int {
	if len(sliceToSort) < 2 {
		return sliceToSort
	}

	left, right := 0, len(sliceToSort)-1
	pivot := rand.Int() % len(sliceToSort)

	sliceToSort[pivot], sliceToSort[right] = sliceToSort[right], sliceToSort[pivot]

	for i := range sliceToSort {
		if sliceToSort[i] < sliceToSort[right] {
			sliceToSort[left], sliceToSort[i] = sliceToSort[i], sliceToSort[left]
			left++
		}
	}

	sliceToSort[left], sliceToSort[right] = sliceToSort[right], sliceToSort[left]

	quicksort(sliceToSort[:left])
	quicksort(sliceToSort[left+1:])

	return sliceToSort
}
