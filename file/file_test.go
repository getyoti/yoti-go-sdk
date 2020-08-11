package file

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFile_ReadFile(t *testing.T) {
	_, err := ReadFile("../test/test-key.pem")
	assert.NilError(t, err)
}

func TestFile_ReadFile_ShouldFailForFileNotFound(t *testing.T) {
	MissingFileName := "/tmp/file_not_found"
	_, err := ReadFile(MissingFileName)
	assert.Check(t, err != nil)
}
