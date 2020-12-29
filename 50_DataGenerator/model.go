package datageneration

import (
	"math/rand"
	"time"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func newNumbers(length int, positive bool, min, max, howMany int) []int {
	rand.Seed(time.Now().UnixNano())
	randNumbers := make([]int, howMany)

	for i := 0; i < howMany; i++ {
		for {
			randNumbers[i] = randInt(min, max)
			if randNumbers[i] > min {
				break
			}
		}
	}
	return randNumbers
}

func newCharacters(length int, howMany int) []string {
	rand.Seed(time.Now().UnixNano())
	randChars := make([]string, 0)

	for h := 0; h < howMany; h++ {
		bytes := make([]byte, length)
		for k := range bytes {
			bytes[k] = byte(randInt(65, 90))
		}
		randChars = append(randChars, string(bytes))
	}
	return randChars
}

func newIDs(howMany int) []int {
	randNumbers := make([]int, howMany)

	for i := 0; i < howMany; i++ {
		randNumbers[i] = i + 1 // TODO: add random number generation
	}
	return randNumbers
}
