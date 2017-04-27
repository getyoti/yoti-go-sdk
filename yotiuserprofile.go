package yoti

import "time"

type YotiUserProfile struct {
	// ID is a unique identifier Yoti assigns to your user, but only for your app
	// if the same user logs into your app again, you get the same id
	// if she/he logs into another application, Yoti will assign a different id for that app
	ID string

	// Selfie is a photograph of the user. This will be nil if not provided by Yoti
	Selfie *Image

	// GivenNames represents the user's given names. This will be an empty string if not provided by Yoti
	GivenNames string

	// Family represents the user's family name. This will be an empty string if not provided by Yoti
	FamilyName string

	// MobileNumber represents the user's mobile phone number. This will be an empty string if not provided by Yoti
	MobileNumber string

	// EmailAddress represents the user's email address. This will be an empty string if not provided by Yoti
	EmailAddress string

	// DateOfBirth represents the user's date of birth. This will be nil if not provided by Yoti
	DateOfBirth *time.Time

	// Address represents the user's address. This will be an empty string if not provided by Yoti
	Address string

	// Gender represents the user's gender. This will be an empty string if not provided by Yoti
	Gender string

	// Nationality represents the user's nationality. This will be an empty string if not provided by Yoti
	Nationality string

	// OtherAttributes is a map of any other information about the user provided by Yoti. The key will be the name
	// of the piece of information, and the keys associated value will be the piece of information itself.
	OtherAttributes map[string]YotiAttributeValue
}
