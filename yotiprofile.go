package yoti

import (
	"github.com/getyoti/yoti-go-sdk/v2/attribute"
)

// Attribute names for user profile attributes
const (
	AttrConstSelfie                  = "selfie"
	AttrConstGivenNames              = "given_names"
	AttrConstFamilyName              = "family_name"
	AttrConstFullName                = "full_name"
	AttrConstMobileNumber            = "phone_number"
	AttrConstEmailAddress            = "email_address"
	AttrConstDateOfBirth             = "date_of_birth"
	AttrConstAddress                 = "postal_address"
	AttrConstStructuredPostalAddress = "structured_postal_address"
	AttrConstGender                  = "gender"
	AttrConstNationality             = "nationality"
	AttrConstDocumentImages          = "document_images"
	AttrConstDocumentDetails         = "document_details"
	AttrConstAgeOver                 = "age_over:%d"
	AttrConstAgeUnder                = "age_under:%d"
)

// Profile represents the details retrieved for a particular user. Consists of
// Yoti attributes: a small piece of information about a Yoti user such as a
// photo of the user or the user's date of birth.
type Profile struct {
	baseProfile
}

// Selfie is a photograph of the user. Will be nil if not provided by Yoti.
func (p Profile) Selfie() *attribute.ImageAttribute {
	return p.GetImageAttribute(AttrConstSelfie)
}

// GivenNames corresponds to secondary names in passport, and first/middle names in English. Will be nil if not provided by Yoti.
func (p Profile) GivenNames() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstGivenNames)
}

// FamilyName corresponds to primary name in passport, and surname in English. Will be nil if not provided by Yoti.
func (p Profile) FamilyName() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstFamilyName)
}

// FullName represents the user's full name.
// If family_name/given_names are present, the value will be equal to the string 'given_names + " " family_name'.
// Will be nil if not provided by Yoti.
func (p Profile) FullName() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstFullName)
}

// MobileNumber represents the user's mobile phone number, as verified at registration time.
// The value will be a number in E.164 format (i.e. '+' for international prefix and no spaces, e.g. "+447777123456").
// Will be nil if not provided by Yoti.
func (p Profile) MobileNumber() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstMobileNumber)
}

// EmailAddress represents the user's verified email address. Will be nil if not provided by Yoti.
func (p Profile) EmailAddress() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstEmailAddress)
}

// DateOfBirth represents the user's date of birth. Will be nil if not provided by Yoti.
// Has an err value which will be filled if there is an error parsing the date.
func (p Profile) DateOfBirth() (*attribute.TimeAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.Name == AttrConstDateOfBirth {
			return attribute.NewTime(a)
		}
	}
	return nil, nil
}

// Address represents the user's address. Will be nil if not provided by Yoti.
func (p Profile) Address() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstAddress)
}

// StructuredPostalAddress represents the user's address in a JSON format.
// Will be nil if not provided by Yoti. This can be accessed as a
// map[string]string{} using a type assertion, e.g.:
// structuredPostalAddress := structuredPostalAddressAttribute.Value().(map[string]string{})
func (p Profile) StructuredPostalAddress() (*attribute.JSONAttribute, error) {
	return p.GetJSONAttribute(AttrConstStructuredPostalAddress)
}

// Gender corresponds to the gender in the registered document; the value will be one of the strings "MALE", "FEMALE", "TRANSGENDER" or "OTHER".
// Will be nil if not provided by Yoti.
func (p Profile) Gender() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstGender)
}

// Nationality corresponds to the nationality in the passport.
// The value is an ISO-3166-1 alpha-3 code with ICAO9303 (passport) extensions.
// Will be nil if not provided by Yoti.
func (p Profile) Nationality() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstNationality)
}

// DocumentImages returns a slice of document images cropped from the image in the capture page.
// There can be multiple images as per the number of regions in the capture in this attribute.
// Will be nil if not provided by Yoti.
func (p Profile) DocumentImages() (*attribute.ImageSliceAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.Name == AttrConstDocumentImages {
			return attribute.NewImageSlice(a)
		}
	}
	return nil, nil
}

// DocumentDetails represents information extracted from a document provided by the user
func (p Profile) DocumentDetails() (*attribute.DocumentDetailsAttribute, error) {
	for _, a := range p.attributeSlice {
		if a.Name == AttrConstDocumentDetails {
			return attribute.NewDocumentDetails(a)
		}
	}
	return nil, nil
}

// GetAttribute retrieve an attribute by name on the Yoti profile. Will return nil if attribute is not present.
func (p Profile) GetAttribute(attributeName string) *attribute.GenericAttribute {
	for _, a := range p.attributeSlice {
		if a.Name == attributeName {
			return attribute.NewGeneric(a)
		}
	}
	return nil
}
