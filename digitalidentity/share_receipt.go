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
	applicationProfile ApplicationProfile
	extraData          *extra.Data
}

type UserContent struct {
	userProfile UserProfile
	extraData   *extra.Data
}
