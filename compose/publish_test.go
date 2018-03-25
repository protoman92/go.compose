package compose

import (
	"testing"
)

func TestPublish(t *testing.T) {
	/// Setup
	published := 0
	var publishedValue interface{}
	var publishedErr error

	var errF Function = func() (interface{}, error) {
		return valueOp, errOp
	}

	publishF := func(value interface{}, err error) {
		published++
		publishedValue = value
		publishedErr = err
	}

	/// When & Then
	value, err := errF.Publish(publishF).Retry(retryCount).Invoke()

	if err != errOp || value != nil {
		t.Errorf("Expected %v, got %v", errOp, err)
	}

	if publishedValue != valueOp {
		t.Errorf("Expected %v, got %v", valueOp, publishedValue)
	}

	if publishedErr != errOp {
		t.Errorf("Expected %v, got %v", errOp, publishedErr)
	}

	if uint(published) != retryCount+1 {
		t.Errorf("Expected %v, got %v", retryCount+1, published)
	}
}