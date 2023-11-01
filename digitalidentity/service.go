package digitalidentity

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity/requests"
)

const identitySessionCreationEndpoint = "/v2/sessions"
const identitySessionRetrieval = "/v2/sessions/%s"
const identitySessionQrCodeCreation = "/v2/sessions/%s/qr-codes"
const identitySessionQrCodeRetrieval = "/v2/qr-codes/%s"

// CreateShareSession creates session using the supplied session specification
func CreateShareSession(httpClient requests.HttpClient, shareSessionRequest *ShareSessionRequest, clientSdkId, apiUrl string, key *rsa.PrivateKey) (*ShareSession, error) {
	endpoint := identitySessionCreationEndpoint

	payload, err := shareSessionRequest.MarshalJSON()
	if err != nil {
		return nil, err
	}

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Body:       payload,
		Params:     map[string]string{"sdkID": clientSdkId},
	}.Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	shareSession := &ShareSession{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBytes, shareSession)
	return shareSession, err
}

// GetShareSession get session info using the supplied sessionID parameter
func GetShareSession(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (*ShareSession, error) {
	endpoint := fmt.Sprintf(identitySessionRetrieval, sessionID)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Params:     map[string]string{"sdkID": clientSdkId},
	}.Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(httpClient, request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	shareSession := &ShareSession{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBytes, shareSession)
	return shareSession, err
}

// CreateShareQrCode generates a sharing qr code using the supplied sessionID parameter
func CreateShareQrCode(httpClient requests.HttpClient, sessionID string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (*QrCode, error) {
	endpoint := fmt.Sprintf(identitySessionQrCodeCreation, sessionID)

	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    requests.AuthHeader(clientSdkId),
		Body:       nil,
		Params:     map[string]string{"sdkID": clientSdkId},
	}.Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	qrCode := &QrCode{}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBytes, qrCode)
	return qrCode, err
}

// GetShareSessionQrCode is used to fetch the qr code by  id.
func GetShareSessionQrCode(httpClient requests.HttpClient, qrCodeId string, clientSdkId, apiUrl string, key *rsa.PrivateKey) (fetchedQrCode ShareSessionQrCode, err error) {
	endpoint := fmt.Sprintf(identitySessionQrCodeRetrieval, qrCodeId)
	headers := requests.AuthHeader(clientSdkId)
	request, err := requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    apiUrl,
		Endpoint:   endpoint,
		Headers:    headers,
	}.Request()
	if err != nil {
		return fetchedQrCode, err
	}

	response, err := requests.Execute(httpClient, request)
	if err != nil {
		return fetchedQrCode, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fetchedQrCode, err
	}

	err = json.Unmarshal(responseBytes, &fetchedQrCode)

	return fetchedQrCode, err
}
