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

// MapData Exported for future use.
type MapString struct {
	Key   string
	Value int
}

func main() {
	s := "abcdadba" // every letter is a vote. if maximum number votes twice, no winner

	m1 := identifyEl(s)
	fmt.Println(wins(m1))

	m2 := identifyString(s)
	fmt.Println(winsString(orderMapString(m2)))
}

func identifyString(s string) map[string]int {
	result := make(map[string]int)

	for i := 0; i < len(s); i++ {
		c := s[i : i+1]

		if _, exists := result[c]; exists {
			result[c]++
			continue
		}

		result[c] = 1
	}

	return result
}

func orderMapString(m map[string]int) []MapString {
	result := []MapString{}

	if len(m) == 0 {
		return result
	}

	for k, v := range m {
		result = append(result, MapString{
			Key:   k,
			Value: v,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	return result
}

func winsString(m []MapString) (string, string, error) {
	if len(m) == 0 {
		return "", "", errors.New("no votes")
	}

	if len(m) == 1 {
		return "", "", errors.New("no winner or loser")
	}

	if m[0].Value == m[len(m)-1].Value {
		return "", "", errors.New("no winner or loser")
	}

	if m[len(m)-1] == m[len(m)-2] {
		return "", m[0].Key, errors.New("but no winner")
	}

	if m[0].Value == m[1].Value {
		return m[len(m)-1].Key, "", errors.New("and no loser")
	}

	return m[len(m)-1].Key, m[0].Key, nil
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
