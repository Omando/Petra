package vm

type CallContext struct {
	Memory   *Memory
	Stack    *Stack
	Contract *Contract
}

type VMInterpreter struct {
}
