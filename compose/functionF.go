package compose

// FuncF transforms a Func into another Func.
type FuncF func(Func) Func

// FuncFConvertible represents an object that can be converted to a FuncF.
type FuncFConvertible interface {
	ToFuncF() FuncF
}

// ToFuncF returns the current FuncF.
func (ff FuncF) ToFuncF() FuncF {
	return ff
}

// Compose composes the functionalities of both FuncF. We can use this to chain
// enhance a base Func without exposing implementation details.
func (ff FuncF) Compose(selector FuncFConvertible) FuncF {
	return func(f Func) Func {
		return ff(selector.ToFuncF().Wrap(f))
	}
}

// Wrap is a convenience method to invoke the wrap on a Func. This may look
// nicer in a function chain.
func (ff FuncF) Wrap(f Func) Func {
	return ff(f)
}
