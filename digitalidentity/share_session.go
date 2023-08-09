package digitalidentity

// ShareSessionResult contains a dynamic share QR code
type ShareSessionResult struct {
	Id      string  `json:"id"`
	Status  string  `json:"status"`
	Expiry  string  `json:"expiry"`
	Created string  `json:"created"`
	Updated string  `json:"updated"`
	QrCode  qrCode  `json:"qrCode"`
	Receipt receipt `json:"receipt"`
}

// ShareSessionResult contains a dynamic share QR code
type qrCode struct {
	Id string `json:"id"`
}

// ShareSessionResult contains a dynamic share QR code
type receipt struct {
	Id string `json:"id"`
}
