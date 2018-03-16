package yoti

type receiptDO struct {
	ReceiptID                  string `json:"receipt_id"`
	OtherPartyProfileContent   string `json:"other_party_profile_content"`
	ProfileContent             string `json:"profile_content"`
	OtherPartyExtraDataContent string `json:"other_party_extra_data_content"`
	ExtraDataContent           string `json:"extra_data_content"`
	WrappedReceiptKey          string `json:"wrapped_receipt_key"`
	PolicyURI                  string `json:"policy_uri"`
	PersonalKey                string `json:"personal_key"`
	RememberMeID               string `json:"remember_me_id"`
	SharingOutcome             string `json:"sharing_outcome"`
	Timestamp                  string `json:"timestamp"`
}

type profileDO struct {
	SessionData string    `json:"session_data"`
	Receipt     receiptDO `json:"receipt"`
}
