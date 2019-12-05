package yoti

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"gotest.tools/assert"
)

func TestProfile_AgeVerifications(t *testing.T) {
	ageOver14 := &yotiprotoattr.Attribute{
		Name:        "age_over:14",
		Value:       []byte("true"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	ageUnder18 := &yotiprotoattr.Attribute{
		Name:        "age_under:18",
		Value:       []byte("true"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	ageOver18 := &yotiprotoattr.Attribute{
		Name:        "age_over:18",
		Value:       []byte("false"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithMultipleAttributes(ageOver14, ageUnder18, ageOver18)
	ageVerifications, err := profile.AgeVerifications()

	assert.NilError(t, err)
	assert.Equal(t, len(ageVerifications), 3)

	assert.Equal(t, ageVerifications[0].Age, 14)
	assert.Equal(t, ageVerifications[0].CheckType, "age_over")
	assert.Equal(t, ageVerifications[0].Result, true)

	assert.Equal(t, ageVerifications[1].Age, 18)
	assert.Equal(t, ageVerifications[1].CheckType, "age_under")
	assert.Equal(t, ageVerifications[1].Result, true)

	assert.Equal(t, ageVerifications[2].Age, 18)
	assert.Equal(t, ageVerifications[2].CheckType, "age_over")
	assert.Equal(t, ageVerifications[2].Result, false)
}
