package yoti

import (
	"fmt"
	"strconv"
	"time"
)

func getProfileEndpoint(token, nonce, sdkID string) string {
	timestamp := getTimestamp()

	return fmt.Sprintf("/profile/%s?nonce=%s&timestamp=%s&appId=%s", token, nonce, timestamp, sdkID)
}

func getAMLEndpoint(nonce, sdkID string) string {
	timestamp := getTimestamp()

	return fmt.Sprintf("/aml-check?appId=%s&timestamp=%s&nonce=%s", sdkID, timestamp, nonce)
}

func getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix()*1000, 10)
}

// GetDynamicShareEndpoint gets the Dynamic Share Endpoint URI
func getDynamicShareEndpoint(client clientInterface) (string, error) {
	timestamp := getTimestamp()
	nonce, err := generateNonce()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"/qrcodes/apps/%s?nonce=%s&timestamp=%s",
		client.GetSdkID(),
		nonce,
		timestamp,
	), nil
}
