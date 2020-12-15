package main

import (
	"strings"
)

// JustWords Structure consolidating sorting methods.
// Implements the sort interface.
type JustWords []string

func (w JustWords) Len() int {
	return len(w)
}

func (w JustWords) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w JustWords) Less(i, j int) bool {
	switch sortMethod {
	case sortLength:
		{
			// sort by length
			return len(w[i]) < len(w[j])
		}
	case sortAlphabetic:
		{
			// alphabetic sort
			return strings.ToLower(w[i]) < strings.ToLower(w[j])
		}
	}
	return false
}
