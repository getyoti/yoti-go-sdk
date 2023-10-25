package digitalidentity

import (
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// UserProfile represents the details retrieved for a particular user. Consists of
// Yoti attributes: a small piece of information about a Yoti user such as a
// photo of the user or the user's date of birth.
type UserProfile struct {
	baseProfile
}

// Creates a new Profile struct
func newUserProfile(attributes *yotiprotoattr.AttributeList) UserProfile {
	return UserProfile{
		baseProfile{
			attributeSlice: createAttributeSlice(attributes),
		},
	}
}

func createAttributeSlice(protoAttributeList *yotiprotoattr.AttributeList) (result []*yotiprotoattr.Attribute) {
	if protoAttributeList != nil {
		result = append(result, protoAttributeList.Attributes...)
	}

	return result
}

// Selfie is a photograph of the user. Will be nil if not provided by Yoti.
func (p UserProfile) Selfie() *attribute.ImageAttribute {
	return p.GetImageAttribute(consts.AttrSelfie)
}

// GetSelfieAttributeByID retrieve a Selfie attribute by ID on the Yoti profile.
// This attribute is a photograph of the user.
// Will return nil if attribute is not present.
func (p UserProfile) GetSelfieAttributeByID(attributeID string) (*attribute.ImageAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.EphemeralId == attributeID {
			return attribute.NewImage(a)
		}
	}
	return nil, nil
}

// GivenNames corresponds to secondary names in passport, and first/middle names in English. Will be nil if not provided by Yoti.
func (p UserProfile) GivenNames() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrGivenNames)
}

// FamilyName corresponds to primary name in passport, and surname in English. Will be nil if not provided by Yoti.
func (p UserProfile) FamilyName() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrFamilyName)
}

// FullName represents the user's full name.
// If family_name/given_names are present, the value will be equal to the string 'given_names + " " family_name'.
// Will be nil if not provided by Yoti.
func (p UserProfile) FullName() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrFullName)
}

// MobileNumber represents the user's mobile phone number, as verified at registration time.
// The value will be a number in E.164 format (i.e. '+' for international prefix and no spaces, e.g. "+447777123456").
// Will be nil if not provided by Yoti.
func (p UserProfile) MobileNumber() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrMobileNumber)
}

// EmailAddress represents the user's verified email address. Will be nil if not provided by Yoti.
func (p UserProfile) EmailAddress() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrEmailAddress)
}

// DateOfBirth represents the user's date of birth. Will be nil if not provided by Yoti.
// Has an err value which will be filled if there is an error parsing the date.
func (p UserProfile) DateOfBirth() (*attribute.DateAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.Name == consts.AttrDateOfBirth {
			return attribute.NewDate(a)
		}
	}
	return nil, nil
}

// Address represents the user's address. Will be nil if not provided by Yoti.
func (p UserProfile) Address() *attribute.StringAttribute {
	addressAttribute := p.GetStringAttribute(consts.AttrAddress)
	if addressAttribute == nil {
		return ensureAddressProfile(&p)
	}

	return addressAttribute
}

// StructuredPostalAddress represents the user's address in a JSON format.
// Will be nil if not provided by Yoti. This can be accessed as a
// map[string]string{} using a type assertion, e.g.:
// structuredPostalAddress := structuredPostalAddressAttribute.Value().(map[string]string{})
func (p UserProfile) StructuredPostalAddress() (*attribute.JSONAttribute, error) {
	return p.GetJSONAttribute(consts.AttrStructuredPostalAddress)
}

// Gender corresponds to the gender in the registered document; the value will be one of the strings "MALE", "FEMALE", "TRANSGENDER" or "OTHER".
// Will be nil if not provided by Yoti.
func (p UserProfile) Gender() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrGender)
}

// Nationality corresponds to the nationality in the passport.
// The value is an ISO-3166-1 alpha-3 code with ICAO9303 (passport) extensions.
// Will be nil if not provided by Yoti.
func (p UserProfile) Nationality() *attribute.StringAttribute {
	return p.GetStringAttribute(consts.AttrNationality)
}

// DocumentImages returns a slice of document images cropped from the image in the capture page.
// There can be multiple images as per the number of regions in the capture in this attribute.
// Will be nil if not provided by Yoti.
func (p UserProfile) DocumentImages() (*attribute.ImageSliceAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.Name == consts.AttrDocumentImages {
			return attribute.NewImageSlice(a)
		}
	}
	return nil, nil
}

// GetDocumentImagesAttributeByID retrieve a Document Images attribute by ID on the Yoti profile.
// This attribute consists of a slice of document images cropped from the image in the capture page.
// There can be multiple images as per the number of regions in the capture in this attribute.
// Will return nil if attribute is not present.
func (p UserProfile) GetDocumentImagesAttributeByID(attributeID string) (*attribute.ImageSliceAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.EphemeralId == attributeID {
			return attribute.NewImageSlice(a)
		}
	}
	return nil, nil
}

// DocumentDetails represents information extracted from a document provided by the user.
// Will be nil if not provided by Yoti.
func (p UserProfile) DocumentDetails() (*attribute.DocumentDetailsAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.Name == consts.AttrDocumentDetails {
			return attribute.NewDocumentDetails(a)
		}
	}
	return nil, nil
}

// IdentityProfileReport represents the JSON object containing identity assertion and the
// verification report. Will be nil if not provided by Yoti.
func (p UserProfile) IdentityProfileReport() (*attribute.JSONAttribute, error) {
	return p.GetJSONAttribute(consts.AttrIdentityProfileReport)
}

// AgeVerifications returns a slice of age verifications for the user.
// Will be an empty slice if not provided by Yoti.
func (p UserProfile) AgeVerifications() (out []attribute.AgeVerification, err error) {
	ageUnderString := strings.Replace(consts.AttrAgeUnder, "%d", "", -1)
	ageOverString := strings.Replace(consts.AttrAgeOver, "%d", "", -1)

	for _, a := range p.attributeSlice {
		if strings.HasPrefix(a.Name, ageUnderString) ||
			strings.HasPrefix(a.Name, ageOverString) {
			verification, err := attribute.NewAgeVerification(a)
			if err != nil {
				return nil, err
			}
			out = append(out, verification)
		}
	}
	return out, err
}
