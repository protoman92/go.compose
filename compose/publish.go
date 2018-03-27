package compose

// PublishF publishes the result of a Func for side effects.
func PublishF(callback func(interface{}, error)) FuncF {
	return func(f Func) Func {
		return func(value interface{}) (interface{}, error) {
			value, err := f(value)
			callback(value, err)
			return value, err
		}
	}
}

// Publish is a convenience function that calls the composable PublishF under
// the hood.
func (f Func) Publish(callback func(interface{}, error)) Func {
	return PublishF(callback).Wrap(f)
}
