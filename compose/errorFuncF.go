package compose

// ErrorFuncF converts an ErrorFunc to another ErrorFunc.
type ErrorFuncF func(ErrorFunc) ErrorFunc

// Wrap is a convenience method that calls the underlying ErrorFuncF.
func (eff ErrorFuncF) Wrap(ef ErrorFunc) ErrorFunc {
	return eff(ef)
}

// ErrorFuncF converts a FuncF to an ErrorFuncF. This method should be used at
// the last stage of a chain so we do not need to redefine other higher-order
// functions.
func (ff FuncF) ErrorFuncF() ErrorFuncF {
	return func(ef ErrorFunc) ErrorFunc {
		var function Func = func() (interface{}, error) {
			return nil, ef()
		}

		wrapped := ff.Wrap(function)

		return func() error {
			_, err := wrapped.Invoke()
			return err
		}
	}
}
