package compose

import (
	"errors"
	"testing"
	"time"
)

const (
	delayDuration = time.Duration(1e7)
	retryCount    = uint(10)
)

var (
	err = errors.New("error")
)

func TestRetryComposeWithAllErrors(t *testing.T) {
	/// Setup
	invoked := uint(0)

	errF := func() error {
		invoked++
		return err
	}

	/// When & Then
	if err1 := Retry(retryCount)(errF)(); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	if invoked != retryCount+1 {
		t.Errorf("Expected %d, got %d", retryCount+1, invoked)
	}
}

func TestRetryComposeWithInitialError(t *testing.T) {
	/// Setup
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
	if err1 := Retry(retryCount)(errF)(); err1 != nil {
		t.Errorf("Should not error, but got %v", err1)
	}

	if invoked != 2 {
		t.Errorf("Invoked %d times", invoked)
	}
}

func TestCountRetryCompose(t *testing.T) {
	/// Setup
	var currentRetry uint

	errF := func(retry uint) error {
		currentRetry = retry
		return err
	}

	/// When & Then
	if err1 := CountRetry(retryCount)(errF)(); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	if currentRetry != retryCount {
		t.Errorf("Expected %v, got %v", retryCount, currentRetry)
	}
}

func TestDelayRetry(t *testing.T) {
	/// Setup
	currentRetry := uint(2)

	errF := func() error {
		return err
	}

	/// When & Then
	start := time.Now()

	if err1 := DelayRetry(delayDuration)(errF)(currentRetry); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	difference := time.Now().Sub(start)

	if difference < delayDuration {
		t.Errorf("Expected %d, got %d", delayDuration, difference)
	}
}

func TestDelayRetryForFirstInvocation(t *testing.T) {
	/// Setup
	errF := func() error {
		return err
	}

	/// When & Then
	start := time.Now()

	if err1 := DelayRetry(delayDuration)(errF)(0); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	difference := time.Now().Sub(start)

	if difference >= delayDuration {
		t.Errorf("Should not have delayed, but got %d", difference)
	}
}

func TestDelayedRetry(t *testing.T) {
	/// Setup
	errF := func() error {
		return err
	}

	/// When & Then
	start := time.Now()

	if err1 := RetryWithDelay(retryCount)(delayDuration)(errF)(); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}

	difference := time.Now().Sub(start)

	if int64(difference) < int64(delayDuration)*int64(retryCount) {
		t.Errorf("Wrong delay duration %d", difference)
	}
}
