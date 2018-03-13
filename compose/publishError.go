package compose

// PublishError publishes an error for side effects.
func PublishError(callback func(error)) ErrorFF {
	return func(f ErrorF) ErrorF {
		return func() error {
			err := f()
			callback(err)
			return err
		}
	}
}
