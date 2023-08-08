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
const identitySessionQrCodeCreation = "/v2/sessions/%s/qr-codes"
const identitySessionQrCodeRetrieval = "/v2/qr-codes/%s"
const identitySessionReceiptRetrieval = "/v2/receipts/%s"
const identitySessionReceiptKeyRetrieval = "/v2/wrapped-item-keys/%s"

// SessionResult contains the information about a created session
type SessionResult struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Expiry string `json:"expiry"`
}

// CreateShareSession creates session using the supplied session specification
func CreateShareSession(httpClient requests.HttpClient, shareSession *ShareSession, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share SessionResult, err error) {
	endpoint := identitySesssionCreationEndpoint

	payload, err := shareSession.MarshalJSON()
	if err != nil {
		return share, err
	}

	headers := requests.AuthHeader(clientSdkId)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
		Body:       payload,
	}.Request()
	if err != nil {
		return share, err
	}

	response, err := requests.Execute(httpClient, request, ShareURLHTTPErrorMessages, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		//fmt.Printf("err 2:=> %s\n\r", err)
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
func GetSession(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (share ShareSessionResult, err error) {
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

	response, err := requests.Execute(httpClient, request, ShareURLHTTPErrorMessages, yotierror.DefaultHTTPErrorMessages)
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
func CreateShareQrCode(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (qrCode CreateShareQrCodeResult, err error) {
	endpoint := identitySessionQrCodeCreation
	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
		Body:       nil,
	}.Request()
	if err != nil {
		return qrCode, err
	}

	response, err := requests.Execute(httpClient, request, ShareURLHTTPErrorMessages, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return qrCode, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return qrCode, err
	}

	err = json.Unmarshal(responseBytes, &qrCode)

	return qrCode, err
}
