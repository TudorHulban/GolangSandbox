package main

import (
	"fmt"
	"sort"
)

// Given a string s, find the first longest substring without repeating characters.
// Adapted as per https://leetcode.com/problems/longest-substring-without-repeating-characters/.

func main() {
	s := "abcd"

	fmt.Println(walkString(s))
}

func walkString(s string) string {
	substrings := getAllSubstrings(s)

	sort.Slice(substrings, func(i, j int) bool {
		if len(substrings[i]) > 0 && len(substrings[j]) > 0 {
			// orders alphabetically
			return (len(substrings[i]) > len(substrings[j])) && ([]rune(substrings[i])[0] < []rune(substrings[j])[0])
		}

		return (len(substrings[i]) > len(substrings[j]))
	})

	fmt.Println(substrings)

	for _, v := range substrings {
		if noRepeating(v) {
			return v
		}
	}

	return ""
}

func getAllSubstrings(s string) []string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			result = append(result, s[i:j])
		}
	}

	return result
}

func noRepeating(s string) bool {
	chars := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		chars[i] = s[i : i+1]
	}

	sort.Strings(chars)

	for j := 0; j < len(chars)-1; j++ {
		if chars[j] == chars[j+1] {
			// log.Printf("Char \"%s\" is repeating. \n", chars[j])
			return false
		}
	}

	return true
}
