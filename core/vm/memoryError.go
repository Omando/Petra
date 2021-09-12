package vm

import "fmt"

type MemorySizeError struct {
	message string
}

func (mse *MemorySizeError) Error() string {
	return fmt.Sprintf("Memory size error: %s", mse.message)
}

type MemoryOffsetError struct {
	size       uint
	offset     uint
	dataLength int
}

func (moe *MemoryOffsetError) Error() string {
	return fmt.Sprintf("offset %d + size %d is > data length %d", moe.offset, moe.size, moe.dataLength)
}
