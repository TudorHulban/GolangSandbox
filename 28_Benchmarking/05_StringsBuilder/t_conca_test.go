package main

import (
	"testing"
)

// Benchmark_WSlice-8   	 6994339	       174.1 ns/op	      51 B/op	       2 allocs/op
func Benchmark_WSlice(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		concaSlice()
	}
}

// Benchmark_WBuilder-8   	 7072774	       171.5 ns/op	      56 B/op	       2 allocs/op
func Benchmark_WBuilder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		concaBuilder()
	}
}
