package main

/*

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

Assumption: integers are positive.

From:
https://leetcode.com/problems/integer-to-roman/

*/

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1948

	deco := decomposeNumber(n)

	fmt.Println(deco)

	compo := recomposeNUmber(deco)

	fmt.Println(compo)
}

func values() map[int]string {
	return map[int]string{
		1:  "I",
		5:  "V",
		10: "X",
	}
}

func decomposeNumber(n int) []int {
	s := strconv.Itoa(n)

	res := make([]int, len(s))

	for i, j := len(s)-1, 0; i >= 0 && j < len(s); i, j = i-1, j+1 {
		v, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			fmt.Println("problems with ", s[i:i+1])
			return []int{}
		}

		res[j] = v
	}

	return res
}

func recomposeNUmber(n []int) int {
	var res int

	for i, v := range n {
		val := v * powersOfTen(i)
		fmt.Println(i, v, val)

		res = res + val
	}

	return res
}

func powersOfTen(p int) int {
	if p == 0 {
		return 1
	}

	res := 1

	i := 1
	for i <= p {
		i++
		res = res * 10
	}

	return res
}
