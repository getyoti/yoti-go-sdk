package test

import (
	"encoding/base64"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

// DecodeTestFile reads a test fixture file
func DecodeTestFile(t *testing.T, filename string) (result []byte) {
	base64Bytes := readTestFile(t, filename)
	base64String := string(base64Bytes)
	filebytes, err := base64.StdEncoding.DecodeString(base64String)

	assert.Assert(t, is.Nil(err))

	return filebytes
}
