package profile

import (
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/share"
)

// ActivityDetails represents the result of an activity between a user and the application.
type ActivityDetails struct {
	UserProfile        Profile
	rememberMeID       string
	parentRememberMeID string
	timestamp          time.Time
	receiptID          string
	ApplicationProfile ApplicationProfile
	extraData          *share.ExtraData
}

// RememberMeID is a unique, stable identifier for a user in the context
// of an application. You can use it to identify returning users.
// This value will be different for the same user in different applications.
func (a ActivityDetails) RememberMeID() string {
	return a.rememberMeID
}

// ParentRememberMeID is a unique, stable identifier for a user in the
// context of an organisation. You can use it to identify returning users.
// This value is consistent for a given user across different applications
// belonging to a single organisation.
func (a ActivityDetails) ParentRememberMeID() string {
	return a.parentRememberMeID
}

// Timestamp is the Time and date of the sharing activity
func (a ActivityDetails) Timestamp() time.Time {
	return a.timestamp
}

// ReceiptID identifies a completed activity
func (a ActivityDetails) ReceiptID() string {
	return a.receiptID
}

// ExtraData represents extra pieces information on the receipt
func (a ActivityDetails) ExtraData() *share.ExtraData {
	return a.extraData
}
