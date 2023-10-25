package attribute

import (
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// StringAttribute is a Yoti attribute which returns a string as its value
type StringAttribute struct {
	attributeDetails
	value string
}

// NewString creates a new String attribute
func NewString(a *yotiprotoattr.Attribute) *StringAttribute {
	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &StringAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
			id:          &a.EphemeralId,
		},
		value: string(a.Value),
	}
}

// Value returns the value of the StringAttribute as a string
func (a *StringAttribute) Value() string {
	return a.value
}
