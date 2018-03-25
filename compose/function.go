package compose

// Function represents an operation that could return an error.
type Function func() (interface{}, error)

// Invoke is a convenience method to call a Function. Although it is the same
// as if we call the function normally, this may look nicer in a chain.
func (f Function) Invoke() (interface{}, error) {
	return f()
}
