package yoti

const (
	attrConstSelfie                  = "selfie"
	attrConstGivenNames              = "given_names"
	attrConstFamilyName              = "family_name"
	attrConstFullName                = "full_name"
	attrConstMobileNumber            = "phone_number"
	attrConstEmailAddress            = "email_address"
	attrConstDateOfBirth             = "date_of_birth"
	attrConstAddress                 = "postal_address"
	attrConstStructuredPostalAddress = "structured_postal_address"
	attrConstGender                  = "gender"
	attrConstNationality             = "nationality"
)

//Profile represents the details retrieved for a particular
type Profile struct {
	attributeSlice []Attribute
}

// AttributeMap represents a map of the Yoti attributes, each attribute is a small piece of information about a Yoti user such as a photo of the user or the
// user's date of birth.
func (p Profile) AttributeMap() (result map[string]Attribute) {
	attributeMap := make(map[string]Attribute)

	for _, v := range p.attributeSlice {
		attributeMap[v.Name()] = v
	}

	return attributeMap
}

// Selfie is a photograph of the user. Will be nil if not provided by Yoti
func (p Profile) Selfie() (result AttributeImage) {
	if p.GetAttribute(attrConstSelfie) != nil {
		return p.GetAttribute(attrConstSelfie).(AttributeImage)
	}

	return
}

// GivenNames represents the user's given names. Will be nil if not provided by Yoti
func (p Profile) GivenNames() (result AttributeString) {
	if p.GetAttribute(attrConstGivenNames) != nil {
		return p.GetAttribute(attrConstGivenNames).(AttributeString)
	}

	return
}

// FamilyName represents the user's family name. Will be nil if not provided by Yoti
func (p Profile) FamilyName() (result AttributeString) {
	if p.GetAttribute(attrConstFamilyName) != nil {
		return p.GetAttribute(attrConstFamilyName).(AttributeString)
	}

	return
}

//FullName represents the user's full name. Will be nil if not provided by Yoti
func (p Profile) FullName() (result AttributeString) {
	if p.GetAttribute(attrConstFullName) != nil {
		return p.GetAttribute(attrConstFullName).(AttributeString)
	}

	return
}

// MobileNumber represents the user's mobile phone number. Will be nil if not provided by Yoti
func (p Profile) MobileNumber() (result AttributeString) {
	if p.GetAttribute(attrConstMobileNumber) != nil {
		return p.GetAttribute(attrConstMobileNumber).(AttributeString)
	}

	return
}

// EmailAddress represents the user's email address. Will be nil if not provided by Yoti
func (p Profile) EmailAddress() (result AttributeString) {
	if p.GetAttribute(attrConstEmailAddress) != nil {
		return p.GetAttribute(attrConstEmailAddress).(AttributeString)
	}

	return
}

// DateOfBirth represents the user's date of birth. Will be nil if not provided by Yoti
func (p Profile) DateOfBirth() (result AttributeTime) {
	if p.GetAttribute(attrConstDateOfBirth) != nil {
		return p.GetAttribute(attrConstDateOfBirth).(AttributeTime)
	}

	return
}

// Address represents the user's address. Will be nil if not provided by Yoti
func (p Profile) Address() (result AttributeString) {
	if p.GetAttribute(attrConstAddress) != nil {
		return p.GetAttribute(attrConstAddress).(AttributeString)
	}

	return
}

// StructuredPostalAddress represents the user's address in a JSON format. Will be empty if not provided by Yoti
func (p Profile) StructuredPostalAddress() (result AttributeJSON) {
	if p.GetAttribute(attrConstStructuredPostalAddress) != nil {
		return p.GetAttribute(attrConstStructuredPostalAddress).(AttributeJSON)
	}

	return
}

// Gender represents the user's gender. Will be nil if not provided by Yoti
func (p Profile) Gender() (result AttributeString) {
	if p.GetAttribute(attrConstGender) != nil {
		return p.GetAttribute(attrConstGender).(AttributeString)
	}

	return
}

// Nationality represents the user's nationality. Will be nil if not provided by Yoti
func (p Profile) Nationality() (result AttributeString) {
	if p.GetAttribute(attrConstNationality) != nil {
		return p.GetAttribute(attrConstNationality).(AttributeString)
	}

	return
}

// GetAttribute retrieve an attribute by name on the Yoti profile. Must be explicitly cast to the desired
// attribute type, e.g. GetAttribute("nationality").(AttributeString)
func (p Profile) GetAttribute(attributeName string) (result Attribute) {
	if attribute, ok := p.AttributeMap()[attributeName]; ok {
		return attribute
	}
	return
}
