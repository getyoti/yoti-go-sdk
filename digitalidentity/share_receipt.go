package digitalidentity

import "github.com/getyoti/yoti-go-sdk/v3/extra"

type SharedReceiptResponse struct {
	ID                 string
	SessionID          string
	RememberMeID       string
	ParentRememberMeID string
	Timestamp          string
	Error              string
	UserContent        UserContent
	ApplicationContent ApplicationContent
}

type ApplicationContent struct {
	ApplicationProfile ApplicationProfile
	ExtraData          *extra.Data
}

type UserContent struct {
	UserProfile UserProfile
	ExtraData   *extra.Data
}
