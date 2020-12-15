/*
Package sequence provides a sequence to start from.
*/
package sequence

import (
	"sync"
)

// Sequence Structure holding the sequence value and locker for protection of operations.
type Sequence struct {
	m            sync.Mutex
	currentValue int64
}

// NewSeq Constructor for sequence.
func NewSeq(startValue int64) *Sequence {
	return &Sequence{
		currentValue: startValue,
	}
}

// GetCurrentValue returns current sequence value.
func (s *Sequence) GetCurrentValue() int64 {
	return s.currentValue
}

// IncrementValue returns sequence value incremented with passed value.
func (s *Sequence) IncrementValue(with int64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.currentValue = s.currentValue + with
}
