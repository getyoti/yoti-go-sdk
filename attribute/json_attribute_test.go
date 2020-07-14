package attribute

import (
	"fmt"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func ExampleNewJSON() {
	proto := yotiprotoattr.Attribute{
		Name:        "exampleJSON",
		Value:       []byte(`{"foo":"bar"}`),
		ContentType: yotiprotoattr.ContentType_JSON,
	}
	attribute, err := NewJSON(&proto)
	if err != nil {
		return
	}
	fmt.Println(attribute.Value())
	// Output: map[foo:bar]
}

func TestNewJSON_ShouldReturnNilForInvalidJSON(t *testing.T) {
	proto := yotiprotoattr.Attribute{
		Name:        "exampleJSON",
		Value:       []byte("Not a json document"),
		ContentType: yotiprotoattr.ContentType_JSON,
	}
	attribute, err := NewJSON(&proto)
	assert.Check(t, attribute == nil)
	assert.ErrorContains(t, err, "Unable to parse JSON value")
}
