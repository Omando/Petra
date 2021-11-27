package vm

import (
	"github.com/cucumber/godog"
	"github.com/holiman/uint256"
)

var operand1 uint256.Int
var operand2 uint256.Int

func stringToUint256(hexValue string) uint256.Int {

}

func addIsCalled() error {
	return godog.ErrPending
}

func addOperands(x, y string) error {
	operand1 = stringToUint256(x)
	operand2 = stringToUint256(y)
	return nil
}

func resultIs(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Add is called$`, addIsCalled)
	ctx.Step(`^"([^"]*)" and "([^"]*)" operands$`, addOperands)
	ctx.Step(`^result is "([^"]*)"$`, resultIs)
}
