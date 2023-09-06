package digitalidentity

import (
	"crypto/rsa"
	"encoding/json"
	"io"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

const identitySesssionCreationEndpoint = "v2/sessions"
const identitySessionRetrieval = "v2/sessions/%s"

// CreateShareSession creates session using the supplied session specification
func CreateShareSession(httpClient requests.HttpClient, shareSessionRequest *ShareSessionRequest, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share ShareSession, err error) {
	endpoint := identitySesssionCreationEndpoint

	payload, err := shareSessionRequest.MarshalJSON()
	if err != nil {
		return share, err
	}

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Body:       payload,
	}.Request()
	if err != nil {
		return share, err
	}

	response, err := requests.Execute(httpClient, request, ShareSessionHTTPErrorMessages, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return share, err
	}

	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return share, err
	}

	err = json.Unmarshal(responseBytes, &share)

	return share, err
}

// GetSession get session info using the supplied sessionID
func GetSession(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share *ShareSession, err error) {
	endpoint := identitySesssionCreationEndpoint
	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
	}.Request()
	if err != nil {
		return share, err
	}

	response, err := requests.Execute(httpClient, request, ShareSessionHTTPErrorMessages, yotierror.DefaultHTTPErrorMessages)

	if err != nil {
		return share, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return share, err
	}

	err = json.Unmarshal(responseBytes, &share)

	return share, err
}
