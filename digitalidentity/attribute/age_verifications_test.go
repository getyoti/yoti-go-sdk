package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func TestNewAgeVerification_ValueTrue(t *testing.T) {
	attribute := &yotiprotoattr.Attribute{
		Name:        "age_over:18",
		Value:       []byte("true"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	ageVerification, err := NewAgeVerification(attribute)

	assert.NilError(t, err)

	assert.Equal(t, ageVerification.Age, 18)
	assert.Equal(t, ageVerification.CheckType, "age_over")
	assert.Equal(t, ageVerification.Result, true)
}

func TestNewAgeVerification_ValueFalse(t *testing.T) {
	attribute := &yotiprotoattr.Attribute{
		Name:        "age_under:30",
		Value:       []byte("false"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	ageVerification, err := NewAgeVerification(attribute)

	assert.NilError(t, err)

	assert.Equal(t, ageVerification.Age, 30)
	assert.Equal(t, ageVerification.CheckType, "age_under")
	assert.Equal(t, ageVerification.Result, false)
}
