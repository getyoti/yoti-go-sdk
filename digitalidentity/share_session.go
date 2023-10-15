package digitalidentity

// ShareSession contains information about the session.
type ShareSession struct {
	Id      string   `json:"id"`
	Status  string   `json:"status"`
	Expiry  string   `json:"expiry"`
	Created string   `json:"created"`
	Updated string   `json:"updated"`
	QrCode  qrCode   `json:"qrCode"`
	Receipt *receipt `json:"receipt"`
}

type qrCode struct {
	Id string `json:"id"`
}

// receipt containing the receipt id as a string.
type receipt struct {
	Id string `json:"id"`
}
