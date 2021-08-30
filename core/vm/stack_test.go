package vm

import (
	"Petra/common"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/holiman/uint256"
	"os"
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

func StackErrorShouldBe(expectedError string) error {

}

func isPushed(data string) error {

}

func popIsCalledTimes(count int) error {

}

func poppedDataIs(data string) error {

}

func stackSizeIs(expectedSize int) error {
}

func InitializeScenario(ctx *godog.ScenarioContext) {
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

func TestMain(m *testing.M) {
	status := common.GetGodogTestSuite("features", "progress", "stack",
		InitializeScenario, nil).Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
