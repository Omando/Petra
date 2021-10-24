package common

import (
	"github.com/cucumber/godog"
	"time"
)

func GetGodogTestSuite(path string, format string, suiteName string, tagName string,
	initScenario func(ctx *godog.ScenarioContext), initSuite func(*godog.TestSuiteContext)) godog.TestSuite {
	opts := godog.Options{
		Tags:      tagName,
		Format:    format,
		Paths:     []string{path},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	}

	var testSuite godog.TestSuite = godog.TestSuite{
		Name:                 suiteName,
		TestSuiteInitializer: initSuite,
		ScenarioInitializer:  initScenario,
		Options:              &opts,
	}

	return testSuite
}
