package digitalidentity

type Content struct {
	Profile   []byte `json:"profile"`
	ExtraData []byte `json:"extraData"`
}

type ReceiptResponse struct {
	ID                 string   `json:"id"`
	SessionID          string   `json:"sessionId"`
	Timestamp          string   `json:"timestamp"`
	RememberMeID       string   `json:"rememberMeId,omitempty"`
	ParentRememberMeID string   `json:"parentRememberMeId,omitempty"`
	Content            *Content `json:"content,omitempty"`
	OtherPartyContent  *Content `json:"otherPartyContent,omitempty"`
	WrappedItemKeyId   string   `json:"wrappedItemKeyId"`
	WrappedKey         []byte   `json:"wrappedKey"`
	Error              string   `json:"error"`
}
