package compose

import "testing"

func TestNoop(t *testing.T) {
	/// Setup
	var errF Func = func(value interface{}) (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	if _, err := errF.Noop().Invoke(nil); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}
