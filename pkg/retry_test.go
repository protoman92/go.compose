package compose

import (
	"errors"
	"testing"
)

func TestRetryComposeWithAllErrors(t *testing.T) {
	/// Setup
	err := errors.New("error")
	retryCount := uint(10)
	invoked := uint(0)

	errF := func() error {
		invoked++
		return err
	}

	/// When & Then
	if err1 := Retry(errF, retryCount)(); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	if invoked != retryCount+1 {
		t.Errorf("Expected %d, got %d", retryCount+1, invoked)
	}
}

func TestRetryComposeWithInitialError(t *testing.T) {
	/// Setup
	err := errors.New("error")
	retryCount := uint(10)
	invoked := uint(0)

	errF := func() error {
		defer func() {
			invoked++
		}()

		if invoked == 0 {
			return err
		}

		return nil
	}

	/// When & Then
	if err1 := Retry(errF, retryCount)(); err1 != nil {
		t.Errorf("Should not error, but got %v", err1)
	}

	if invoked != 2 {
		t.Errorf("Invoked %d times", invoked)
	}
}
