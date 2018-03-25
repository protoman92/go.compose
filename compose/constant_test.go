package compose

import (
	"errors"
	"time"
)

const (
	delay   = time.Duration(1e8)
	retries = uint(10)

	// This value should be returned by test Functions.
	valueOp = 1
)

var (
	// This error should be returned by test Functions.
	errOp = errors.New("error")
)
