package yotierror

import "fmt"

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
