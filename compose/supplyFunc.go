package compose

// SupplyFunc is a specialized form of Func that produces a value.
type SupplyFunc func() (interface{}, error)

// SupplyFuncConvertible represents an object that can be converted to a SupplyFunc.
type SupplyFuncConvertible interface {
	ToSupplyFunc() SupplyFunc
}

// ToSupplyFunc returns the current SupplyFunc.
func (sf SupplyFunc) ToSupplyFunc() SupplyFunc {
	return sf
}

// Invoke is a convenience method that calls the underlying supply function.
func (sf SupplyFunc) Invoke() (interface{}, error) {
	return sf()
}

// ToFunc converts the current SupplyFunc into a ToFunc.
func (sf SupplyFunc) ToFunc() Func {
	return func(value interface{}) (interface{}, error) {
		return sf.Invoke()
	}
}

// ToSupplyFunc converts the current Func into a ToSupplyFunc.
func (f Func) ToSupplyFunc() SupplyFunc {
	return func() (interface{}, error) {
		return f.Invoke(nil)
	}
}
