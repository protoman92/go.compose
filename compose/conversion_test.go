package compose

import (
	"testing"
)

func TestConvertBetweenFuncTypes(t *testing.T) {
	/// Setup
	var errF Func = func(value interface{}) (interface{}, error) {
		return valueOp, errOp
	}

	/// When
	value, err := errF.
		ToCallbackFunc().
		ToCallbackFunc().
		ToFunc().
		ToFunc().
		ToSupplyFunc().
		ToSupplyFunc().
		ToFunc().
		Invoke(0)

	/// Then
	if err != errOp || value != nil {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}

func TestCovertBetweenFuncF(t *testing.T) {
	/// Setup
	var errF Func = func(value interface{}) (interface{}, error) {
		return valueOp, errOp
	}

	/// When
	value, err := RetryF(retries).
		ToCallbackFuncF().
		ToCallbackFuncF().
		ToFuncF().
		ToFuncF().
		ToSupplyFuncF().
		ToSupplyFuncF().
		ToFuncF().
		Wrap(errF).
		Invoke(0)

	/// Then
	if err != errOp || value != nil {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}
