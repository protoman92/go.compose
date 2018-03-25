package compose

// PublishF publishes the result of a Function for side effects.
func PublishF(callback func(interface{}, error)) FunctionF {
	return func(f Function) Function {
		return func() (interface{}, error) {
			value, err := f()
			callback(value, err)
			return value, err
		}
	}
}

// Publish is a convenience function that calls the composable PublishF under
// the hood.
func (f Function) Publish(callback func(interface{}, error)) Function {
	return PublishF(callback).Wrap(f)
}
