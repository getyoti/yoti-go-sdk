package dynamic_sharing_service

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/web"
)

func getDynamicShareEndpoint(clientSdkId string) string {
	return fmt.Sprintf(
		"/qrcodes/apps/%s",
		clientSdkId,
	)
}

// CreateShareURL creates a QR code for a dynamic scenario
func CreateShareURL(httpClient web.HttpClient, scenario *DynamicScenario, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share ShareURL, err error) {
	endpoint := getDynamicShareEndpoint(clientSdkId)

	payload, err := scenario.MarshalJSON()
	if err != nil {
		return
	}

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    nil,
		Body:       payload,
	}.Request()
	if err != nil {
		return
	}

	response, err := web.MakeRequest(httpClient, request, ShareURLHTTPErrorMessages, web.DefaultHTTPErrorMessages)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(response), &share)

	return
}
