package vm

import "github.com/cucumber/godog"

func addIsCalled() error {
	return godog.ErrPending
}

func resultIsSum() error {
	return godog.ErrPending
}

func xAndY() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Add is called$`, addIsCalled)
	ctx.Step(`^result is sum$`, resultIsSum)
	ctx.Step(`^x and y$`, xAndY)
}
