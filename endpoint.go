package yoti

import (
	"fmt"
)

func getProfileEndpoint(token, sdkID string) string {
	return fmt.Sprintf("/profile/%s?appId=%s", token, sdkID)
}

func getAMLEndpoint(sdkID string) string {
	return fmt.Sprintf("/aml-check?appId=%s", sdkID)
}

// GetDynamicShareEndpoint gets the Dynamic Share Endpoint URI
func getDynamicShareEndpoint(client clientInterface) (string, error) {
	return fmt.Sprintf(
		"/qrcodes/apps/%s",
		client.GetSdkID(),
	), nil
}
