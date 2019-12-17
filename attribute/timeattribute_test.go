package attribute

import (
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"gotest.tools/assert"
)

func TestTimeAttribute_NewTime_DateOnly(t *testing.T) {
	proto := yotiprotoattr.Attribute{
		Value: []byte("2011-12-25"),
	}

	timeAttribute, err := NewTime(&proto)
	assert.NilError(t, err)

	assert.Equal(t, *timeAttribute.Value(), time.Date(2011, 12, 25, 0, 0, 0, 0, time.UTC))
}
