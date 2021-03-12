package main

import (
	"fmt"
	"sort"
)

/*

Given two sorted []int, merge them into a single sorted list.
The lists might have different length.

Resources: https://adaickalavan.github.io/portfolio/coding-questions-in-golang/#findPair&gsc.tab=0
item 4.

*/

var a1 = []int{1, 3, 5, 7}
var a2 = []int{2, 4, 6, 8}
var result = []int{1, 2, 3, 4}

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

	for indexA1 := range a1 {
		if a1[indexA1] == a2[indexA1] {
			result = append(result, a1[indexA1], a2[indexA1])
			continue
		}

		if a1[indexA1] < a2[indexA1] {
			result = append(result, a1[indexA1])

			indexA2 = indexA1
			j := indexA1 + 1

			for j < len(a1)-1 {
				fmt.Println("j:", j, a1[j], a2[indexA2])

				if a1[j] == a2[indexA2] {
					result = append(result, a1[j], a2[indexA2])
					break
				}

				if a1[j] < a2[indexA2] {
					result = append(result, a1[j])
					j++
					continue
				}

				result = append(result, a2[indexA2])
				indexA2++
				j++
			}

			if j == len(a1) {
				result = append(result, a2[indexA2:]...)
				break
			}
		}
	}

	return result
}
