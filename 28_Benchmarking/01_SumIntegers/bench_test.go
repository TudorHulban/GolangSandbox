package main

import (
	"testing"
)

var sample []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

func benchmarkSumLoop(pIntegers []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumLoop(pIntegers)
	}
}

func BenchmarkSumLoop(b *testing.B) {
	benchmarkSumLoop(sample, b)
}

func benchmarkSumRecursive(pIntegers []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumRecurs(pIntegers)
	}
}

func BenchmarkSumRecurs(b *testing.B) {
	benchmarkSumRecursive(sample, b)
}
