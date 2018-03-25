package compose

// FuncF transforms a Func into another Func.
type FuncF func(Func) Func

// Compose composes the functionalities of both FuncF. We can use this to chain
// enhance a base Func without exposing implementation details.
func (ff FuncF) Compose(selector FuncF) FuncF {
	return func(f Func) Func {
		return ff(selector(f))
	}
}

// ComposeFn is similar to Compose, but it is more convenient when we deal with
// functions that return FuncF.
func (ff FuncF) ComposeFn(selectorFn func() FuncF) FuncF {
	return ff.Compose(selectorFn())
}

// Wrap is a convenience method to invoke the wrap on a Func. This may look
// nicer in a function chain.
func (ff FuncF) Wrap(f Func) Func {
	return ff(f)
}
