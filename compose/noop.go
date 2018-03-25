package compose

// NoopF does nothing and simply returns the function.
func NoopF() FuncF {
	return func(f Func) Func {
		return f
	}
}

// Noop is a convenience function that calls the composable NoopF under the hood.
func (f Func) Noop() Func {
	return NoopF().Wrap(f)
}
