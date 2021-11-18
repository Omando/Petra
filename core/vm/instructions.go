package vm

func opAdd(pc *uint64, interpreter *EVMInterpreter, scope *CallContext) ([]byte, error) {
	x, xerr := scope.Stack.pop()
	y, yerr := scope.Stack.peek()

	y.Add(&x, y)
	return nil, nil
}
