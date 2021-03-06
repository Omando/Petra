package vm

import (
	"errors"
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

func (stack *Stack) isEmpty() bool {
	return len(stack.data) == 0
}

func (stack *Stack) size() int {
	return len(stack.data)
}

// push adds as item at the top of the stack
func (stack *Stack) push(entry uint256.Int) {
	stack.data = append(stack.data, entry)
}

// pop retrieves and removes the item at the top of the stack
func (stack *Stack) pop() (ret uint256.Int, err error) {
	if stack.isEmpty() {
		return uint256.Int{}, errors.New("stack is empty")
	}

	// Retrieve and remove item from the end of the slice
	index := len(stack.data) - 1
	ret = stack.data[index]
	stack.data = stack.data[:index]
	return ret, nil
}

// peek retrieves but does not remove the item at the top of the stack
func (stack *Stack) peek() (*uint256.Int, error) {
	if stack.isEmpty() {
		return &uint256.Int{}, errors.New("stack is empty")
	}

	return &stack.data[len(stack.data)-1], nil
}

func (stack *Stack) peekN(n int) (uint256.Int, error) {
	if stack.isEmpty() {
		return uint256.Int{}, errors.New("stack is empty")
	}

	return stack.data[len(stack.data)-1-n], nil
}

func (stack *Stack) print() {
	fmt.Println("------------- Start stack dump -------------")
	if len(stack.data) == 0 {
		fmt.Println("Stack is empty")
	} else {
		for i, v := range stack.data {
			fmt.Printf("%-5d %s\n", i, v.String())
		}
	}
	fmt.Println("------------- End stack dump -------------")
}
