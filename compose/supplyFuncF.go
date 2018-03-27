package compose

// SupplyFuncF maps a SupplyFunc to another SupplyFunc.
type SupplyFuncF func(SupplyFunc) SupplyFunc

// SupplyFuncFConvertible represents an object that can be converted to a SupplyFuncF.
type SupplyFuncFConvertible interface {
	ToSupplyFuncF() SupplyFuncF
}

// ToSupplyFuncF returns the current SupplyFuncF.
func (sff SupplyFuncF) ToSupplyFuncF() SupplyFuncF {
	return sff
}

// Wrap is a convenience method to call the underlying SupplyFuncF.
func (sff SupplyFuncF) Wrap(sf SupplyFunc) SupplyFunc {
	return sff(sf)
}

// ToFuncF converts the current SupplyFuncF into a ToFuncF.
func (sff SupplyFuncF) ToFuncF() FuncF {
	return func(f Func) Func {
		sf := func() (interface{}, error) {
			return f(nil)
		}

		return sff.Wrap(sf).ToFunc()
	}
}

// ToSupplyFuncF converts the current FuncF to a ToSupplyFuncF.
func (ff FuncF) ToSupplyFuncF() SupplyFuncF {
	return func(sf SupplyFunc) SupplyFunc {
		f := func(value interface{}) (interface{}, error) {
			return sf.Invoke()
		}

		return ff.Wrap(f).ToSupplyFunc()
	}
}
