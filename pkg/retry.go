package compose

import (
	"time"
)

// ErrorFunc represents an error-returning function.
type ErrorFunc func() error

// ErrorFuncMap transforms an ErrorFunc into another ErrorFunc.
type ErrorFuncMap func(ErrorFunc) ErrorFunc

// RetryIndexErrorFunc represents an error-returning function that also tracks
// the current retry index.
type RetryIndexErrorFunc func(uint) error

// CountRetry composes an error function with retry capabilities. The error
// function has access to the current retry count in its first parameter, which
// is useful e.g when we are implementing a delay mechanism.
func CountRetry(retryCount uint) func(RetryIndexErrorFunc) ErrorFunc {
	return func(f RetryIndexErrorFunc) ErrorFunc {
		var retryF RetryIndexErrorFunc

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
func Retry(retryCount uint) func(ErrorFunc) ErrorFunc {
	return func(f ErrorFunc) ErrorFunc {
		return CountRetry(retryCount)(func(retry uint) error {
			return f()
		})
	}
}

// delayRetry composes a function with retry-delaying capabilities. The output
// of the return function can be fed to a CountRetry composition.
func delayRetry(d time.Duration) func(ErrorFunc) RetryIndexErrorFunc {
	return func(f ErrorFunc) RetryIndexErrorFunc {
		return func(retry uint) error {
			if retry > 0 {
				time.Sleep(d)
			}

			return f()
		}
	}
}

// DelayRetry composes retry with delay capabilities.
func DelayRetry(retryCount uint) func(time.Duration) ErrorFuncMap {
	return func(delay time.Duration) ErrorFuncMap {
		return func(f ErrorFunc) ErrorFunc {
			return CountRetry(retryCount)(delayRetry(delay)(f))
		}
	}
}

// PublishError publishes an error for side effects.
func PublishError(callback func(error)) ErrorFuncMap {
	return func(f ErrorFunc) ErrorFunc {
		return func() error {
			err := f()
			callback(err)
			return err
		}
	}
}
