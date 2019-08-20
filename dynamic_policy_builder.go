package yoti

import (
	"encoding/json"
	"fmt"
)

const (
	authTypeSelfieConst = iota + 1
	authTypePinConst
)

// DynamicPolicyBuilder constructs a json payload specifying the dynamic policy
// for a dynamic scenario
type DynamicPolicyBuilder struct {
	wantedAttributes   map[string]WantedAttribute
	wantedAuthTypes    map[int]bool
	isWantedRememberMe bool
}

// DynamicPolicy represents a dynamic policy for a share
type DynamicPolicy struct {
	attributes   []WantedAttribute
	authTypes    []int
	rememberMeID bool
}

// New initializes a DynamicPolicyBuilder
func (b *DynamicPolicyBuilder) New() *DynamicPolicyBuilder {
	b.wantedAttributes = make(map[string]WantedAttribute)
	b.wantedAuthTypes = make(map[int]bool)
	b.isWantedRememberMe = false
	return b
}

// WithWantedAttribute adds an attribute from WantedAttributeBuilder to the policy
func (b *DynamicPolicyBuilder) WithWantedAttribute(attribute WantedAttribute) *DynamicPolicyBuilder {
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
func (b *DynamicPolicyBuilder) WithWantedAttributeByName(name string, options ...interface{}) *DynamicPolicyBuilder {
	attributeBuilder := (&WantedAttributeBuilder{}).New().WithName(name)

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

	attribute := attributeBuilder.Build()
	b.WithWantedAttribute(attribute)
	return b
}

// WithFamilyName adds the family name attribute
func (b *DynamicPolicyBuilder) WithFamilyName(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstFamilyName, options...)
}

// WithGivenNames adds the given names attribute
func (b *DynamicPolicyBuilder) WithGivenNames(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstGivenNames, options...)
}

// WithFullName adds the full name attribute
func (b *DynamicPolicyBuilder) WithFullName(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstFullName, options...)
}

// WithDateOfBirth adds the date of birth attribute
func (b *DynamicPolicyBuilder) WithDateOfBirth(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstDateOfBirth, options...)
}

// WithGender adds the gender attribute
func (b *DynamicPolicyBuilder) WithGender(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstGender, options...)
}

// WithPostalAddress adds the postal address attribute
func (b *DynamicPolicyBuilder) WithPostalAddress(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstAddress, options...)
}

// WithStructuredPostalAddress adds the structured postal address attribute
func (b *DynamicPolicyBuilder) WithStructuredPostalAddress(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstStructuredPostalAddress, options...)
}

// WithNationality adds the nationality attribute
func (b *DynamicPolicyBuilder) WithNationality(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstNationality, options...)
}

// WithPhoneNumber adds the phone number attribute
func (b *DynamicPolicyBuilder) WithPhoneNumber(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstMobileNumber, options...)
}

// WithSelfie adds the selfie attribute
func (b *DynamicPolicyBuilder) WithSelfie(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstSelfie, options...)
}

// WithEmail adds the email address attribute
func (b *DynamicPolicyBuilder) WithEmail(options ...interface{}) *DynamicPolicyBuilder {
	return b.WithWantedAttributeByName(AttrConstEmailAddress, options...)
}

// WithAgeDerivedAttribute is a helper method for setting age based derivations
// Prefer to use WithAgeOver and WithAgeUnder instead of using this directly
func (b *DynamicPolicyBuilder) WithAgeDerivedAttribute(derivation string, options ...interface{}) *DynamicPolicyBuilder {
	var attributeBuilder WantedAttributeBuilder
	attributeBuilder.New().
		WithName(AttrConstDateOfBirth).
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

	return b.WithWantedAttribute(attributeBuilder.Build())
}

// WithAgeOver sets this dynamic policy as requesting whether the user is older
// than a certain age
func (b *DynamicPolicyBuilder) WithAgeOver(age int, options ...interface{}) *DynamicPolicyBuilder {
	return b.WithAgeDerivedAttribute(fmt.Sprintf(AttrConstAgeOver, age), options...)
}

// WithAgeUnder sets this dynamic policy as requesting whether the user is younger
// than a certain age
func (b *DynamicPolicyBuilder) WithAgeUnder(age int, options ...interface{}) *DynamicPolicyBuilder {
	return b.WithAgeDerivedAttribute(fmt.Sprintf(AttrConstAgeUnder, age), options...)
}

// WithWantedRememberMe sets the Policy as requiring a "Remember Me ID"
func (b *DynamicPolicyBuilder) WithWantedRememberMe() *DynamicPolicyBuilder {
	b.isWantedRememberMe = true
	return b
}

// WithWantedAuthType sets this dynamic policy as requiring a specific authentication type
func (b *DynamicPolicyBuilder) WithWantedAuthType(wantedAuthType int) *DynamicPolicyBuilder {
	b.wantedAuthTypes[wantedAuthType] = true
	return b
}

// WithSelfieAuth sets this dynamic policy as requiring Selfie-based authentication
func (b *DynamicPolicyBuilder) WithSelfieAuth() *DynamicPolicyBuilder {
	return b.WithWantedAuthType(authTypeSelfieConst)
}

// WithPinAuth sets this dynamic policy as requiring PIN authentication
func (b *DynamicPolicyBuilder) WithPinAuth() *DynamicPolicyBuilder {
	return b.WithWantedAuthType(authTypePinConst)
}

// Build constructs a dynamic policy object
func (b *DynamicPolicyBuilder) Build() DynamicPolicy {
	return DynamicPolicy{
		attributes:   b.attributesAsList(),
		authTypes:    b.authTypesAsList(),
		rememberMeID: b.isWantedRememberMe,
	}

}

func (b *DynamicPolicyBuilder) attributesAsList() []WantedAttribute {
	attributeList := make([]WantedAttribute, 0)
	for _, attr := range b.wantedAttributes {
		attributeList = append(attributeList, attr)
	}
	return attributeList
}

func (b *DynamicPolicyBuilder) authTypesAsList() []int {
	authTypeList := make([]int, 0)
	for auth, b := range b.wantedAuthTypes {
		if b {
			authTypeList = append(authTypeList, auth)
		}
	}
	return authTypeList
}

// MarshalJSON returns the JSON encoding
func (policy *DynamicPolicy) MarshalJSON() ([]byte, error) {
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
