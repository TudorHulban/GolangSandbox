package main

import "fmt"

// Generic Structure to contain priority data that is related to heap.
type Generic struct {
	Priority int
}

// Priority Structure to keep heap.
type Priority struct {
	State []*Generic
}

// NewPriority Constructor for priority queue kind of type.
func NewPriority() *Priority {
	return &Priority{
		State: []*Generic{},
	}
}

// Insert Method adds data to the queue.
func (m *Priority) Insert(data *Generic) *Priority {
	m.State = append(m.State, data)

	if len(m.State) == 1 {
		return m
	}

	m.heapifyMax(len(m.State) - 1)

	fmt.Println(m.Values())

	return m
}

// Extract Method extracts top priority data from queue.
func (m *Priority) Extract() *Generic {
	l := len(m.State)

	if l == 0 {
		return nil
	}

	if l == 1 {
		result := (*m).State[0]

		m.State = []*Generic{}
		return result
	}

	result := (*m).State[0]
	fmt.Println("extracted:", result.Priority)

	// swap root with last element
	m.State[0] = m.State[l-1]
	// chop array
	m.State = m.State[:l-1]

	m.heapifyDown(0)
	fmt.Println(m.Values())

	return result
}

// heapifyDown Helper which roots element with top priority.
func (m *Priority) heapifyDown(index int) {
	lastIndex := len(m.State) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = 1
		} else if m.State[l].Priority > m.State[r].Priority {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.State[index].Priority < m.State[childToCompare].Priority {
			m.swap(index, childToCompare)

			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Values Method can be used to nicely print priority values.
func (m *Priority) Values() []int {
	result := make([]int, len(m.State))

	for i, v := range m.State {
		result[i] = v.Priority
	}

	return result
}

func (m *Priority) heapifyMax(index int) {
	for m.State[parent(index)].Priority < m.State[index].Priority {
		m.swap(parent(index), index)

		index = parent(index)
	}
}

func (m *Priority) swap(i1, i2 int) {
	m.State[i1], m.State[i2] = m.State[i2], m.State[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {

}
