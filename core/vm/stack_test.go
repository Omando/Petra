package vm

import (
	"Petra/common"
	"fmt"
	"github.com/cucumber/godog"
	"os"
	"testing"
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

func TestMain(m *testing.M) {
	status := common.GetGodogTestSuite("features", "progress", "stack",
		InitializeScenario, nil).Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
