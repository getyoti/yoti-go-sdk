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
