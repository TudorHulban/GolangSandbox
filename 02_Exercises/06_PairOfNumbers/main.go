package main

import (
	"fmt"
)

/*

Find every pair of numbers (inclusive of duplicates) from two integer arrays
(i.e. one number from each array), whose sum equals a given number.

ex.:
INPUT:
arr1 := []int{-1, -2, 4, 4, -2, -6, 5, -7}
arr2 := []int{0, 6, 3, 4, 0}

for 4 the output should be:

OUTPUT:
[[4 0] [4 0] [-2 6] [-2 6] [4 0] [4 0]]

*/

func main() {
	a1 := []int{-1, -2, 4, 4, -2, -6, 5, -7, 4, 4, 10, 12}
	a2 := []int{0, 6, 3, 4, 0, 7}

	r1 := parse1(a1, a2, 4)
	fmt.Println(r1)

	r2 := parse2(a1, a2, 4)
	fmt.Println(r2)
}

func parse1(a1, a2 []int, sum int) [][2]int {
	var result [][2]int

	for _, v1 := range a1 {
		for _, v2 := range a2 {
			if sum == v1+v2 {
				result = append(result, [2]int{v1, v2})
			}
		}
	}

	return result
}

func parse2(a1, a2 []int, n int) [][]int {
	a1Map := make(map[int]int)
	for _, v1 := range a1 {
		a1Map[v1] = a1Map[v1] + 1
	}

	result := [][]int{}

	for _, v2 := range a2 {
		temp := n - v2

		if count, ok := a1Map[temp]; ok {
			for j := count; j > 0; j-- {
				result = append(result, []int{temp, v2})
			}
		}
	}

	return result
}

func parse3(a1, a2 []int, n int) [][2]int {
	a1Map := make(map[int]int)
	for _, v1 := range a1 {
		a1Map[v1] = a1Map[v1] + 1
	}

	result := [][2]int{}

	for _, v2 := range a2 {
		temp := n - v2

		if count, ok := a1Map[temp]; ok {
			for j := count; j > 0; j-- {
				result = append(result, [2]int{temp, v2})
			}
		}
	}

	return result
}
