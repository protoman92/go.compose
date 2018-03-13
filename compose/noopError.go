package compose

// NoopError does nothing and simply returns the error function.
func NoopError() ErrorFF {
	return func(f ErrorF) ErrorF {
		return f
	}
}
