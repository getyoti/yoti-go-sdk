package service

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/retrieve"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/supported"
	"github.com/getyoti/yoti-go-sdk/v3/media"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

// CreateSession creates a Doc Scan session using the supplied session specification
func CreateSession(httpClient requests.HttpClient, sdkID string, key *rsa.PrivateKey, APIURL string, sessionSpec *create.SessionSpecification, jsonMarshaler jsonMarshaler) (*retrieve.CreateSessionResult, error) {
	requestBody, err := marshalJSON(jsonMarshaler, sessionSpec)
	if err != nil {
		return nil, err
	}

	var request *http.Request
	request, err = (&requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodPost,
		BaseURL:    APIURL,
		Endpoint:   createSessionPath(),
		Headers:    requests.JSONHeaders(),
		Body:       requestBody,
		Params:     map[string]string{"sdkID": sdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(httpClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *retrieve.CreateSessionResult
	err = json.Unmarshal(responseBytes, result)

	return result, err
}

// GetSession retrieves the state of a previously created Yoti Doc Scan session
func GetSession(httpClient requests.HttpClient, sdkID string, key *rsa.PrivateKey, APIURL string, sessionID string) (*retrieve.GetSessionResult, error) {
	request, err := (&requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    APIURL,
		Endpoint:   getSessionPath(sessionID),
		Params:     map[string]string{"sdkID": sdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(httpClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *retrieve.GetSessionResult
	err = json.Unmarshal(responseBytes, result)

	return result, err
}

// DeleteSession deletes a previously created Yoti Doc Scan session and all of its related resources
func DeleteSession(httpClient requests.HttpClient, sdkID string, key *rsa.PrivateKey, APIURL string, sessionID string) error {
	request, err := (&requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodDelete,
		BaseURL:    APIURL,
		Endpoint:   deleteSessionPath(sessionID),
		Params:     map[string]string{"sdkID": sdkID},
	}).Request()
	if err != nil {
		return err
	}

	_, err = requests.Execute(httpClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return err
	}

	return nil
}

// GetMediaContent retrieves media related to a Yoti Doc Scan session based on the supplied media ID
func GetMediaContent(httpClient requests.HttpClient, sdkID string, key *rsa.PrivateKey, APIURL, sessionID, mediaID string) (media.Media, error) {
	request, err := (&requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    APIURL,
		Endpoint:   getMediaContentPath(sessionID, mediaID),
		Params:     map[string]string{"sdkID": sdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(httpClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	contentType := response.Header.Get("Content-type") // TODO: check this
	if contentType == "" {
		return nil, errors.New("unable to parse content type from response")
	}

	return media.NewMedia(contentType, responseBytes), nil
}

// DeleteMediaContent deletes media related to a Yoti Doc Scan session based on the supplied media ID
func DeleteMediaContent(httpClient requests.HttpClient, sdkID string, key *rsa.PrivateKey, APIURL, sessionID, mediaID string) error {
	request, err := (&requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodDelete,
		BaseURL:    APIURL,
		Endpoint:   deleteMediaPath(sessionID, mediaID),
		Params:     map[string]string{"sdkID": sdkID},
	}).Request()
	if err != nil {
		return err
	}

	_, err = requests.Execute(httpClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return err
	}

	return nil
}

// GetSupportedDocuments gets a list of supported documents
func GetSupportedDocuments(httpClient requests.HttpClient, key *rsa.PrivateKey, APIURL string) (*supported.DocumentsResponse, error) {
	request, err := (&requests.SignedRequest{
		Key:        key,
		HTTPMethod: http.MethodGet,
		BaseURL:    APIURL,
		Endpoint:   getSupportedDocumentsPath(),
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(httpClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *supported.DocumentsResponse
	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// jsonMarshaler is a mockable JSON marshaler
type jsonMarshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

func marshalJSON(jsonMarshaler jsonMarshaler, v interface{}) ([]byte, error) {
	if jsonMarshaler != nil {
		return jsonMarshaler.Marshal(v)
	}
	return json.Marshal(v)
}
