package vm

import "fmt"

type MemorySizeError struct {
	message string
}

func (mse *MemorySizeError) Error() string {
	return fmt.Sprintf("Memory size error: %s", mse.message)
}
