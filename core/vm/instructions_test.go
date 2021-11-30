package vm

import (
	"Petra/common"
	"errors"
	"fmt"
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
	var err error
	operand1, err = stringToUint256(x)
	if err != nil {
		return errors.New(fmt.Sprint("Could not convert %s to uint256.int", x))
	}

	operand2, err = stringToUint256(y)
	if err != nil {
		return errors.New(fmt.Sprint("Could not convert %s to uint256.int", y))
	}
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
