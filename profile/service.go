package profile

import "fmt"

// GetProfileEndpoint gets the Profile Endpoint URI
func GetProfileEndpoint(token, sdkID string) string {
	return fmt.Sprintf("/profile/%s?appId=%s", token, sdkID)
}
