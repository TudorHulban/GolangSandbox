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

func Benchmark_Order1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		order1(a1, a2)
	}
}
