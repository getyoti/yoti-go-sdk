package digitalidentity

var (
	// ShareURLHTTPErrorMessages specifies the HTTP error status codes used
	// by the Share URL API
	ShareSessionHTTPErrorMessages = map[int]string{
		400: "JSON is incorrect, contains invalid data",
		404: "Application was not found",
	}
)

// ShareSession object
type ShareSession struct {
	Id      string   `json:"id"`
	Status  string   `json:"status"`
	Expiry  string   `json:"expiry"`
	Created string   `json:"created"`
	Updated string   `json:"updated"`
	QrCode  qrCode   `json:"qrCode"`
	Receipt *receipt `json:"receipt"`
}

// ShareSession contains QR code as string
type qrCode struct {
	Id string `json:"id"`
}

// receipt containin id as string
type receipt struct {
	Id string `json:"id"`
}
