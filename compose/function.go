package compose

// Func represents an operation that could return an error.
type Func func(interface{}) (interface{}, error)

// FuncConvertible represents an object that can be converted to a Func.
type FuncConvertible interface {
	ToFunc() Func
}

// ToFunc returns the current Func.
func (f Func) ToFunc() Func {
	return f
}

// Invoke is a convenience method to call a Func. Although it is the same as if
// we call the function normally, this may look nicer in a chain.
func (f Func) Invoke(value interface{}) (interface{}, error) {
	return f(value)
}
