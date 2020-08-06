package yotierror

import "fmt"

// TemporaryError indicates that a temporary outage has occured and the
// previous request can be reattempted without modification.
type TemporaryError struct {
	Err error
}

// NewTemporary marks another error as a temporary error
func NewTemporary(err error) TemporaryError {
	return TemporaryError{Err: err}
}

func (e TemporaryError) Error() string {
	return e.Err.Error()
}

// Unwrap returns the internal error for debugging
func (e TemporaryError) Unwrap() error {
	return e.Err
}

// Temporary indicates this error is a temporary error
func (e TemporaryError) Temporary() bool {
	return true
}

// MultiError wraps one or more errors into a single error
type MultiError struct {
	This error
	Next error
}

func (e MultiError) Error() string {
	if e.Next != nil {
		return fmt.Sprintf("%s, %s", e.This.Error(), e.Next.Error())
	}
	return e.This.Error()
}

// Unwrap the next error
func (e MultiError) Unwrap() error {
	return e.Next
}
