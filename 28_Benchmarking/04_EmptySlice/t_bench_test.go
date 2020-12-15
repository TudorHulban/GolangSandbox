package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ex1 = []string{}
var ex2 = []string{"1", "2", "3", "4", "5"}

func Test_Nil(t *testing.T) {
	assert.True(t, byLength(ex1))
}

// Benchmark_Len-4   	1000000000	         0.767 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Len1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byLength(ex1)
	}
}

// Benchmark_Nil-4   	1000000000	         0.764 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Nil1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byNil(ex1)
	}
}

// Benchmark_Len-4   	1000000000	         0.767 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Len2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byLength(ex2)
	}
}

// Benchmark_Nil-4   	1000000000	         0.764 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Nil2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byNil(ex2)
	}
}
