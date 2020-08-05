package dynamic

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

const (
	authTypeSelfieConst = iota + 1
	authTypePinConst
)

// PolicyBuilder constructs a json payload specifying the dynamic policy
// for a dynamic scenario
type PolicyBuilder struct {
	wantedAttributes   map[string]WantedAttribute
	wantedAuthTypes    map[int]bool
	isWantedRememberMe bool
	err                error
}

// Policy represents a dynamic policy for a share
type Policy struct {
	attributes   []WantedAttribute
	authTypes    []int
	rememberMeID bool
}

// WithWantedAttribute adds an attribute from WantedAttributeBuilder to the policy
func (b *PolicyBuilder) WithWantedAttribute(attribute WantedAttribute) *PolicyBuilder {
	if b.wantedAttributes == nil {
		b.wantedAttributes = make(map[string]WantedAttribute)
	}
	var key string
	if attribute.derivation != "" {
		key = attribute.derivation
	} else {
		key = attribute.name
	}
	b.wantedAttributes[key] = attribute
	return b
}

// WithWantedAttributeByName adds an attribute by its name. This is not the preferred
// way of adding an attribute - instead use the other methods below
func (b *PolicyBuilder) WithWantedAttributeByName(name string, options ...interface{}) *PolicyBuilder {
	attributeBuilder := (&WantedAttributeBuilder{}).WithName(name)

	for _, option := range options {
		switch value := option.(type) {
		case SourceConstraint:
			attributeBuilder.WithConstraint(&value)
		case constraintInterface:
			attributeBuilder.WithConstraint(value)
		default:
			panic(fmt.Sprintf("Not a valid option type, %v", value))
		}
	}

	attribute, err := attributeBuilder.Build()
	if err != nil {
		b.err = yotierror.MultiError{This: err, Next: b.err}
	}
	b.WithWantedAttribute(attribute)
	return b
}

// WithFamilyName adds the family name attribute
func (b *PolicyBuilder) WithFamilyName(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrFamilyName, options...)
}

// WithGivenNames adds the given names attribute
func (b *PolicyBuilder) WithGivenNames(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrGivenNames, options...)
}

// WithFullName adds the full name attribute
func (b *PolicyBuilder) WithFullName(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrFullName, options...)
}

// WithDateOfBirth adds the date of birth attribute
func (b *PolicyBuilder) WithDateOfBirth(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrDateOfBirth, options...)
}

// WithGender adds the gender attribute
func (b *PolicyBuilder) WithGender(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrGender, options...)
}

// WithPostalAddress adds the postal address attribute
func (b *PolicyBuilder) WithPostalAddress(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrAddress, options...)
}

// WithStructuredPostalAddress adds the structured postal address attribute
func (b *PolicyBuilder) WithStructuredPostalAddress(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrStructuredPostalAddress, options...)
}

// WithNationality adds the nationality attribute
func (b *PolicyBuilder) WithNationality(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrNationality, options...)
}

// WithPhoneNumber adds the phone number attribute
func (b *PolicyBuilder) WithPhoneNumber(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrMobileNumber, options...)
}

// WithSelfie adds the selfie attribute
func (b *PolicyBuilder) WithSelfie(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrSelfie, options...)
}

// WithEmail adds the email address attribute
func (b *PolicyBuilder) WithEmail(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrEmailAddress, options...)
}

// WithDocumentImages adds the document images attribute
func (b *PolicyBuilder) WithDocumentImages(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrDocumentImages, options...)
}

// WithDocumentDetails adds the document details attribute
func (b *PolicyBuilder) WithDocumentDetails(options ...interface{}) *PolicyBuilder {
	return b.WithWantedAttributeByName(consts.AttrDocumentDetails, options...)
}

// WithAgeDerivedAttribute is a helper method for setting age based derivations
// Prefer to use WithAgeOver and WithAgeUnder instead of using this directly
func (b *PolicyBuilder) WithAgeDerivedAttribute(derivation string, options ...interface{}) *PolicyBuilder {
	var attributeBuilder WantedAttributeBuilder
	attributeBuilder.
		WithName(consts.AttrDateOfBirth).
		WithDerivation(derivation)

	for _, option := range options {
		switch value := option.(type) {
		case SourceConstraint:
			attributeBuilder.WithConstraint(&value)
		case constraintInterface:
			attributeBuilder.WithConstraint(value)
		default:
			panic(fmt.Sprintf("Not a valid option type, %v", value))
		}
	}

	attr, err := attributeBuilder.Build()
	if err != nil {
		b.err = yotierror.MultiError{This: err, Next: b.err}
	}
	return b.WithWantedAttribute(attr)
}

// WithAgeOver sets this dynamic policy as requesting whether the user is older
// than a certain age
func (b *PolicyBuilder) WithAgeOver(age int, options ...interface{}) *PolicyBuilder {
	return b.WithAgeDerivedAttribute(fmt.Sprintf(consts.AttrAgeOver, age), options...)
}

// WithAgeUnder sets this dynamic policy as requesting whether the user is younger
// than a certain age
func (b *PolicyBuilder) WithAgeUnder(age int, options ...interface{}) *PolicyBuilder {
	return b.WithAgeDerivedAttribute(fmt.Sprintf(consts.AttrAgeUnder, age), options...)
}

// WithWantedRememberMe sets the Policy as requiring a "Remember Me ID"
func (b *PolicyBuilder) WithWantedRememberMe() *PolicyBuilder {
	b.isWantedRememberMe = true
	return b
}

// WithWantedAuthType sets this dynamic policy as requiring a specific authentication type
func (b *PolicyBuilder) WithWantedAuthType(wantedAuthType int) *PolicyBuilder {
	if b.wantedAuthTypes == nil {
		b.wantedAuthTypes = make(map[int]bool)
	}
	b.wantedAuthTypes[wantedAuthType] = true
	return b
}

// WithSelfieAuth sets this dynamic policy as requiring Selfie-based authentication
func (b *PolicyBuilder) WithSelfieAuth() *PolicyBuilder {
	return b.WithWantedAuthType(authTypeSelfieConst)
}

// WithPinAuth sets this dynamic policy as requiring PIN authentication
func (b *PolicyBuilder) WithPinAuth() *PolicyBuilder {
	return b.WithWantedAuthType(authTypePinConst)
}

// Build constructs a dynamic policy object
func (b *PolicyBuilder) Build() (Policy, error) {
	return Policy{
		attributes:   b.attributesAsList(),
		authTypes:    b.authTypesAsList(),
		rememberMeID: b.isWantedRememberMe,
	}, b.err
}

func (b *PolicyBuilder) attributesAsList() []WantedAttribute {
	attributeList := make([]WantedAttribute, 0)
	for _, attr := range b.wantedAttributes {
		attributeList = append(attributeList, attr)
	}
	return attributeList
}

func (b *PolicyBuilder) authTypesAsList() []int {
	authTypeList := make([]int, 0)
	for auth, boolValue := range b.wantedAuthTypes {
		if boolValue {
			authTypeList = append(authTypeList, auth)
		}
	}
	return authTypeList
}

// MarshalJSON returns the JSON encoding
func (policy *Policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Wanted           []WantedAttribute `json:"wanted"`
		WantedAuthTypes  []int             `json:"wanted_auth_types"`
		WantedRememberMe bool              `json:"wanted_remember_me"`
	}{
		Wanted:           policy.attributes,
		WantedAuthTypes:  policy.authTypes,
		WantedRememberMe: policy.rememberMeID,
	})
}
