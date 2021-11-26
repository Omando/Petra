package vm

import (
	"github.com/cucumber/godog"
	"github.com/holiman/uint256"
)

var operand1 uint256.Int
var operand2 uint256.Int

func addIsCalled() error {
	return godog.ErrPending
}

func addOperands(arg1, arg2 string) error {
	return godog.ErrPending
}

func resultIs(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Add is called$`, addIsCalled)
	ctx.Step(`^"([^"]*)" and "([^"]*)" operands$`, addOperands)
	ctx.Step(`^result is "([^"]*)"$`, resultIs)
}
