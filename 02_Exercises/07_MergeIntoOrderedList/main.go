package main

import (
	"sort"
)

/*

Given two sorted []int, merge them into a single sorted list.
The lists might have different length.

Resources: https://adaickalavan.github.io/portfolio/coding-questions-in-golang/#findPair&gsc.tab=0
item 4.

*/

var a1 = []int{1, 3, 5, 7}
var a2 = []int{2, 4, 6, 8, 9}
var result = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {

}

func order1(a1, a2 []int) []int {
	var result []int

	result = append(a1, a2...)
	sort.Ints(result)

	return result
}

func order2(a1, a2 []int) []int {
	var result []int
	var indexA2 int

	for _, v1 := range a1 {
		for indexA2 <= len(a2)-1 {
			if v1 < a2[indexA2] {
				break
			}

			result = append(result, a2[indexA2])
			indexA2++
		}

		result = append(result, v1)
	}

	return append(result, a2[indexA2:]...)
}
