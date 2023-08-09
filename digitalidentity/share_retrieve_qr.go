package digitalidentity

type ShareSessionFetchedQrCode struct {
	ID          string              `json:"id"`
	Expiry      string              `json:"expiry"`
	Policy      string              `json:"policy"`
	Extensions  []interface{}       `json:"extensions"`
	Session     ShareSessionCreated `json:"session"`
	RedirectURI string              `json:"redirectUri"`
}
