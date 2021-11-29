package vm

import (
	"Petra/common"
	"github.com/cucumber/godog"
	"github.com/holiman/uint256"
)

var operand1 uint256.Int
var operand2 uint256.Int

func stringToUint256(hexValue string) (uint256.Int, error) {
	bytes, err := common.FromHex(hexValue)
	if err != nil {
		return uint256.Int{}, err
	}
}

func addOperands(x, y string) error {

}

func addIsCalled() error {
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
