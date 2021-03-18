package main

// Find all letter combinations from a word.
// Number of letters to be deleted can be added in breaking condition of recursion.

// Result:
/*

[a b c d]
[ab ac ad bc bd cd]
[abc abd acd bcd]
[abcd]

*/

import (
	"fmt"
	"sort"
)

func main() {
	s := "abcd"
	res := make(map[string]bool)

	deleteOne(s, &res)

	order := make([][]string, len(s)+1)

	for k := range res {
		order[len(k)] = append(order[len(k)], k)
	}

	order[len(s)] = append(order[len(s)], s)

	for k := range order {
		if k == 0 {
			continue
		}

		sort.Strings(order[k])
		fmt.Println(order[k])
	}
}

// deleteOne Deletes recursively a letter from given string 
// and adds combinations to result pointer.
func deleteOne(s string, res *map[string]bool) {
	if len(s) == 1 {
		(*res)[s] = true
		return
	}

	for i := 0; i < len(s); i++ {
		b := s[:i] + s[i+1:]

		(*res)[b] = true

		deleteOne(b, res)
	}
}
