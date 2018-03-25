package compose

import "testing"

func TestNoop(t *testing.T) {
	/// Setup
	var errF Function = func() (interface{}, error) {
		return valueOp, errOp
	}

	/// When & Then
	if _, err := errF.Noop().Invoke(); err != errOp {
		t.Errorf("Expected %v, got %v", errOp, err)
	}
}
