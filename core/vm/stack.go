package vm

import (
	"fmt"
	"github.com/holiman/uint256"
	"sync"
)

// Stack Each entry in a stack is a slice of 4 64-bit integers
type Stack struct {
	data []uint256.Int
}

// stackPool a concurrent collection of a set of Stacks
// Cache allocated but unused Stack objects for quick access
var pool = sync.Pool{
	New: func() interface{} {
		// Stack capacity is 512 bytes = (16 * 256 bits / 8)
		return &Stack{data: make([]uint256.Int, 0, 16)}
	},
}

// getStack gets a stack from the stack pool
func getStack() *Stack {
	return pool.Get().(*Stack)
}

// recycleStack clears stack contents then returns it to the stack pool
func recycleStack(stack *Stack) {
	stack.data = stack.data[:0] // Keep allocated array, but reset slice to zero length
	pool.Put(stack)
}

// push adds as item at the top of the stack
func (stack *Stack) push(entry uint256.Int) {
	stack.data = append(stack.data, entry)
}

// pop retrieves and removes the item at the top of the stack
func (stack *Stack) pop() (ret uint256.Int) {
	ret = stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)]
	return
}

// peek retrieves but does not remove the item at the top of the stack
func (stack *Stack) peek() (ret uint256.Int) {
	return stack.data[len(stack.data)-1]
}

func (stack *Stack) peekN(n int) (ret uint256.Int) {
	return stack.data[len(stack.data)-1-n]
}

func (stack *Stack) print() {
	fmt.Println("-------------Start stack dump-------------")
	if len(stack.data) == 0 {
		fmt.Println("Stack is empty")
	} else {
		for i, v := range stack.data {
			fmt.Printf("%-5d %s\n", i, v)
		}
	}
	fmt.Println("-------------End stack dump-------------")
}
