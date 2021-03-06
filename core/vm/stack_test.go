package vm

import (
	"Petra/common"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/holiman/uint256"
	"strings"
	"testing"
)

var stack *Stack
var stackError error
var popped []uint256.Int

func createStack() {
	stack = getStack()
	stackError = nil
	popped = nil
}

func stackShouldBeEmpty() error {
	if stack.size() != 0 {
		return fmt.Errorf("expected length 0, but found %d", stack.size())
	}
	return nil
}

func anEmptyStack() {
	createStack()
}

func peekIsCalled() {
	_, stackError = stack.peek()
}

func popIsCalled() {
	_, stackError = stack.pop()
}

func errorShouldBe(expectedError *godog.DocString) error {
	if !strings.EqualFold(expectedError.Content, stackError.Error()) {
		return errors.New("unexpected error")
	}
	return nil
}

func isPushed(data string) error {
	if data == "" {
		return nil
	}

	// Convert to numbers
	var items = strings.Split(data, ",")
	for _, item := range items {
		if bytes, err := hex.DecodeString(strings.TrimSpace(item)); err == nil {
			// Can replace below with: stack.push(*(&uint256.Int{}).SetBytes(bytes))
			value := uint256.Int{}
			value.SetBytes(bytes)
			stack.push(value)
		} else {
			return err
		}
	}
	return nil
}

func popIsCalledTimes(count int) error {
	for i := 0; i < count; i++ {
		var value uint256.Int
		if value, stackError = stack.pop(); stackError == nil {
			popped = append(popped, value)
		}
	}
	return nil
}

func poppedDataIs(data string) error {
	// Handle case when no items are popped
	if data == "" && popped == nil {
		return nil
	}

	// Items popped: check lengths are same
	expectedItems := strings.Split(data, ",")
	if len(expectedItems) != len(popped) {
		return fmt.Errorf("expected %d popped items, but actual is %d popped items",
			len(expectedItems), len(popped))
	}

	// Items popped: check values are the same
	for i, actual := range popped {
		bytes, _ := hex.DecodeString(expectedItems[i])
		expected := (&uint256.Int{}).SetBytes(bytes)
		if actual.Cmp(expected) != 0 {
			return fmt.Errorf("Testcase %d, expected  %x, got %x", i, expected, actual)
		}
	}

	return nil
}

func stackSizeIs(expectedSize int) error {
	actualSize := len(stack.data)
	if expectedSize == actualSize {
		return nil
	}
	return fmt.Errorf("expected size is %d, but actual is %d", expectedSize, actualSize)
}

func StackErrorShouldBe(expectedError string) error {
	expectedError = strings.TrimSpace(expectedError)

	if (expectedError != "" && stackError == nil) ||
		(expectedError == "" && stackError != nil) ||
		(expectedError != "" && stackError != nil && !strings.EqualFold(expectedError, stackError.Error())) {
		return fmt.Errorf("expected error: '%s', but got '%s'", expectedError, stackError)
	}
	return nil
}

func InitializeStackScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a stack is created$`, createStack)
	ctx.Step(`^stack should be empty$`, stackShouldBeEmpty)
	ctx.Step(`^an empty stack$`, anEmptyStack)
	ctx.Step(`^peek is called$`, peekIsCalled)
	ctx.Step(`^pop is called$`, popIsCalled)
	ctx.Step(`^error should be:$`, errorShouldBe)

	ctx.Step(`^error should be "([^"]*)"$`, StackErrorShouldBe)
	ctx.Step(`^"([^"]*)" is pushed$`, isPushed)
	ctx.Step(`^pop is called "([^"]*)" times$`, popIsCalledTimes)
	ctx.Step(`^popped data is "([^"]*)"$`, poppedDataIs)
	ctx.Step(`^stack size is "([^"]*)"$`, stackSizeIs)
}

func TestStack(t *testing.T) {
	suite := common.GetGodogTestSuite("features", "progress", "stack", "",
		InitializeStackScenario, nil)

	if suite.Run() != 0 {
		t.Fatal("Failed to run memory feature test")
	}
}
