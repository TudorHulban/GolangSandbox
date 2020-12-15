package datageneration
import (
	"math/rand"
	"time"
)

func randInt(pMin, pMax int) int {
	return pMin + rand.Intn(pMax-pMin)
}

func newNumbers(pLength int, pPositive bool, pMin, pMax, pHowMany int) []int {
	rand.Seed(time.Now().UnixNano())
	randNumbers := make([]int, pHowMany)

	for i := 0; i < pHowMany; i++ {
		for {
			randNumbers[i] = randInt(pMin, pMax)
			if randNumbers[i] > pMin {
				break
			}
		}
	}
	return randNumbers
}

func newCharacters(pLength int, pHowMany int) []string {
	rand.Seed(time.Now().UnixNano())
	randChars := make([]string, 0)

	for h := 0; h < pHowMany; h++ {
		bytes := make([]byte, pLength)
		for k := range bytes {
			bytes[k] = byte(randInt(65, 90))
		}
		randChars = append(randChars, string(bytes))
	}
	return randChars
}

func newIDs(pHowMany int) []int {
	randNumbers := make([]int, pHowMany)

	for i := 0; i < pHowMany; i++ {
		randNumbers[i] = i + 1
	}
	return randNumbers
}