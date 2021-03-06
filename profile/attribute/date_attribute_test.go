package attribute

import (
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func TestTimeAttribute_NewDate_DateOnly(t *testing.T) {
	proto := yotiprotoattr.Attribute{
		Value: []byte("2011-12-25"),
	}

	timeAttribute, err := NewDate(&proto)
	assert.NilError(t, err)

	assert.Equal(t, *timeAttribute.Value(), time.Date(2011, 12, 25, 0, 0, 0, 0, time.UTC))
}

func TestTimeAttribute_DateOfBirth(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../../test/fixtures/test_attribute_date_of_birth.txt")

	dateOfBirthAttribute, err := NewDate(protoAttribute)

	assert.NilError(t, err)

	expectedDateOfBirth := time.Date(1970, time.December, 01, 0, 0, 0, 0, time.UTC)
	actualDateOfBirth := dateOfBirthAttribute.Value()

	assert.Assert(t, actualDateOfBirth.Equal(expectedDateOfBirth))
}

func TestNewTime_ShouldReturnErrorForInvalidDate(t *testing.T) {
	proto := yotiprotoattr.Attribute{
		Name:        "example",
		Value:       []byte("2006-60-20"),
		ContentType: yotiprotoattr.ContentType_DATE,
	}
	attribute, err := NewDate(&proto)
	assert.Check(t, attribute == nil)
	assert.ErrorContains(t, err, "month out of range")
}
