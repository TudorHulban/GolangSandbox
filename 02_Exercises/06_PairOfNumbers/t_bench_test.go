package main

import (
	"testing"
)

const sum = 4

var a1 = []int{-1, -2, 4, 4, -2, -6, 5, -7}
var a2 = []int{0, 6, 3, 4, 0}

// Benchmark_Parse1-8   	 6006733	       204 ns/op	     240 B/op	       4 allocs/op
func Benchmark_Parse1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		parse1(a1, a2, sum)
	}
}

// Benchmark_Parse2-8   	 1484632	       814 ns/op	     464 B/op	      10 allocs/op
func Benchmark_Parse2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		parse2(a1, a2, sum)
	}
}

// Benchmark_Parse3-8   	 2591080	       465 ns/op	     240 B/op	       4 allocs/op
func Benchmark_Parse3(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		parse3(a1, a2, sum)
	}
}

// Benchmark_Parse4-8   	 4306646	       288 ns/op	     112 B/op	       4 allocs/op
func Benchmark_Parse4(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		parse4(a1, a2, sum)
	}
}
