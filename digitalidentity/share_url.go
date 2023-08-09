package digitalidentity

var (
	// ShareURLHTTPErrorMessages specifies the HTTP error status codes used
	// by the Share URL API
	ShareHTTPErrorMessages = map[int]string{
		400: "JSON is incorrect, contains invalid data",
		404: "Application was not found",
	}
)

// ShareURL contains a dynamic share QR code
type ShareURL struct {
	ShareURL string `json:"qrcode"`
	RefID    string `json:"ref_id"`
}
