package compose

// ErrorF represents an error-returning function.
type ErrorF func() error

// ErrorFF transforms an ErrorF into another ErrorF.
type ErrorFF func(ErrorF) ErrorF

// Compose composes the functionalities of both ErrorFF. For e.g, we have two
// ErrorFF Retry and Publish, assuming callback is a side effect function that
// prints the error to output.
//
// If we call Retry(2).Compose(Publish(callback))(errF)(), we will see 3 lines
// of error message being printed. This is because the output of publish becomes
// the input for Retry, meaning that every time Retry invokes the function, the
// inner error will be invoked. For 2 retries, there are actually 3 invocations
// (including the initial call), so it is logical that there are 3 publications.
//
// If we call Publish(callback).Compose(Retry(2))(errF)(), we only see 1 line of
// error printed out. The logic is similar: when the error is being published,
// the actual error functions would already have been retried twice. The publish
// function is not aware of this, however; it only knows the output of Retry.
func (ff ErrorFF) Compose(selector ErrorFF) ErrorFF {
	return func(f ErrorF) ErrorF {
		return ff(selector(f))
	}
}

// ComposeFn is similar to Compose, but it is more convenient when we deal with
// functions that return ErrorFF.
func (ff ErrorFF) ComposeFn(selectorFn func() ErrorFF) ErrorFF {
	return ff.Compose(selectorFn())
}
