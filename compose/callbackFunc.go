package compose

// CallbackFunc represents an function that accepts a value and perform side.
// effects. This is a specialized form of Func, and thus is constructible from
// a Func.
type CallbackFunc func(interface{}) error

// CallbackFuncConvertible represents an object that can be converted to a
// CallbackFunc.
type CallbackFuncConvertible interface {
	ToCallbackFunc() CallbackFunc
}

// ToCallbackFunc returns the current CallbackFunc.
func (cf CallbackFunc) ToCallbackFunc() CallbackFunc {
	return cf
}

// Invoke invokes the underlying function.
func (cf CallbackFunc) Invoke(value interface{}) error {
	return cf(value)
}

// ToFunc converts the current CallbackFunc into a ToFunc.
func (cf CallbackFunc) ToFunc() Func {
	return func(value interface{}) (interface{}, error) {
		err := cf(value)
		return nil, err
	}
}

// ToCallbackFunc converts the current Func into an ToCallbackFunc.
func (f Func) ToCallbackFunc() CallbackFunc {
	return func(value interface{}) error {
		_, err := f.Invoke(value)
		return err
	}
}
