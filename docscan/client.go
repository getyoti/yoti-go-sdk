package docscan

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/facecapture"
	"io"
	"net/http"
	"os"
	"strconv"
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
	responseBytes, err = io.ReadAll(response.Body)
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
	responseBytes, err = io.ReadAll(response.Body)
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
	responseBytes, err = io.ReadAll(response.Body)
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

// GetSupportedDocuments gets a slice of supported documents (defaults includeNonLatin to false)
func (c *Client) GetSupportedDocuments() (*supported.DocumentsResponse, error) {
	return c.GetSupportedDocumentsWithNonLatin(false)
}

// GetSupportedDocuments gets a slice of supported documents with bool param includeNonLatin
func (c *Client) GetSupportedDocumentsWithNonLatin(includeNonLatin bool) (*supported.DocumentsResponse, error) {

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodGet,
		BaseURL:    c.apiURL,
		Endpoint:   getSupportedDocumentsPath(),
		Params:     map[string]string{"includeNonLatin": strconv.FormatBool(includeNonLatin)},
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
	responseBytes, err = io.ReadAll(response.Body)
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

func (c *Client) CreateFaceCaptureResource(sessionID string, payload *facecapture.CreateFaceCaptureResourcePayload) (*retrieve.FaceCaptureResourceResponse, error) {
	if sessionID == "" {
		return nil, fmt.Errorf(mustNotBeEmptyString, "sessionID")
	}

	body, err := marshalJSON(c.jsonMarshaler, payload)
	if err != nil {
		return nil, err
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodPost,
		BaseURL:    c.apiURL,
		Endpoint:   fmt.Sprintf("/sessions/%s/resources/face-capture", sessionID),
		Params:     map[string]string{"sdkID": c.SdkID},
		Headers:    requests.JSONHeaders(),
		Body:       body,
	}).Request()
	if err != nil {
		return nil, err
	}

	resp, err := requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var result retrieve.FaceCaptureResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UploadFaceCaptureImage(sessionID, resourceID string, payload *facecapture.UploadFaceCaptureImagePayload) error {
	if sessionID == "" || resourceID == "" {
		return fmt.Errorf("sessionID and resourceID must not be empty")
	}

	if err := payload.Prepare(); err != nil {
		return fmt.Errorf("failed to prepare multipart payload: %w", err)
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodPut,
		BaseURL:    c.apiURL,
		Endpoint:   fmt.Sprintf("/sessions/%s/resources/face-capture/%s/image", sessionID, resourceID),
		Params:     map[string]string{"sdkID": c.SdkID},
		Body:       payload.MultipartFormBody().Bytes(),
		Headers:    payload.Headers(),
	}).Request()

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	return err
}

func (c *Client) GetSessionConfiguration(sessionID string) (*retrieve.SessionConfigurationResponse, error) {
	if sessionID == "" {
		return nil, fmt.Errorf(mustNotBeEmptyString, "sessionID")
	}

	request, err := (&requests.SignedRequest{
		Key:        c.Key,
		HTTPMethod: http.MethodGet,
		BaseURL:    c.apiURL,
		Endpoint:   fmt.Sprintf("/sessions/%s/configuration", sessionID),
		Params:     map[string]string{"sdkID": c.SdkID},
	}).Request()
	if err != nil {
		return nil, err
	}

	response, err := requests.Execute(c.HTTPClient, request, yotierror.DefaultHTTPErrorMessages)
	if err != nil {
		return nil, err
	}

	var responseBytes []byte
	responseBytes, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result retrieve.SessionConfigurationResponse

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

func (c *Client) AddFaceCaptureResourceToSession(sessionID string) error {
	sessionConfig, err := c.GetSessionConfiguration(sessionID)
	if err != nil {
		return err
	}

	if sessionConfig == nil {
		return fmt.Errorf("sessionConfig is nil")
	}

	capture := sessionConfig.GetCapture()
	if capture == nil {
		return fmt.Errorf("capture info is missing in sessionConfig")
	}

	requirements := capture.GetFaceCaptureResourceRequirements()
	if len(requirements) == 0 {
		// No face capture resource requirement, nothing to add
		return nil
	}

	firstRequirement := requirements[0]
	if firstRequirement == nil || firstRequirement.ID == "" {
		return fmt.Errorf("invalid face capture resource requirement")
	}

	payload := facecapture.NewCreateFaceCaptureResourcePayload(firstRequirement.ID)

	resource, err := c.CreateFaceCaptureResource(sessionID, payload)
	if err != nil {
		return err
	}

	base64Image := "iVBORw0KGgoAAAANSUhEUgAAAsAAAAGMAQMAAADuk4YmAAAAA1BMVEX///+nxBvIAAAAAXRSTlMAQObYZgAAADlJREFUeF7twDEBAAAAwiD7p7bGDlgYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwAGJrAABgPqdWQAAAABJRU5ErkJggg=="
	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return err
	}

	imagePayload := facecapture.NewUploadFaceCaptureImagePayload("image/png", imageBytes)
	return c.UploadFaceCaptureImage(sessionID, resource.ID, imagePayload)
}
