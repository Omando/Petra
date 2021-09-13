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
