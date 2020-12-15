package main

// Actual sort
func bubblesort(theSlice []int) []int {
	if len(theSlice) < 2 {
		return theSlice
	}

	swapped := true
	j := 0

	for swapped {
		swapped = false
		j++

		for i := 0; i < (len(theSlice) - j); i++ {
			if theSlice[i+1] < theSlice[i] {
				theSlice[i+1], theSlice[i] = theSlice[i], theSlice[i+1]
				swapped = true
			}

		}
	}

	return theSlice
}
