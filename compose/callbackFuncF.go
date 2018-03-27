package compose

// CallbackFuncF converts an CallbackFunc to another CallbackFunc.
type CallbackFuncF func(CallbackFunc) CallbackFunc

// CallbackFuncFConvertible represents an object that can be converted to a
// CallbackFuncF.
type CallbackFuncFConvertible interface {
	ToCallbackFuncF() CallbackFuncF
}

// ToCallbackFuncF returns the current CallbackFuncF.
func (cff CallbackFuncF) ToCallbackFuncF() CallbackFuncF {
	return cff
}

// Wrap is a convenience method that calls the underlying CallbackFuncF.
func (cff CallbackFuncF) Wrap(ef CallbackFunc) CallbackFunc {
	return cff(ef)
}

// ToFuncF converts the current CallbackF into a ToFuncF.
func (cff CallbackFuncF) ToFuncF() FuncF {
	return func(f Func) Func {
		callback := func(value interface{}) error {
			_, err := f.Invoke(value)
			return err
		}

		return cff.Wrap(callback).ToFunc()
	}
}

// ToCallbackFuncF converts a FuncF to an ToCallbackFuncF. This method should be
// used at the last stage of a chain so we do not need to redefine other higher
// order functions.
func (ff FuncF) ToCallbackFuncF() CallbackFuncF {
	return func(cf CallbackFunc) CallbackFunc {
		var function Func = func(value interface{}) (interface{}, error) {
			return nil, cf(value)
		}

		return ff.Wrap(function).ToCallbackFunc()
	}
}
