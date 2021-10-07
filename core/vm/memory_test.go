package vm

import (
	"Petra/common"
	"errors"
	"github.com/cucumber/godog"
	"github.com/smartystreets/assertions"
	"testing"
)

var memory *Memory

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

func sizeIs(oldSize int) error {
	dummyData := make([]byte, oldSize)
	for i := 0; i < oldSize; i++ {
		dummyData[i] = byte(i)
	}
	err := memory.Set(0, 10, dummyData)
	return err
}

func resizedTo(newSize int) {
	memory.ResizeIfLessThan(newSize)
}

func aMemoryIsCreatedAndInitializedWith(arg1 string) error {
	return godog.ErrPending
}

func dataIsACopy() error {
	return godog.ErrPending
}

func dataIsNotACopy() error {
	return godog.ErrPending
}

func dataShouldBe(arg1 string) error {
	return godog.ErrPending
}

func errorIs(arg1 string) error {
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
	ctx.Step(`^a memory is created and initialized with "([^"]*)"$`, aMemoryIsCreatedAndInitializedWith)
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
}

func TestMemory(t *testing.T) {
	suite := common.GetGodogTestSuite("features", "progress", "memory",
		InitializeMemoryScenario, nil)

	if suite.Run() != 0 {
		t.Fatal("Failed to run memory feature test")
	}
}
