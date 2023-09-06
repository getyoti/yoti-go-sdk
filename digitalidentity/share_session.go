package digitalidentity

var (
	// ShareSessionHTTPErrorMessages specifies the HTTP error status codes used
	// by the Share Session API
	ShareSessionHTTPErrorMessages = map[int]string{
		400: "JSON is incorrect, contains invalid data",
		404: "Application was not found",
	}
)

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

// ShareSession contains QR code id as string
type qrCode struct {
	Id string `json:"id"`
}

// receipt containing the receipt id as a string.
type receipt struct {
	Id string `json:"id"`
}