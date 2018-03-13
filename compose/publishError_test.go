package compose

import (
	"testing"
)

func TestPublishError(t *testing.T) {
	/// Setup
	published := 0
	var publishedErr error

	errF := func() error {
		return err
	}

	publishF := func(err error) {
		published++
		publishedErr = err
	}

	/// When & Then
	if err1 := Retry(retryCount)(PublishError(publishF)(errF))(); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	if publishedErr != err {
		t.Errorf("Expected %v, got %v", err, publishedErr)
	}

	if uint(published) != retryCount+1 {
		t.Errorf("Expected %v, got %v", retryCount+1, published)
	}
}
