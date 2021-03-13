package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	fmt.Println(reverseArray(a))

	s := "abc"
	fmt.Println(reverseString(s))
}

func reverseArray(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return a
}

func reverseString(s string) string {
	a := []rune(s)

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return string(a)
}
