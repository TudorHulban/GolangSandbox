package main

import (
	"testing"
)

func TestLogic(t *testing.T) {
	cases := []struct {
		description string
		a1          []int
		a2          []int
		expected    []int
	}{
		{"minimum arrays 1", []int{1}, []int{2}, []int{1, 2}},
		{"minimum arrays 2", []int{2}, []int{1}, []int{1, 2}},
		{"diff sizes 1", []int{1}, []int{3, 4}, []int{1, 3, 4}},
		{"diff sizes 2", []int{3, 4}, []int{1}, []int{1, 3, 4}},
		{"diff positions", []int{1, 3, 5, 7}, []int{2, 4, 6, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tc := range cases {
		result := order2(tc.a1, tc.a2)

		if len(result) != len(tc.expected) {
			t.FailNow()
		}

		for i := range result {
			if result[i] != tc.expected[i] {
				t.Log(i, result[i], tc.expected[i])
				t.FailNow()
			}
		}
	}
}

// Benchmark_Order1-8   	 5858282	       197 ns/op	      96 B/op	       2 allocs/op
func Benchmark_Order1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		order1(a1, a2)
	}
}

// Benchmark_Order2-8   	 7419392	       159 ns/op	     120 B/op	       4 allocs/op
func Benchmark_Order2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		order2(a1, a2)
	}
}
