package compose

import "testing"

func TestConvertToErrorFunc(t *testing.T) {
	/// Setup
	var errF Func = func() (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	if err := errF.ErrorFunc().Invoke(); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}
