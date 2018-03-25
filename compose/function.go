package compose

// Func represents an operation that could return an error.
type Func func() (interface{}, error)

// Invoke is a convenience method to call a Func. Although it is the same as if
// we call the function normally, this may look nicer in a chain.
func (f Func) Invoke() (interface{}, error) {
	return f()
}

// ErrorFunc converts the current function into an ErrorFunc.
func (f Func) ErrorFunc() ErrorFunc {
	return func() error {
		_, err := f.Invoke()
		return err
	}
}
