package common

import (
	"github.com/cucumber/godog"
)

func GetGodogTestSuite(path string, initScenario func(ctx *godog.ScenarioContext)) godog.TestSuite {
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{path},
		// Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	}

	var testSuite godog.TestSuite = godog.TestSuite{
		Name: "stack",
		// TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer: initScenario,
		Options:             &opts,
	}

	return testSuite
}
