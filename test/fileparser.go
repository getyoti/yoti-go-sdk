package test

import (
	"encoding/base64"
	"io/ioutil"
	"testing"

	"gotest.tools/v3/assert"
)

// GetTestFileBytes takes a filepath, decodes it from base64, and returns a byte representation of it
func GetTestFileBytes(t *testing.T, filename string) (result []byte) {
	base64Bytes := readTestFile(t, filename)
	base64String := string(base64Bytes)
	filebytes, err := base64.StdEncoding.DecodeString(base64String)

	assert.NilError(t, err)

	return filebytes
}

// GetTestFileAsString returns a file as a string
func GetTestFileAsString(t *testing.T, filename string) string {
	base64Bytes := readTestFile(t, filename)
	return string(base64Bytes)
}

func readTestFile(t *testing.T, filename string) (result []byte) {
	b, err := ioutil.ReadFile(filename)
	assert.NilError(t, err)

	return b
}
