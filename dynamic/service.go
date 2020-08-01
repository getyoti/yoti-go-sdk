package dynamic

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
)

func getDynamicShareEndpoint(clientSdkId string) string {
	return fmt.Sprintf(
		"/qrcodes/apps/%s",
		clientSdkId,
	)
}

// CreateShareURL creates a QR code for a dynamic scenario
func CreateShareURL(httpClient requests.HttpClient, scenario *Scenario, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share ShareURL, err error) {
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

	response, err := requests.Execute(httpClient, request, ShareURLHTTPErrorMessages, requests.DefaultHTTPErrorMessages)
	if err != nil {
		return
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseBytes, &share)

	return
}
