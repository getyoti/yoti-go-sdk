package yoti

import (
	"encoding/json"
)

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

// CreateShareURL creates a QR code for a dynamic scenario
func CreateShareURL(client clientInterface, scenario *DynamicScenario) (share ShareURL, err error) {
	httpMethod := "POST"
	endpoint, err := getDynamicShareEndpoint(client)
	if err != nil {
		return
	}
	payload, err := scenario.MarshalJSON()
	if err != nil {
		return
	}

	response, err := client.makeRequest(httpMethod, endpoint, payload, false, ShareURLHTTPErrorMessages, DefaultHTTPErrorMessages)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(response), &share)

	return
}
