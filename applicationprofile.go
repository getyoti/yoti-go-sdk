package yoti

import (
	"github.com/getyoti/yoti-go-sdk/v2/attribute"
)

// Attribute names for application attributes
const (
	AttrConstApplicationName           = "application_name"
	AttrConstApplicationURL            = "application_url"
	AttrConstApplicationLogo           = "application_logo"
	AttrConstApplicationReceiptBGColor = "application_receipt_bgcolor"
)

// ApplicationProfile is the profile of an application with convenience methods
// to access well-known attributes.
type ApplicationProfile struct {
	baseProfile
}

// ApplicationName is the name of the application
func (p ApplicationProfile) ApplicationName() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstApplicationName)
}

// ApplicationURL is the URL where the application is available at
func (p ApplicationProfile) ApplicationURL() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstApplicationURL)
}

// ApplicationReceiptBgColor is the background colour that will be displayed on
// each receipt the user gets as a result of a sharing with the application.
func (p ApplicationProfile) ApplicationReceiptBgColor() *attribute.StringAttribute {
	return p.GetStringAttribute(AttrConstApplicationReceiptBGColor)
}

// ApplicationLogo is the logo of the application that will be displayed to
// those users that perform a sharing with it.
func (p ApplicationProfile) ApplicationLogo() *attribute.ImageAttribute {
	return p.GetImageAttribute(AttrConstApplicationLogo)
}
