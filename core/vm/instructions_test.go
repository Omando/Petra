package vm

import (
	"Petra/common"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/holiman/uint256"
	"testing"
)

var operand1 uint256.Int
var operand2 uint256.Int

func stringToUint256(hexValue string) (uint256.Int, error) {
	// Convert a hex string to byte[]
	bytes, err := common.FromHex(hexValue)
	if err != nil {
		return uint256.Int{}, err
	}

	result := uint256.Int{}
	result.SetBytes(bytes)
	return result, nil
}

func addOperands(x, y string) error {
	var err error
	operand1, err = stringToUint256(x)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not convert %s to uint256.int", x))
	}

	operand2, err = stringToUint256(y)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not convert %s to uint256.int", y))
	}

	return nil
}

func addIsCalled() error {
	return godog.ErrPending
}

func resultIs(arg1 string) error {
	return godog.ErrPending
}

func InitializeInstructionsScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Add is called$`, addIsCalled)
	ctx.Step(`^"([^"]*)" and "([^"]*)" operands$`, addOperands)
	ctx.Step(`^result is "([^"]*)"$`, resultIs)
}

func TestInstructions(t *testing.T) {
	suite := common.GetGodogTestSuite(
		"features",                     // path to features folder
		"progress",                     // godog formatter name
		"instructions",                 // suite name
		"run",                          // tag name (scenarios without this tag will not run)
		InitializeInstructionsScenario, // function used to setup godog steps
		nil)

	if suite.Run() != 0 {
		t.Fatal("Failed to run memory feature test")
	}
}
