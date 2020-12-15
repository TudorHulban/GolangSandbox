package main

import (
	"testing"
)

// Benchmark_aFunction1-4   	1000000000	         0.765 ns/op	       0 B/op	       0 allocs/op
func Benchmark_aFunction1(b *testing.B) {
	v := aStruct1{
		state: "x",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aFunction1(v)
	}
}

// Benchmark_aFunction2-4   	239866978	         4.95 ns/op	       0 B/op	       0 allocs/op
func Benchmark_aFunction2(b *testing.B) {
	theState := "x"
	v := aStruct2{
		state: &theState,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aFunction2(v)
	}
}

// Benchmark_aFunction3-4   	1000000000	         0.767 ns/op	       0 B/op	       0 allocs/op
func Benchmark_aFunction3(b *testing.B) {
	v := aStruct1{
		state: "x",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aFunction3(&v)
	}
}

// Benchmark_aFunction4-4   	1000000000	         0.761 ns/op	       0 B/op	       0 allocs/op
func Benchmark_aFunction4(b *testing.B) {
	v := aStruct1{
		state: "x",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aFunction4(&v)
	}
}

// Benchmark_aFunction5-4   	584905333	         1.91 ns/op	       0 B/op	       0 allocs/op
func Benchmark_aFunction5(b *testing.B) {
	theState := "x"
	v := aStruct2{
		state: &theState,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aFunction5(&v)
	}
}

// Benchmark_aFunction6-4   	1000000000	         0.765 ns/op	       0 B/op	       0 allocs/op
func Benchmark_aFunction6(b *testing.B) {
	v := "x"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		aFunction6(v)
	}
}
