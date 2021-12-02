# GO Petra 

## Installations
### GODOG
```
go install github.com/cucumber/godog/cmd/godog@v0.12.0
```
Additionally install Gherkin plugin

### Other libraries
```go
 go get github.com/holiman/uint256       
```

## Code structure


## Building the source

## Running
## Running tests
Change to relevant directory and run godog or go test
```go
godog
go test -v
```

To run a specific feature, specify folder and feature name:
```go
godog .\features\memory.feature
```

To run all features in a specific features folder
```go
godog .\features
```

To debug a features test, add the following to the test file
```go
func TestMemory(t *testing.T) {
	suite := common.GetGodogTestSuite(
		"features",                 // path to features folder 
		"progress",                 // godog formatter name 
		"memory",                   // suite name
		"run",                      // tag name (scenarios without this tag will not run)
		InitializeMemoryScenario,   // function used to setup godog steps 
		nil)

	if suite.Run() != 0 {
		t.Fatal("Failed to run memory feature test")
	}
}
```
## Technical Notes

## Links
[godog](https://github.com/cucumber/godog)