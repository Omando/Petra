package vm

type Memory struct {
	data        []byte
	lastGasCost uint64
}

// NewMemory return a new memory model
func NewMemory() *Memory {
	return &Memory{
		data:        make([]byte, 0),
		lastGasCost: 0,
	}
}

// Length returns the length of the backing slice
func (m *Memory) Length() int {
	return len(m.data)
}

// ResizeIfLessThan increases the Memory to newSize if current size is less than the new size
func (m *Memory) ResizeIfLessThan(newSize int) {
	if len(m.data) < newSize {
		sizeDiff := newSize - len(m.data)
		m.data = append(m.data, make([]byte, sizeDiff)...)
	}
}
