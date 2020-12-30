package sequence

/*
File tests sequence generation.
Passes with go test -race.
*/

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	startFrom     = 1
	incrementWith = 1
)

func Test1_CurrentValue(t *testing.T) {
	s := NewSeq(startFrom)

	assert.EqualValues(t, startFrom, s.GetCurrentValue())
}

func Test2_GetNextValue(t *testing.T) {
	s := NewSeq(startFrom)

	s.IncrementValue(incrementWith)
	assert.EqualValues(t, startFrom+incrementWith, s.GetCurrentValue())
}

func Test3_Race(t *testing.T) {
	s := NewSeq(startFrom)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		s.IncrementValue(incrementWith)
		wg.Done()
	}()

	go func() {
		s.IncrementValue(incrementWith)
		wg.Done()
	}()

	go func() {
		s.IncrementValue(0)
		wg.Done()
	}()
	wg.Wait()

	assert.EqualValues(t, startFrom+2*incrementWith, s.GetCurrentValue())
}

// go test -bench=.
// Benchmark_WMutex-8   	70632560	        16.5 ns/op	       0 B/op	       0 allocs/op
func Benchmark_WMutex(b *testing.B) {
	s := NewSeq(77)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IncrementValue(1)
	}
}
