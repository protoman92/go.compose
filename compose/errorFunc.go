package compose

// ErrorFunc represents an error-returning function. This is a specialized
// form of Func, and thus is constructible from a Func.
type ErrorFunc func() error

// Invoke invokes the underlying error function.
func (ef ErrorFunc) Invoke() error {
	return ef()
}
