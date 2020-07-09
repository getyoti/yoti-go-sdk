package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func TestParseValue_ShouldParseInt(t *testing.T) {
	parsed, err := parseValue(yotiprotoattr.ContentType_INT, []byte("7"))
	assert.NilError(t, err)
	integer, ok := parsed.(int)
	assert.Check(t, ok)
	assert.Equal(t, integer, 7)
}
