package digitalidentity

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

func getIdentitySesssionCreationEndpoint() string {
	return "/v2/sessions"
}

// SessionResult contains the information about a created session
type SessionResult struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Expiry string `json:"expiry"`
}

// CreateShareSession
func CreateShareSession(httpClient requests.HttpClient, shareSession *ShareSession, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share ShareURL, err error) {
	endpoint := getIdentitySesssionCreationEndpoint()

	payload, err := shareSession.MarshalJSON()
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

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseBytes, &share)

	return
}
