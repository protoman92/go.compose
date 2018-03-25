package compose

// NoopF does nothing and simply returns the function.
func NoopF() FunctionF {
	return func(f Function) Function {
		return f
	}
}

// Noop is a convenience function that calls the composable NoopF under the hood.
func (f Function) Noop() Function {
	return NoopF().Wrap(f)
}
