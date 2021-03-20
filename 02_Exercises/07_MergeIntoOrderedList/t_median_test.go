package main

import (
	"testing"
)

func TestMedian(t *testing.T) {
	cases := []struct {
		description string
		a           []int
		expected    float64
	}{
		{"zero array", []int{0}, 0.00},
		{"one array", []int{1}, 1.00},
		{"two array", []int{2}, 2.00},
		{"a - odd length", []int{1, 2, 3}, 2.00},
		{"a - even length", []int{1, 2, 3, 4}, 2.5},
	}

	for _, tc := range cases {
		m := median(tc.a)
		if m != tc.expected {
			t.Log(tc.description, m)
			t.FailNow()
		}
	}
}
