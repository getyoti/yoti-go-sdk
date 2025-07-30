package digitalidentity

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func TestUserProfile_EstimatedAge_ReturnsNilWhenNotPresent(t *testing.T) {
	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	result := userProfile.EstimatedAge()

	assert.Assert(t, result == nil)
}

func TestUserProfile_EstimatedAge_ReturnsEstimatedAgeWhenPresent(t *testing.T) {
	estimatedAgeValue := "18-24"
	estimatedAgeAttr := &yotiprotoattr.Attribute{
		Name:        consts.AttrEstimatedAge,
		Value:       []byte(estimatedAgeValue),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{estimatedAgeAttr},
		},
	}

	result := userProfile.EstimatedAge()

	assert.Assert(t, result != nil)
	assert.Equal(t, result.Value(), estimatedAgeValue)
}

func TestUserProfile_EstimatedAgeWithFallback_ReturnsEstimatedAge_WhenPresent(t *testing.T) {
	estimatedAgeValue := "18-24"
	estimatedAgeAttr := &yotiprotoattr.Attribute{
		Name:        consts.AttrEstimatedAge,
		Value:       []byte(estimatedAgeValue),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{estimatedAgeAttr},
		},
	}

	result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()

	assert.Assert(t, result != nil)
	assert.Assert(t, isEstimatedAge) // true indicates estimated_age was found
	stringAttr, ok := result.(*attribute.StringAttribute)
	assert.Assert(t, ok)
	assert.Equal(t, stringAttr.Value(), estimatedAgeValue)
}

func TestUserProfile_EstimatedAgeWithFallback_ReturnsDateOfBirth_WhenEstimatedAgeNotPresent(t *testing.T) {
	dateOfBirthValue := "1985-03-15"
	dateOfBirthAttr := &yotiprotoattr.Attribute{
		Name:        consts.AttrDateOfBirth,
		Value:       []byte(dateOfBirthValue),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{dateOfBirthAttr},
		},
	}

	result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()

	assert.Assert(t, result != nil)
	assert.Assert(t, !isEstimatedAge) // false indicates date_of_birth was used as fallback
	dateAttr, ok := result.(*attribute.DateAttribute)
	assert.Assert(t, ok)
	assert.Equal(t, dateAttr.Value().Format("2006-01-02"), dateOfBirthValue)
}

func TestUserProfile_EstimatedAgeWithFallback_PrefersEstimatedAge_WhenBothPresent(t *testing.T) {
	estimatedAgeValue := "18-24"
	dateOfBirthValue := "1985-03-15"

	estimatedAgeAttr := &yotiprotoattr.Attribute{
		Name:        consts.AttrEstimatedAge,
		Value:       []byte(estimatedAgeValue),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	dateOfBirthAttr := &yotiprotoattr.Attribute{
		Name:        consts.AttrDateOfBirth,
		Value:       []byte(dateOfBirthValue),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{estimatedAgeAttr, dateOfBirthAttr},
		},
	}

	result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()

	assert.Assert(t, result != nil)
	assert.Assert(t, isEstimatedAge) // true indicates estimated_age was preferred
	stringAttr, ok := result.(*attribute.StringAttribute)
	assert.Assert(t, ok)
	assert.Equal(t, stringAttr.Value(), estimatedAgeValue)
}

func TestUserProfile_EstimatedAgeWithFallback_ReturnsNil_WhenNeitherPresent(t *testing.T) {
	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	result, isEstimatedAge := userProfile.EstimatedAgeWithFallback()

	assert.Assert(t, result == nil)
	assert.Assert(t, !isEstimatedAge)
}
