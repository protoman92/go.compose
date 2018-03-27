package compose

import "time"

// CountRetryF composes a Func with retry capabilities. The error function has
// access to the current retry count in its first parameter, which is useful
// e.g when we are implementing a delay mechanism.
func CountRetryF(retryCount uint) func(func(uint, interface{}) (interface{}, error)) Func {
	return func(f func(uint, interface{}) (interface{}, error)) Func {
		var retryF func(uint, interface{}) (interface{}, error)

		retryF = func(current uint, value interface{}) (interface{}, error) {
			value, err := f(current, value)

			if err != nil {
				if current < retryCount {
					return retryF(current+1, value)
				}

				return nil, err
			}

			return value, nil
		}

		return func(value interface{}) (interface{}, error) {
			return retryF(0, value)
		}
	}
}

// RetryF has the same semantics as CountRetry, but ignores the current retry
// count.
func RetryF(retryCount uint) FuncF {
	return func(f Func) Func {
		return CountRetryF(retryCount)(func(retry uint, value interface{}) (interface{}, error) {
			return f(value)
		})
	}
}

// Retry is a convenience method to chain Func, using the compose RetryF under
// the hood.
func (f Func) Retry(retryCount uint) Func {
	return RetryF(retryCount).Wrap(f)
}

// delayRetry composes a function with retry-delaying capabilities. The output
// of the return function can be fed to a CountRetry composition.
func delayRetry(d time.Duration) func(Func) func(uint, interface{}) (interface{}, error) {
	return func(f Func) func(uint, interface{}) (interface{}, error) {
		return func(retry uint, value interface{}) (interface{}, error) {
			if retry > 0 {
				time.Sleep(d)
			}

			return f(value)
		}
	}
}

// DelayRetryF composes retry with delay capabilities.
func DelayRetryF(retryCount uint) func(time.Duration) FuncF {
	return func(delay time.Duration) FuncF {
		return func(f Func) Func {
			return CountRetryF(retryCount)(delayRetry(delay)(f))
		}
	}
}

// DelayRetry is a convenience function that calls the composable DelayRetryF
// function under the hood.
func (f Func) DelayRetry(retryCount uint) func(time.Duration) Func {
	return func(d time.Duration) Func {
		return DelayRetryF(retryCount)(d).Wrap(f)
	}
}
