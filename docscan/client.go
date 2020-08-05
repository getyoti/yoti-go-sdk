package docscan

import (
	"crypto/rsa"
	"net/http"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/service"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/retrieve"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/supported"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
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
}

// NewClient constructs a Client object
func NewClient(sdkID string, key []byte) (*Client, error) {
	decodedKey, err := cryptoutil.ParseRSAKey(key)

	if err != nil {
		return nil, err
	}

	return &Client{
		SdkID: sdkID,
		Key:   decodedKey,
	}, err
}

// OverrideAPIURL overrides the default API URL for this Yoti Client
func (client *Client) OverrideAPIURL(apiURL string) {
	client.apiURL = apiURL
}

func (client *Client) getAPIURL() string {
	if client.apiURL == "" {
		if value, exists := os.LookupEnv("YOTI_DOC_SCAN_API_URL"); exists && value != "" {
			client.apiURL = value
		} else {
			client.apiURL = "https://api.yoti.com/sandbox/idverify/v1"
		}
	}
	return client.apiURL
}

func (client *Client) getHTTPClient() requests.HttpClient {
	if client.HTTPClient != nil {
		return client.HTTPClient
	}
	return http.DefaultClient
}

// CreateSession creates a Doc Scan session using the supplied session specification
func (client *Client) CreateSession(sessionSpec *create.SessionSpecification) (*retrieve.CreateSessionResult, error) {
	return service.CreateSession(client.getHTTPClient(), client.SdkID, client.Key, client.getAPIURL(), sessionSpec, nil)
}

// GetSession retrieves the state of a previously created Yoti Doc Scan session
func (client *Client) GetSession(sessionID string) (*retrieve.GetSessionResult, error) {
	return service.GetSession(client.getHTTPClient(), client.SdkID, client.Key, client.getAPIURL(), sessionID)
}

// DeleteSession deletes a previously created Yoti Doc Scan session and all of its related resources
func (client *Client) DeleteSession(sessionID string) error {
	return service.DeleteSession(client.getHTTPClient(), client.SdkID, client.Key, client.getAPIURL(), sessionID)
}

// GetMediaContent retrieves media related to a Yoti Doc Scan session based on the supplied media ID
func (client *Client) GetMediaContent(sessionID, mediaID string) (*attribute.Image, error) { // TODO: change to media.Value
	return service.GetMediaContent(client.getHTTPClient(), client.SdkID, client.Key, client.getAPIURL(), sessionID, mediaID)
}

// DeleteMediaContent deletes media related to a Yoti Doc Scan session based on the supplied media ID
func (client *Client) DeleteMediaContent(sessionID, mediaID string) error {
	return service.DeleteMediaContent(client.getHTTPClient(), client.SdkID, client.Key, client.getAPIURL(), sessionID, mediaID)
}

// GetSupportedDocuments gets a list of supported documents
func (client *Client) GetSupportedDocuments() (*supported.DocumentsResponse, error) {
	return service.GetSupportedDocuments(client.getHTTPClient(), client.Key, client.getAPIURL())
}
