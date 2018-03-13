package compose

import (
	"testing"
)

func TestCompose(t *testing.T) {
	published := 0

	var retryF ErrorFF = Retry(retryCount)

	publishF := PublishError(func(err error) {
		published++
	})

	errF := func() error {
		return err
	}

	/// When && Then 1
	retryF.Compose(publishF).Compose(NoopError())(errF)()

	if uint(published) != retryCount+1 {
		t.Errorf("Expected %d, got %d", retryCount+1, published)
	}

	/// When && Then 2
	published = 0
	publishF.Compose(retryF).Compose(NoopError())(errF)()

	if published != 1 {
		t.Errorf("Expected %d, got %d", 1, published)
	}
}

func BenchmarkComposition(b *testing.B) {
	errF := func() error {
		return err
	}

	publishF := func(err error) {}
	composeF := PublishError(publishF)

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
	errF := func() error {
		return err
	}

	for i := 0; i < b.N; i++ {
		errF()
	}
}
