package attribute

import (
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestTimeAttribute_NewTime_DateOnly(t *testing.T) {
	proto := yotiprotoattr.Attribute{
		Value: []byte("2011-12-25"),
	}

	timeAttribute, err := NewTime(&proto)
	assert.NilError(t, err)

	assert.Equal(t, *timeAttribute.Value(), time.Date(2011, 12, 25, 0, 0, 0, 0, time.UTC))
}

func TestTimeAttribute_DateOfBirth(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../test/fixtures/test_attribute_date_of_birth.txt")

	dateOfBirthAttribute, err := NewTime(protoAttribute)

	assert.Assert(t, is.Nil(err))

	expectedDateOfBirth := time.Date(1970, time.December, 01, 0, 0, 0, 0, time.UTC)
	actualDateOfBirth := dateOfBirthAttribute.Value()

	assert.Assert(t, actualDateOfBirth.Equal(expectedDateOfBirth))
}
