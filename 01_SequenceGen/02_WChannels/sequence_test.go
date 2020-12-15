package sequence

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
	defer s.Stop()

	assert.EqualValues(t, startFrom, s.GetCurrentValue())
}

func Test2_GetNextValue(t *testing.T) {
	s := NewSeq(startFrom)
	defer s.Stop()

	assert.EqualValues(t, startFrom+incrementWith, s.IncrementValue(incrementWith))
}

func Test3_Race(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	s := NewSeq(1)
	defer s.Stop()

	go func() {
		s.IncrementValue(incrementWith)
		wg.Done()
	}()

	go func() {
		s.IncrementValue(incrementWith)
		wg.Done()
	}()

	go func() {
		s.IncrementValue(incrementWith)
		wg.Done()
	}()
	wg.Wait()

	assert.EqualValues(t, startFrom+3*incrementWith, s.GetCurrentValue())
}

// go test -bench=.
// Benchmark_WChannels-4   	  904759	      2004 ns/op	       0 B/op	       0 allocs/op
func Benchmark_WChannels(b *testing.B) {
	s := NewSeq(77)
	defer s.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.IncrementValue(1)
	}
}
