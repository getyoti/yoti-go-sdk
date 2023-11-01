package digitalidentity

type ShareSessionQrCode struct {
	ID         string        `json:"id"`
	Expiry     string        `json:"expiry"`
	Policy     string        `json:"policy"`
	Extensions []interface{} `json:"extensions"`
	Session    struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Expiry string `json:"expiry"`
	} `json:"session"`
	RedirectURI string `json:"redirectUri"`
}
