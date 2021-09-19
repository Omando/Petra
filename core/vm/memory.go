package vm

import (
	"fmt"
	"github.com/holiman/uint256"
)

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

// GetCopy returns data starting at given offset and up to offset+size as a new slice
func (m *Memory) GetCopy(offset, size uint) (dataCopy []byte, err error) {
	// Check inputs
	if size == 0 {
		return nil, &MemorySizeError{"size is zero"}
	}

	if int(offset+size) > len(m.data) {
		return nil, &MemoryOffsetError{size, offset, len(m.data)}
	}

	// Allocate a new slice of the given size, and copy data starting at
	// the given offset and upto offset+size
	dataCopy = make([]byte, int(size))
	copy(dataCopy, m.data[offset:offset+size])
	return dataCopy, nil
}

// GetPtr returns the offset + size
func (m *Memory) GetPtr(offset, size uint) ([]byte, error) {
	// Check inputs
	if size == 0 {
		return nil, &MemorySizeError{"size is zero"}
	}

	if int(offset+size) > len(m.data) {
		return nil, &MemoryOffsetError{size, offset, len(m.data)}
	}

	// Return data starting at the given offset
	return m.data[offset : offset+size], nil
}

// Set sets offset + size to value
func (m *Memory) Set(offset, size uint64, value []byte) error {
	// Check inputs
	if size == 0 {
		return &MemorySizeError{"size is zero"}
	}

	// Length of store must be at least offset+size
	if int(offset+size) > len(m.data) {
		return &MemoryOffsetError{uint(size), uint(offset), len(m.data)}
	}

	copy(m.data[offset:offset+size], value)
	return nil
}

// Set32 sets the 32 bytes starting at offset to the value of val,
// left-padded with zeroes to 32 bytes
func (m *Memory) Set32(offset uint64, val *uint256.Int) error {
	// Length of store must be at least offset+size
	if int(offset+32) > len(m.data) {
		return &MemoryOffsetError{32, uint(offset), len(m.data)}
	}

	copy(m.data[offset:offset+32], []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return nil
}

// Dump dumps the content of Memory
func (m *Memory) Dump() {
	fmt.Printf("### Total %d bytes in memory ###\n", len(m.data))
	if len(m.data) > 0 {
		address := 0
		for i := 0; i+32 <= len(m.data); i += 32 {
			fmt.Printf("%04d: %x\n", address, m.data[i:i+32])
			address++
		}
	}

	fmt.Println("###")
}
