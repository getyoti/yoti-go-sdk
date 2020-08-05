package yotierror

import (
	"errors"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMultiError(t *testing.T) {
	err := errors.New("inner err")
	err = MultiError{This: errors.New("outer err"), Next: err}
	result := err.(MultiError)

	assert.Equal(t, result.Error(), "outer err, inner err")
	assert.Error(t, result.Unwrap(), "inner err")
}
