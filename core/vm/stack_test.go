package vm

import (
	"fmt"
	"github.com/cucumber/godog"
)

var stack *Stack

func createStack() {
	stack = getStack()
}

func stackShouldBeEmpty() error {
	length := len(stack.data)
	if length != 0 {
		return fmt.Errorf("expected length 0, but found %d", length)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a stack is created$`, createStack)
	ctx.Step(`^stack should be empty$`, stackShouldBeEmpty)
}
