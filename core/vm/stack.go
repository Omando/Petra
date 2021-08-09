package vm

import (
	"github.com/holiman/uint256"
	"sync"
)

// Stack Each entry in a stack is an array of 4 64-bit integers
type Stack struct {
	data[] uint256.Int
}

var pool = sync.Pool{
	New: func() interface{} {
		// Stack capacity is 512 bytes = (16 * 256 bits / 8)
		return &Stack{data: make([]uint256.Int,0,16)}
	},
}

func create() *Stack {
	return pool.Get().(*Stack);
}



