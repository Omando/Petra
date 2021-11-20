package vm

func opAdd(pc *uint64, interpreter *EVMInterpreter, scope *CallContext) ([]byte, error) {
	x, xerr := scope.Stack.pop()
	if xerr != nil {
		return nil, xerr
	}

	y, yerr := scope.Stack.peek()
	if yerr != nil {
		return nil, yerr
	}

	y.Add(&x, y)
	return nil, nil
}

func opSub(pc *uint64, interpreter *EVMInterpreter, scope *CallContext) ([]byte, error) {
	x, xerr := scope.Stack.pop()
	if xerr != nil {
		return nil, xerr
	}

	y, yerr := scope.Stack.peek()
	if yerr != nil {
		return nil, yerr
	}

	y.Sub(&x, y)
	return nil, nil
}

func opMul(pc *uint64, interpreter *EVMInterpreter, scope *CallContext) ([]byte, error) {
	x, xerr := scope.Stack.pop()
	if xerr != nil {
		return nil, xerr
	}

	y, yerr := scope.Stack.peek()
	if yerr != nil {
		return nil, yerr
	}

	y.Mul(&x, y)
	return nil, nil
}

func opDiv(pc *uint64, interpreter *EVMInterpreter, scope *CallContext) ([]byte, error) {
	x, xerr := scope.Stack.pop()
	if xerr != nil {
		return nil, xerr
	}

	y, yerr := scope.Stack.peek()
	if yerr != nil {
		return nil, yerr
	}

	y.Div(&x, y)
	return nil, nil
}
