package shareurl

import (
	"encoding/json"

	dynamic "github.com/getyoti/yoti-go-sdk/v3/dynamic_sharing_service"
	"github.com/getyoti/yoti-go-sdk/v3/web"
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
func CreateShareURL(client web.ClientInterface, scenario *dynamic.DynamicScenario) (share ShareURL, err error) {
	httpMethod := "POST"
	endpoint, err := web.GetDynamicShareEndpoint(client)
	if err != nil {
		return
	}
	payload, err := scenario.MarshalJSON()
	if err != nil {
		return
	}

	response, err := client.MakeRequest(httpMethod, endpoint, payload, false, ShareURLHTTPErrorMessages, web.DefaultHTTPErrorMessages)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(response), &share)

	return
}
