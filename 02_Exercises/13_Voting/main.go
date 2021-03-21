package main

// "aaabbbcddd" a -> 3 b -> 3 c -1 d -> 3

// => Winner is : a  (max votes, single person)
// => loser is : b (min votes, single person)
// => no winner

// Length of the string is N (i.e. n different votes)
// "aaaabbbcddd" -> Winner is a, Loser is c
// "aaabbbcddd" -> No winner
// "aaabcd" -> Winner is a
// "abcd" -> No Winner

import (
	"errors"
	"fmt"
	"sort"
)

// MapData Exported for future use.
type MapData struct {
	Key   rune
	Value int
}

func main() {
	s := "abcd" // every letter is a vote. if maximum number votes twice, no winner

	m := identifyEl(s)
	fmt.Println(wins(m))
}

func identifyEl(s string) map[rune]int {
	result := make(map[rune]int)

	for _, v := range s {
		if _, exists := result[v]; exists {
			result[v]++
			continue
		}

		result[v] = 1
	}

	return result
}

func orderMapData(m map[rune]int) []MapData {
	if len(m) == 0 {
		return []MapData{}
	}

	var result []MapData

	for k, v := range m {
		result = append(result, MapData{
			Key:   k,
			Value: v,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	return result
}

// returns winner, loser, error
func wins(m map[rune]int) (string, string, error) {
	if len(m) == 0 {
		return "", "", errors.New("no votes")
	}

	var a []int
	for _, v := range m {
		a = append(a, v)
	}

	if len(a) == 1 {
		return "", "", errors.New("no winner or loser")
	}

	sort.Ints(a)

	// check if occurences are the same
	if a[0] == a[len(a)-1] {
		return "", "", errors.New("no winner or loser")
	}

	// find runes in map
	var winner, loser rune

	for k, r := range m {
		if r == a[len(a)-1] {
			winner = k
		}

		if r == a[0] {
			loser = k
		}

		if winner != 0 && loser != 0 {
			break
		}
	}

	if a[len(a)-1] == a[len(a)-2] {
		return "", string(loser), errors.New("but no winner")
	}

	if a[0] == a[1] {
		return string(winner), "", errors.New("and no loser")
	}

	return string(winner), string(loser), nil
}
