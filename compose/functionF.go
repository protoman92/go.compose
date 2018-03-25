package compose

// FunctionF transforms a Function into another Function.
type FunctionF func(Function) Function

// Compose composes the functionalities of both FunctionF. We can use this to
// chain enhance a base Function without exposing implementation details.
func (ff FunctionF) Compose(selector FunctionF) FunctionF {
	return func(f Function) Function {
		return ff(selector(f))
	}
}

// ComposeFn is similar to Compose, but it is more convenient when we deal with
// functions that return FunctionF.
func (ff FunctionF) ComposeFn(selectorFn func() FunctionF) FunctionF {
	return ff.Compose(selectorFn())
}

// Wrap is a convenience method to invoke the wrap on a Function. This may look
// nice in a function chain.
func (ff FunctionF) Wrap(f Function) Function {
	return ff(f)
}
