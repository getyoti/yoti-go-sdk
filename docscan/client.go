package docscan

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/retrieve"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/supported"
	"github.com/getyoti/yoti-go-sdk/v3/media"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

// Client is responsible for setting up test data in the sandbox instance.
type Client struct {
	// SDK ID. This can be found in the Yoti Hub after you have created and activated an application.
	SdkID string
	// Private Key associated for your application, can be downloaded from the Yoti Hub.
	Key *rsa.PrivateKey
	// Mockable HTTP Client Interface
	HTTPClient requests.HttpClient
	// API URL to use. This is not required, and a default will be set if not provided.
	apiURL string
	// Mockable JSON marshaler
	jsonMarshaler jsonMarshaler
}

var mustNotBeEmptyString = "%s cannot be an empty string"

// NewClient constructs a Client object
func NewClient(sdkID string, key []byte) (*Client, error) {
	if sdkID == "" {
		return nil, fmt.Errorf(mustNotBeEmptyString, "SdkID")
	}

	decodedKey, err := cryptoutil.ParseRSAKey(key)
	if err != nil {
		return nil, err
	}

	return &Client{
		SdkID:      sdkID,
		Key:        decodedKey,
		HTTPClient: http.DefaultClient,
		apiURL:     getAPIURL(),
	}, err
}

// OverrideAPIURL overrides the default API URL for this Yoti Client
func (c *Client) OverrideAPIURL(apiURL string) {
	c.apiURL = apiURL
}

func getAPIURL() string {
	if value, exists := os.LookupEnv("YOTI_DOC_SCAN_API_URL"); exists && value != "" {
		return value
	} else {
		return "https://api.yoti.com/idverify/v1"
	}
}

// CreateSession creates a Doc Scan (IDV) session using the supplied session specification
func (c *Client) CreateSession(sessionSpec *create.SessionSpecification) (*create.SessionResult, error) {
	requestBody, err := marshalJSON(c.jsonMarshaler, sessionSpec)
	if err != nil {
		return nil, err
	}

	var request *http.Request
	request, err = (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodPost,
		BaseURL:    c.apiURL,
		Endpoint:   createSessionPath(),
		Headers:    requests.JSONHeaders(),
		Body:       requestBody,
		Params:     map[string]string{"sdkID": c.SdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result create.SessionResult
	err = json.Unmarshal(responseBytes, &result)

	return &result, err
}

// GetSession retrieves the state of a previously created Yoti Doc Scan (IDV) session
func (c *Client) GetSession(sessionID string) (*retrieve.GetSessionResult, error) {
	if sessionID == "" {
		return nil, fmt.Errorf(mustNotBeEmptyString, "sessionID")
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodGet,
		BaseURL:    c.apiURL,
		Endpoint:   getSessionPath(sessionID),
		Params:     map[string]string{"sdkID": c.SdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result retrieve.GetSessionResult
	err = json.Unmarshal(responseBytes, &result)

	return &result, err
}

// DeleteSession deletes a previously created Yoti Doc Scan (IDV) session and all of its related resources
func (c *Client) DeleteSession(sessionID string) error {
	if sessionID == "" {
		return fmt.Errorf(mustNotBeEmptyString, "sessionID")
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodDelete,
		BaseURL:    c.apiURL,
		Endpoint:   deleteSessionPath(sessionID),
		Params:     map[string]string{"sdkID": c.SdkID},
	}).Request()
	if err != nil {
		return err
	}

	_, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return err
	}

	return nil
}

// GetMediaContent retrieves media related to a Yoti Doc Scan (IDV) session based on the supplied media ID
func (c *Client) GetMediaContent(sessionID, mediaID string) (media.Media, error) {
	if sessionID == "" {
		return nil, fmt.Errorf(mustNotBeEmptyString, "sessionID")
	}

	if mediaID == "" {
		return nil, fmt.Errorf(mustNotBeEmptyString, "mediaID")
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodGet,
		BaseURL:    c.apiURL,
		Endpoint:   getMediaContentPath(sessionID, mediaID),
		Params:     map[string]string{"sdkID": c.SdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	contentTypes := strings.Split(response.Header.Get("Content-type"), ";")
	if len(contentTypes[0]) < 1 {
		err = errors.New("unable to parse content type from response")
	}

	result := media.NewMedia(contentTypes[0], responseBytes)

	return result, err
}

// DeleteMediaContent deletes media related to a Yoti Doc Scan (IDV) session based on the supplied media ID
func (c *Client) DeleteMediaContent(sessionID, mediaID string) error {
	if sessionID == "" {
		return fmt.Errorf(mustNotBeEmptyString, "sessionID")
	}

	if mediaID == "" {
		return fmt.Errorf(mustNotBeEmptyString, "mediaID")
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodDelete,
		BaseURL:    c.apiURL,
		Endpoint:   deleteMediaPath(sessionID, mediaID),
		Params:     map[string]string{"sdkID": c.SdkID},
	}).Request()
	if err != nil {
		return err
	}

	_, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return err
	}

	return nil
}

// GetSupportedDocuments gets a slice of supported documents
func (c *Client) GetSupportedDocuments() (*supported.DocumentsResponse, error) {
	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodGet,
		BaseURL:    c.apiURL,
		Endpoint:   getSupportedDocumentsPath(),
	}).Request()
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result supported.DocumentsResponse
	err = json.Unmarshal(responseBytes, &result)

	return &result, err
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
