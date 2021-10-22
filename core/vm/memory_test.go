package vm

import (
	"Petra/common"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/smartystreets/assertions"
	"reflect"
	"testing"
)

var memory *Memory
var memoryError error

func aNewMemoryStoreIsCreated() {
	memory = NewMemory()
}

func storeIsEmpty() error {
	var result = assertions.ShouldBeZeroValue(len(memory.data))
	if result != "" {
		return errors.New(result)
	}
	return nil
}

func sizeIs(existingSize int) {
	memory.data = append(memory.data, make([]byte, existingSize)...)
}

func resizedTo(newSize int) {
	memory.ResizeIfLessThan(newSize)
}

func updatedSizeIs(expectedSize int) error {
	var result string = assertions.ShouldEqual(len(memory.data), expectedSize)
	if result != "" {
		return errors.New(result)
	}
	return nil
}

func storeIsInitializedWith(data []byte) {
	return godog.ErrPending
}

func settingOffsetAndSizeTo(offset, size uint64, data []byte) {
	return godog.ErrPending
}

func dataShouldBe(expectedData []byte) error {
	return godog.ErrPending
}

func errorIs(expectedErrorType string) error {
	var errType reflect.Type = reflect.TypeOf(memoryError)
	if errType.String() != expectedErrorType {
		return fmt.Errorf("expected %s but got %s", expectedErrorType, errType.String())
	}
	return nil
}

func dataIsACopy() error {
	return godog.ErrPending
}

func dataIsNotACopy() error {
	return godog.ErrPending
}

func gettingACopyAtOffsetAndSize(arg1, arg2 string) error {
	return godog.ErrPending
}

func gettingAPtrAtOffsetAndSize(arg1, arg2 string) error {
	return godog.ErrPending
}

func settingStartingAtAndSize(arg1, arg2, arg3 string) error {
	return godog.ErrPending
}

func InitializeMemoryScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a new memory store is created$`, aNewMemoryStoreIsCreated)
	ctx.Step(`^data is a copy$`, dataIsACopy)
	ctx.Step(`^data is not a copy$`, dataIsNotACopy)
	ctx.Step(`^data should be "([^"]*)"$`, dataShouldBe)
	ctx.Step(`^error is "([^"]*)"$`, errorIs)
	ctx.Step(`^getting a copy at offset "([^"]*)" and size "([^"]*)"$`, gettingACopyAtOffsetAndSize)
	ctx.Step(`^getting a ptr at offset "([^"]*)" and size "([^"]*)"$`, gettingAPtrAtOffsetAndSize)
	ctx.Step(`^resized to "([^"]*)"$`, resizedTo)
	ctx.Step(`^setting "([^"]*)" starting at "([^"]*)" and size "([^"]*)"$`, settingStartingAtAndSize)
	ctx.Step(`^size is "([^"]*)"$`, sizeIs)
	ctx.Step(`^store is empty$`, storeIsEmpty)
	ctx.Step(`^updated size is "([^"]*)"$`, updatedSizeIs)
	ctx.Step(`^store is initialized with "([^"]*)"$`, storeIsInitializedWith)
	ctx.Step(`^Setting offset "([^"]*)" and size "([^"]*)" to "([^"]*)"$`, settingOffsetAndSizeTo)
}

func TestMemory(t *testing.T) {
	suite := common.GetGodogTestSuite("features", "progress", "memory",
		InitializeMemoryScenario, nil)

	if suite.Run() != 0 {
		t.Fatal("Failed to run memory feature test")
	}
}
