package digitalidentity

type Content struct {
	Profile   string `json:"profile"`
	ExtraData string `json:"extraData"`
}

type OtherPartyContent struct {
	Profile   string `json:"profile"`
	ExtraData string `json:"extraData"`
}

type ReceiptResponse struct {
	ID                 string             `json:"id"`
	SessionID          string             `json:"sessionId"`
	Timestamp          string             `json:"timestamp"`
	RememberMeID       string             `json:"rememberMeId,omitempty"`
	ParentRememberMeID string             `json:"parentRememberMeId,omitempty"`
	Content            *Content           `json:"content,omitempty"`
	OtherPartyContent  *OtherPartyContent `json:"otherPartyContent,omitempty"`
	WrappedItemKeyId   string             `json:"wrappedItemKeyId"`
	WrappedKey         string             `json:"wrappedKey"`
	Error              string             `json:"error"`
}
