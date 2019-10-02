package sandbox

import (
	"encoding/base64"
	"encoding/json"
	"time"

	yoti "github.com/getyoti/yoti-go-sdk/v2"
)

type Profile struct {
	RememberMeID string      `json:"remember_me_id"`
	Attributes   []Attribute `json:"profile_attributes"`
}

func (profile *Profile) WithAttribute(name, value string, anchors []Anchor) *Profile {
	if anchors == nil {
		anchors = make([]Anchor, 0)
	}
	attribute := Attribute{
		Name:    name,
		Value:   value,
		Anchors: anchors,
	}
	profile.Attributes = append(profile.Attributes, attribute)
	return profile
}

func (profile *Profile) WithGivenNames(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstGivenNames, value, anchors)
}

func (profile *Profile) WithFamilyName(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstFamilyName, value, anchors)
}

func (profile *Profile) WithFullName(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstFullName, value, anchors)
}

func (profile *Profile) WithDateOfBirth(value time.Time, anchors []Anchor) *Profile {
	formattedTime := value.Format("2006-01-02")
	return profile.WithAttribute(yoti.AttrConstDateOfBirth, formattedTime, anchors)
}

func (profile *Profile) WithAgeVerification(dateOfBirth time.Time, derivation Derivation, anchors []Anchor) *Profile {
	if anchors == nil {
		anchors = []Anchor{}
	}
	attribute := Attribute{
		Name:       yoti.AttrConstDateOfBirth,
		Value:      dateOfBirth.Format("2006-01-02"),
		Derivation: derivation.ToString(),
		Anchors:    anchors,
	}
	profile.Attributes = append(profile.Attributes, attribute)
	return profile
}

func (profile *Profile) WithGender(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstGender, value, anchors)
}

func (profile *Profile) WithPhoneNumber(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstMobileNumber, value, anchors)
}

func (profile *Profile) WithNationality(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstNationality, value, anchors)
}

func (profile *Profile) WithPostalAddress(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstAddress, value, anchors)
}

func (profile *Profile) WithStructuredPostalAddress(value map[string]string, anchors []Anchor) *Profile {
	data, _ := json.Marshal(value)
	return profile.WithAttribute(yoti.AttrConstStructuredPostalAddress, string(data), anchors)
}

func (profile *Profile) WithSelfie(value []byte, anchors []Anchor) *Profile {
	return profile.WithAttribute(
		yoti.AttrConstSelfie,
		base64.StdEncoding.EncodeToString(value),
		anchors,
	)
}

func (profile *Profile) WithEmailAddress(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstEmailAddress, value, anchors)
}

func (profile *Profile) WithDocumentDetails(value string, anchors []Anchor) *Profile {
	return profile.WithAttribute(yoti.AttrConstDocumentDetails, value, anchors)
}

func (profile *Profile) WithoutAttributes() *Profile {
	profile.Attributes = []Attribute{Attribute{
		Name:    "unused",
		Value:   "unused",
		Anchors: make([]Anchor, 0),
	}}
	return profile
}
