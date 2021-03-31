package main

/*

Check if passed string is a palindrome.
A palindrome is a word, number, phrase, or other
sequence of characters which reads the same backward as forward,
such as madam or racecar.

ex. `tenet`, `aa`

*/

import (
	"fmt"
)

func main() {
	s1 := "tenet"
	s2 := "aa"

	fmt.Println(check(s1))
	fmt.Println(check(s2))
}

func check(s string) bool {
	l := len(s)
	half := l / 2

	h1 := s[:half]
	h2 := s[half+1:]

	// did not use runes and reverse string
	for i := 0; i < half-1; i++ {
		if h1[i] != h2[len(h2)-1-i] {
			return false
		}
	}

	return true
}
