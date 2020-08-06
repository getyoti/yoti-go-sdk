package sandbox

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/response/docscanerr"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
	yotirequest "github.com/getyoti/yoti-go-sdk/v3/requests"
)

type jsonMarshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

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

func (client *Client) marshalJSON(v interface{}) ([]byte, error) {
	if client.jsonMarshaler != nil {
		return client.jsonMarshaler.Marshal(v)
	}
	return json.Marshal(v)
}

func (client *Client) makeConfigureResponseRequest(request *http.Request) error {
	_, err := requests.Execute(client.getHTTPClient(), request)

	if err != nil {
		return err
	}

	return nil
}

// ConfigureSessionResponse configures the response for the session
func (client *Client) ConfigureSessionResponse(sessionID string, responseConfig *request.ResponseConfig) error {
	requestEndpoint := "/sessions/" + sessionID + "/response-config"
	requestBody, err := client.marshalJSON(responseConfig)
	if err != nil {
		return err
	}

	request, err := (&yotirequest.SignedRequest{
		Key:        client.Key,
		HTTPMethod: http.MethodPut,
		BaseURL:    client.getAPIURL(),
		Endpoint:   requestEndpoint,
		Headers:    yotirequest.JSONHeaders(),
		Body:       requestBody,
		Params:     map[string]string{"sdkId": client.SdkID},
	}).Request()
	if err != nil {
		return err
	}

	return client.makeConfigureResponseRequest(request)
}

// ConfigureApplicationResponse configures the response for the application
func (client *Client) ConfigureApplicationResponse(responseConfig *request.ResponseConfig) error {
	requestEndpoint := "/apps/" + client.SdkID + "/response-config"
	requestBody, err := client.marshalJSON(responseConfig)
	if err != nil {
		return err
	}

	request, err := (&yotirequest.SignedRequest{
		Key:        client.Key,
		HTTPMethod: http.MethodPut,
		BaseURL:    client.getAPIURL(),
		Endpoint:   requestEndpoint,
		Headers:    yotirequest.JSONHeaders(),
		Body:       requestBody,
	}).Request()
	if err != nil {
		return err
	}

	return client.makeConfigureResponseRequest(request)
}
