package compose

import "testing"

func TestConvertToCallbackFunc(t *testing.T) {
	/// Setup
	var errF Func = func(value interface{}) (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	if err := errF.ToCallbackFunc().Invoke(nil); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}
