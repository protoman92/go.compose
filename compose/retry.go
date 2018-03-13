package compose

import (
	"time"
)

// CountRetry composes an error function with retry capabilities. The error
// function has access to the current retry count in its first parameter, which
// is useful e.g when we are implementing a delay mechanism.
func CountRetry(retryCount uint) func(func(uint) error) ErrorF {
	return func(f func(uint) error) ErrorF {
		var retryF func(uint) error

		retryF = func(current uint) error {
			if err := f(current); err != nil {
				if current < retryCount {
					return retryF(current + 1)
				}

				return err
			}

			return nil
		}

		return func() error {
			return retryF(0)
		}
	}
}

// Retry has the same semantics as CountRetry, but ignores the current retry
// count.
func Retry(retryCount uint) func(ErrorF) ErrorF {
	return func(f ErrorF) ErrorF {
		return CountRetry(retryCount)(func(retry uint) error {
			return f()
		})
	}
}

// delayRetry composes a function with retry-delaying capabilities. The output
// of the return function can be fed to a CountRetry composition.
func delayRetry(d time.Duration) func(ErrorF) func(uint) error {
	return func(f ErrorF) func(uint) error {
		return func(retry uint) error {
			if retry > 0 {
				time.Sleep(d)
			}

			return f()
		}
	}
}

// DelayRetry composes retry with delay capabilities.
func DelayRetry(retryCount uint) func(time.Duration) ErrorFF {
	return func(delay time.Duration) ErrorFF {
		return func(f ErrorF) ErrorF {
			return CountRetry(retryCount)(delayRetry(delay)(f))
		}
	}
}
