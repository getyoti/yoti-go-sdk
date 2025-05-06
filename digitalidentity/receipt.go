package digitalidentity

type Content struct {
	Profile   []byte `json:"profile"`
	ExtraData []byte `json:"extraData"`
}

type ErrorReason struct {
	RequirementsNotMetDetails []struct {
		Details                string `json:"details"`
		AuditId                string `json:"audit_id"`
		FailureType            string `json:"failure_type"`
		DocumentType           string `json:"document_type"`
		DocumentCountryIsoCode string `json:"document_country_iso_code"`
	} `json:"requirements_not_met_details"`
}

type ReceiptResponse struct {
	ID                 string      `json:"id"`
	SessionID          string      `json:"sessionId"`
	Timestamp          string      `json:"timestamp"`
	RememberMeID       string      `json:"rememberMeId,omitempty"`
	ParentRememberMeID string      `json:"parentRememberMeId,omitempty"`
	Content            *Content    `json:"content,omitempty"`
	OtherPartyContent  *Content    `json:"otherPartyContent,omitempty"`
	WrappedItemKeyId   string      `json:"wrappedItemKeyId"`
	WrappedKey         []byte      `json:"wrappedKey"`
	Error              string      `json:"error"`
	ErrorReason        ErrorReason `json:"errorReason"`
}
