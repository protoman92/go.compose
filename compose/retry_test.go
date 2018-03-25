package compose

import (
	"testing"
	"time"
)

func TestCountRetryCompose(t *testing.T) {
	/// Setup
	var currentRetry uint

	errF := func(retry uint) (interface{}, error) {
		currentRetry = retry
		return valueOp, errOp
	}

	/// When & Then
	if _, err := CountRetryF(retries)(errF)(); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}

	if currentRetry != retries {
		t.Errorf("Expected %v, got %v", retries, currentRetry)
	}
}

func TestRetryComposeWithInitialError(t *testing.T) {
	/// Setup
	invoked := uint(0)

	errF := func() (interface{}, error) {
		defer func() {
			invoked++
		}()

		if invoked == 0 {
			return nil, errOp
		}

		return valueOp, nil
	}

	/// When & Then
	if value, err := RetryF(retries)(errF)(); err != nil || value != valueOp {
		t.Errorf("Should not error, but got %v", err)
	}

	if invoked != 2 {
		t.Errorf("Invoked %d times", invoked)
	}
}

func TestRetryComposeWithAllErrors(t *testing.T) {
	/// Setup
	invoked := uint(0)

	var errF Func = func() (interface{}, error) {
		invoked++
		return valueOp, errOp
	}

	/// When & Then
	if _, err := errF.Retry(retries).Invoke(); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}

	if invoked != retries+1 {
		t.Errorf("Expected %d, got %d", retries+1, invoked)
	}
}

func TestDelayRetry(t *testing.T) {
	/// Setup
	currentRetry := uint(2)

	errF := func() (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	start := time.Now()

	if _, err := delayRetry(delay)(errF)(currentRetry); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}

	difference := time.Now().Sub(start)

	if difference < delay {
		t.Errorf("Expected %d, got %d", delay, difference)
	}
}

func TestDelayRetryForFirstInvocation(t *testing.T) {
	/// Setup
	errF := func() (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	start := time.Now()

	if _, err := delayRetry(delay)(errF)(0); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}

	difference := time.Now().Sub(start)

	if difference >= delay {
		t.Errorf("Should not have delayed, but got %d", difference)
	}
}

func TestDelayedRetry(t *testing.T) {
	/// Setup
	var errF Func = func() (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	start := time.Now()

	if _, err := errF.DelayRetry(retries)(delay).Invoke(); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}

	difference := time.Now().Sub(start)

	if int64(difference) < int64(delay)*int64(retries) {
		t.Errorf("Wrong delay duration %d", difference)
	}
}
