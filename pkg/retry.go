package compose

// ErrorFunc represents a function that returns an error.
type ErrorFunc func() error

// Retry composes an ErrorFunc with retry capabilities.
func Retry(f ErrorFunc, retryCount uint) ErrorFunc {
	return func() error {
		var retryF func(uint) error

		retryF = func(current uint) error {
			if err := f(); err != nil {
				if current < retryCount {
					return retryF(current + 1)
				}

				return err
			}

			return nil
		}

		return retryF(0)
	}
}
