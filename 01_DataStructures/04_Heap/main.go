package main

// Generic Structure to contain priority data that is related to heap.
type Generic struct {
	Priority int
}

// Priority Structure to keep heap.
type Priority struct {
	State []*Generic
}

func NewPriority() *Priority {
	return &Priority{
		State: []*Generic{},
	}
}

func (m *Priority) Insert(data *Generic) *Priority {
	m.State = append(m.State, data)

	if len(m.State) == 1 {
		return m
	}

	m.heapifyMax(len(m.State) - 1)
	return m
}

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
