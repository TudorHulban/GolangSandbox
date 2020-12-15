/*
Package sequence provides a sequence to start from.
Tests pass race condition.
*/
package sequence

// Sequence Structure holding the sequence value and locker for protection of operations.
type Sequence struct {
	currentValue      int64
	chRequestSequence chan int64
	chReceiveSequence chan int64
	chStop            chan struct{}
}

/*
NewSeq provides sequence value on a read channel based on request on request channel.
Requests for new sequence and stop are empty struct.
*/
func NewSeq(startValue int64) *Sequence {
	result := &Sequence{
		currentValue:      startValue,
		chRequestSequence: make(chan int64), // channel to send the increment value
		chReceiveSequence: make(chan int64), // read the new current value.
		chStop:            make(chan struct{}),
	}

	go func(s *Sequence) {
		defer func() {
			close(s.chReceiveSequence)
			close(s.chReceiveSequence)
			close(s.chStop)
		}()

		for {
			select {
			case with := <-s.chRequestSequence:
				{
					s.currentValue = s.currentValue + with
					s.chReceiveSequence <- s.currentValue
				}
			case <-s.chStop:
				{
					break
				}
			}
		}
	}(result)
	return result
}

// GetCurrentValue returns current sequence value.
func (s *Sequence) GetCurrentValue() int64 {
	return s.currentValue
}

// IncrementValue returns sequence value incremented with passed value.
func (s *Sequence) IncrementValue(with int64) int64 {
	s.chRequestSequence <- with
	return <-s.chReceiveSequence
}

// Stop closes the channels opened.
func (s *Sequence) Stop() {
	s.chStop <- struct{}{}
}
