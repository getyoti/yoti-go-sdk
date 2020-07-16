package dynamic_sharing_service

var (
	// ShareURLHTTPErrorMessages specifies the HTTP error status codes used
	// by the Share URL API
	ShareURLHTTPErrorMessages = map[int]string{
		400: "JSON is incorrect, contains invalid data: %[2]s",
		404: "Application was not found: %[2]s",
	}
)

// ShareURL contains a dynamic share QR code
type ShareURL struct {
	ShareURL string `json:"qrcode"`
	RefID    string `json:"ref_id"`
}
