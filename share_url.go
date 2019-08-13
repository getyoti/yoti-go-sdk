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

// ShareURL contains a dynamic share qr code
type ShareURL struct {
	ShareURL string `json:"qrcode"`
	RefID    string `json:"ref_id"`
}

// CreateShareURL creates a QR Code for a dynamic scenario
func CreateShareURL(client ClientInterface, scenario *DynamicScenario) (share ShareURL, err error) {
	httpMethod := "POST"
	endpoint, err := GetDynamicShareEndpoint(client)
	if err != nil {
		return
	}
	payload, err := scenario.MarshalJSON()
	if err != nil {
		return
	}

	response, err := client.MakeRequest(httpMethod, endpoint, payload, ShareURLHTTPErrorMessages, DefaultHTTPErrorMessages)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(response), &share)

	return
}
