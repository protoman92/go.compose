package compose

import "testing"

func TestNoop(t *testing.T) {
	/// Setup
	errF := func() error {
		return err
	}

	/// When & Then
	if err1 := NoopError()(errF)(); err1 != err {
		t.Errorf("Expected %v, got %v", err, err1)
	}
}
