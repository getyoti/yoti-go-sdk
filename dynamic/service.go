package dynamic

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
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

	response, err := requests.Execute(httpClient, request, ShareURLHTTPErrorMessages, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return share, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseBytes, &share)

	return
}
