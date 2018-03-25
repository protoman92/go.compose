package compose

import (
	"testing"
)

func TestCompose(t *testing.T) {
	published := 0

	retryF := RetryF(retries)

	publishF := PublishF(func(value interface{}, err error) {
		published++
	})

	errF := func() (interface{}, error) {
		return valueOp, errOp
	}

	/// When && Then 1
	retryF.Compose(publishF).ComposeFn(NoopF)(errF)()

	if uint(published) != retries+1 {
		t.Errorf("Expected %d, got %d", retries+1, published)
	}

	/// When && Then 2
	published = 0
	publishF.Compose(retryF).ComposeFn(NoopF)(errF)()

	if published != 1 {
		t.Errorf("Expected %d, got %d", 1, published)
	}
}

func TestComposeConvertToErrorFuncF(t *testing.T) {
	/// Setup
	errF := func() error {
		return errOp
	}

	retryF := RetryF(retries).ErrorFuncF()

	/// When & Then
	if err := retryF.Wrap(errF).Invoke(); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}

func BenchmarkComposition(b *testing.B) {
	errF := func() (interface{}, error) {
		return valueOp, errOp
	}

	publishF := func(value interface{}, err error) {}
	composeF := PublishF(publishF)

	composed := composeF.
		Compose(composeF).
		Compose(composeF).
		Compose(composeF).
		Compose(composeF)(errF)

	for i := 0; i < b.N; i++ {
		composed()
	}
}

func BenchmarkOrdinaryErrorF(b *testing.B) {
	errF := func() (interface{}, error) {
		return valueOp, errOp
	}

	for i := 0; i < b.N; i++ {
		errF()
	}
}
