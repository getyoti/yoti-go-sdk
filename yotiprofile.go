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
	// AttributeSlice represents a map of the Yoti attributes, each attribute is a small piece of information about a Yoti user such as a photo of the user or the
	// user's date of birth.
	AttributeSlice []*Attribute
}

// Selfie is a photograph of the user. Will be nil if not provided by Yoti
func (p Profile) Selfie() *ImageAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstSelfie {
			return newImageAttribute(a)
		}
	}
	return nil
}

// GivenNames represents the user's given names. Will be nil if not provided by Yoti
func (p Profile) GivenNames() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstGivenNames {
			return newStringAttribute(a)
		}
	}
	return nil
}

// FamilyName represents the user's family name. Will be nil if not provided by Yoti
func (p Profile) FamilyName() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstFamilyName {
			return newStringAttribute(a)
		}
	}
	return nil
}

//FullName represents the user's full name. Will be nil if not provided by Yoti
func (p Profile) FullName() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstFullName {
			return newStringAttribute(a)
		}
	}
	return nil
}

// MobileNumber represents the user's mobile phone number. Will be nil if not provided by Yoti
func (p Profile) MobileNumber() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstMobileNumber {
			return newStringAttribute(a)
		}
	}
	return nil
}

// EmailAddress represents the user's email address. Will be nil if not provided by Yoti
func (p Profile) EmailAddress() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstEmailAddress {
			return newStringAttribute(a)
		}
	}
	return nil
}

// DateOfBirth represents the user's date of birth. Will be nil if not provided by Yoti. Has an err value which will be filled if there is an error parsing the date.
func (p Profile) DateOfBirth() *TimeAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstDateOfBirth {
			return newTimeAttribute(a)
		}
	}
	return nil
}

// Address represents the user's address. Will be nil if not provided by Yoti
func (p Profile) Address() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstAddress {
			return newStringAttribute(a)
		}
	}
	return nil
}

// StructuredPostalAddress represents the user's address in a JSON format. Will be nil if not provided by Yoti
func (p Profile) StructuredPostalAddress() *JSONAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstStructuredPostalAddress {
			return newJSONAttribute(a)
		}
	}
	return nil
}

// Gender represents the user's gender. Will be nil if not provided by Yoti
func (p Profile) Gender() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstGender {
			return newStringAttribute(a)
		}
	}
	return nil
}

// Nationality represents the user's nationality. Will be nil if not provided by Yoti
func (p Profile) Nationality() *StringAttribute {
	for _, a := range p.AttributeSlice {
		if a.Name == attrConstNationality {
			return newStringAttribute(a)
		}
	}
	return nil
}

// GetAttribute retrieve an attribute by name on the Yoti profile. Will return nil if attribute is not present.
func (p Profile) GetAttribute(attributeName string) *GenericAttribute {
	for _, attribute := range p.AttributeSlice {
		if attribute.Name == attributeName {
			return newGenericAttribute(attribute)
		}
	}
	return nil
}
