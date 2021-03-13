package main

import (
	"testing"
)

func TestLogic(t *testing.T) {
	a := order2(a1, a2)

	t.Log("order:", a)

	if len(result) != len(a) {
		t.FailNow()
	}

	for i := range result {
		if result[i] != a[i] {
			t.Log(i, result[i], a[i])
			t.FailNow()
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
