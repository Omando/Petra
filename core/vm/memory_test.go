package vm

import (
	"Petra/common"
	"bytes"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/smartystreets/assertions"
	"reflect"
	"testing"
)

var memory *Memory
var memoryError error
var dataCopy []byte
var dataOffset, dataSize int

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
	memory.data = append(memory.data, data...)
}

func settingOffsetAndSizeTo(offset, size int, data []byte) {
	memoryError = memory.Set(uint64(offset), uint64(size), data)
}

func dataShouldBe(expectedData []byte) error {
	result := bytes.Equal(memory.Data(), expectedData)
	if result == false {
		return errors.New("Not equal")
	}
	return nil
}

func gettingACopyAtOffsetAndSize(offset, size int) {
	dataOffset = offset
	dataSize = size
	dataCopy, memoryError = memory.GetCopy(uint(offset), uint(size))
}

func gettingAPtrAtOffsetAndSize(offset, size int) {
	dataOffset = offset
	dataSize = size
	dataCopy, memoryError = memory.GetPtr(uint(offset), uint(size))
}

func getptrDataShouldBe(expectedData []byte) error {
	result := bytes.Equal(dataCopy, expectedData)
	if result == false {
		return errors.New("Not equal")
	}
	return nil
}

func errorIs(expectedErrorType string) error {
	// Get type of error reported by last operation
	var errType reflect.Type = reflect.TypeOf(memoryError)

	/* Check if error reported by last operation matched expectation */
	// No error should been generated
	if errType == nil && len(expectedErrorType) == 0 {
		return nil
	}

	// Does last error match our expectation?
	if errType.String() != expectedErrorType {
		return fmt.Errorf("expected %s but got %s", expectedErrorType, errType.String())
	}
	return nil
}

func dataIsACopy() error {
	// Ignore this operation memory copy is nil
	if dataCopy == nil {
		return nil
	}

	// We did get data. Compare if the slices refer to the same underlying array
	var addr1 *byte = &dataCopy[0]
	var addr2 *byte = &memory.data[dataOffset]

	if addr1 == addr2 {
		return errors.New("copied data points to the same underlying array as memory")
	}
	return nil
}

func dataIsNotACopy() error {
	// Ignore this operation memory copy is nil
	if dataCopy == nil {
		return nil
	}

	// We did get data. Compare if the slices refer to the same underlying array
	var addr1 *byte = &dataCopy[0]
	var addr2 *byte = &memory.data[dataOffset]

	if addr1 != addr2 {
		return errors.New("getptr data does not point to the same underlying array as memory")
	}
	return nil
}

func InitializeMemoryScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a new memory store is created$`, aNewMemoryStoreIsCreated)
	ctx.Step(`^data is a copy$`, dataIsACopy)
	ctx.Step(`^data is not a copy$`, dataIsNotACopy)
	ctx.Step(`^data should be "([^"]*)"$`, dataShouldBe)
	ctx.Step(`^copied data should be "([^"]*)"$`, copiedDataShouldBe)
	ctx.Step(`^getptr data should be "([^"]*)"$`, getptrDataShouldBe)
	ctx.Step(`^error is "([^"]*)"$`, errorIs)
	ctx.Step(`^getting a copy at offset "([^"]*)" and size "([^"]*)"$`, gettingACopyAtOffsetAndSize)
	ctx.Step(`^getting a ptr at offset "([^"]*)" and size "([^"]*)"$`, gettingAPtrAtOffsetAndSize)
	ctx.Step(`^resized to "([^"]*)"$`, resizedTo)
	ctx.Step(`^size is "([^"]*)"$`, sizeIs)
	ctx.Step(`^store is empty$`, storeIsEmpty)
	ctx.Step(`^updated size is "([^"]*)"$`, updatedSizeIs)
	ctx.Step(`^store is initialized with "([^"]*)"$`, storeIsInitializedWith)
	ctx.Step(`^Setting offset "([^"]*)" and size "([^"]*)" to "([^"]*)"$`, settingOffsetAndSizeTo)
}

func TestMemory(t *testing.T) {
	suite := common.GetGodogTestSuite("features", "progress", "memory", "run",
		InitializeMemoryScenario, nil)

	if suite.Run() != 0 {
		t.Fatal("Failed to run memory feature test")
	}
}
