package web

import (
	"fmt"
)

// GetProfileEndpoint gets the Profile Endpoint URI
func GetProfileEndpoint(token, sdkID string) string {
	return fmt.Sprintf("/profile/%s?appId=%s", token, sdkID)
}

// GetAMLEndpoint gets the AML Endpoint URI
func GetAMLEndpoint(sdkID string) string {
	return fmt.Sprintf("/aml-check?appId=%s", sdkID)
}

// GetDynamicShareEndpoint gets the Dynamic Share Endpoint URI
func GetDynamicShareEndpoint(client ClientInterface) (string, error) {
	return fmt.Sprintf(
		"/qrcodes/apps/%s",
		client.GetSdkID(),
	), nil
}
