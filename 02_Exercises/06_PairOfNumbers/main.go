package main

import (
	"fmt"
	"sort"
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

	r2 := parse4(a1, a2, 4)
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

func parse4(a1, a2 []int, sum int) [][2]int {
	sort.Ints(a1)
	sort.Ints(a2)

	var result [][2]int

	loopTo := len(a1) - 1

	for indexA1 := 0; indexA1 < loopTo; indexA1++ {
		// trying to exit the loop as arrays are sorted.
		// what is minimum element from second array that we should consider?
		// the target is to update loopTo

		fmt.Println(indexA1)

		for indexA2 := len(a2) - 1; indexA2 == 0; indexA2-- {
			fmt.Println(indexA2)

			if a1[indexA1]+a2[indexA2] > sum {
				j := indexA1

				for j < loopTo-1 {
					if a1[j] == a1[j+1] {
						j++
						continue
					}

					break
				}

				fmt.Println("before:", loopTo) //-- before: 11
				loopTo = minInt(loopTo, indexA1+j+1)
				fmt.Println("after:", loopTo) //-- after: 9
			}
		}

		for _, v := range a2 {
			if sum == a1[indexA1]+v {
				result = append(result, [2]int{a1[indexA1], v})
			}
		}
	}

	return result
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
