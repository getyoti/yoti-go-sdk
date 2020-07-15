package aml

import "fmt"

// GetAMLEndpoint gets the AML Endpoint URI
func GetAMLEndpoint(sdkID string) string {
	return fmt.Sprintf("/aml-check?appId=%s", sdkID)
}
