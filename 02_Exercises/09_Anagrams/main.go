package main

/*

Check if passed strings are anagram of each other.
Anagram of string is another string with same letters
but in different order.

*/

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "abc"
	s2 := "bac"

	fmt.Println(checkAnagramStrngs(s1, s2))

	a1 := []int{1, 2, 3, 1}
	a2 := []int{1, 1, 3, 2}

	fmt.Println(checkAnagramArrays(a1, a2))
}

func checkAnagramArrays(a1, a2 []int) bool {
	return reflect.DeepEqual(identifyElemsInt(a1), identifyElemsInt(a2))
}

func identifyElemsInt(a []int) map[int]int {
	result := make(map[int]int)

	for _, v := range a {
		if _, exists := result[v]; exists {
			result[v]++
			continue
		}

		result[v] = 1
	}

	return result
}

func checkAnagramStrngs(s1, s2 string) bool {
	return reflect.DeepEqual(identifyElemsRune(s1), identifyElemsRune(s2))
}

func identifyElemsRune(s string) map[rune]int {
	result := make(map[rune]int)

	for _, v := range []rune(s) {
		if _, exists := result[v]; exists {
			result[v]++
			continue
		}

		result[v] = 1
	}

	return result
}
