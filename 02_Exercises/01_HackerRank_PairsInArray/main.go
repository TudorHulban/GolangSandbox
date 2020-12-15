package main

import (
	"fmt"
	"log"
)

/*
Code gets number of pairs from an array.
The number of occurences for a pair is configurable.
*/

func main() {
	data := []int{1, 2, 1, 1, 3, 2, 4, 2, 2}

	log.Println(getPairs(2, data))
}

// occurencesForPair Is the number of occurences that constitutes a pair.
func getPairs(occurencesForPair int, theArray []int) (int, error) {
	if occurencesForPair < 1 {
		return 0, fmt.Errorf("passed occurences for pair number value: %v smaller than one occurence", occurencesForPair)
	}

	log.Println("")
	log.Printf("Computing number of pairs for %v in pair.", occurencesForPair)

	// defining new map to hold the number of occurences per each value
	// key is value, value is number of occurences
	data := make(map[int]int)

	for _, value := range theArray {
		_, exists := data[value]
		if exists {
			data[value]++ // increment map value for found key.
			continue
		}

		// key does not exist therefore adding new key with one occurence.
		data[value] = 1
	}

	log.Println("data:", data)

	// extract number of pairs
	// defining holder for the pairs
	// key is the number for which occurences are calculated and value the number of occurences
	thePairs := make(map[int]int)

	for k, theValue := range data {
		numberPairs := theValue / occurencesForPair
		log.Println("key:", k, "occurences:", theValue, "number of pairs:", numberPairs)

		if numberPairs >= 1 {
			log.Println("key:", k, "adding for number of pairs:", numberPairs)
			thePairs[k] = numberPairs
		}
	}

	log.Println("pairs:", thePairs)

	var result int
	for _, n := range thePairs {
		result = result + n
	}

	return result, nil
}
