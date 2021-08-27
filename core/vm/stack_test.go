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
var data uint256.Int
var err error

func createStack() {
	stack = getStack()
}

func stackShouldBeEmpty() error {
	if stack.size() != 0 {
		return fmt.Errorf("expected length 0, but found %d", stack.size())
	}
	return nil
}

func anEmptyStack() {
	stack = getStack()
}

func peekIsCalled() {
	data, err = stack.peek()
}

func popIsCalled() {
	data, err = stack.pop()
}

func errorShouldBe(expectedError *godog.DocString) error {
	if !strings.EqualFold(expectedError.Content, err.Error()) {
		return errors.New("unexpected error")
	}
	return nil
}

func isPushed(data sc
return godog.ErrPending
}

func popIsCalledTimes(count int) error {
	return godog.ErrPending
}

func poppedDataIs(data string) error {
	return godog.ErrPending
}

func stackSizeIs(size int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a stack is created$`, createStack)
	ctx.Step(`^stack should be empty$`, stackShouldBeEmpty)
	ctx.Step(`^an empty stack$`, anEmptyStack)
	ctx.Step(`^peek is called$`, peekIsCalled)
	ctx.Step(`^pop is called$`, popIsCalled)
	ctx.Step(`^error should be:$`, errorShouldBe)

	//ctx.Step(`^error should be "([^"]*)"$`, errorShouldBe)
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
