package sandbox

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"

	yoti "github.com/getyoti/yoti-go-sdk/v2"
	yotirequest "github.com/getyoti/yoti-go-sdk/v2/requests"
)

// Client is responsible for setting up test data in the sandbox instance. BaseURL is not required.
type Client struct {
	// Client SDK ID. This can be found in the Yoti Hub after you have created and activated an application.
	ClientSdkID string
	// Private Key associated for your application, can be downloaded from the Yoti Hub.
	Key *rsa.PrivateKey
	// Base URL to use. This is not required, and a default will be set if not provided.
	BaseURL string
	// Mockable HTTP Client Interface
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}
}

func (client *Client) do(request *http.Request) (*http.Response, error) {
	if client.HTTPClient != nil {
		return client.HTTPClient.Do(request)
	}
	return http.DefaultClient.Do(request)
}

// SetupSharingProfile creates a user profile in the sandbox instance
func (client *Client) SetupSharingProfile(tokenRequest TokenRequest) (token string, err error) {
	if client.BaseURL == "" {
		if value, exists := os.LookupEnv("YOTI_API_URL"); exists && value != "" {
			client.BaseURL = value
		} else {
			client.BaseURL = "https://api.yoti.com/sandbox/v1"
		}
	}

	requestEndpoint := "/apps/" + client.ClientSdkID + "/tokens"
	requestBody, err := json.Marshal(tokenRequest)
	if err != nil {
		return
	}

	request, err := (&yotirequest.SignedRequest{
		Key:        client.Key,
		HTTPMethod: "POST",
		BaseURL:    client.BaseURL,
		Endpoint:   requestEndpoint,
		Headers:    yotirequest.JSONHeaders(),
		Body:       requestBody,
	}).Request()
	if err != nil {
		return
	}

	response, err := client.do(request)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(response.Body)
		return "", fmt.Errorf("Sharing Profile not created (HTTP %d) %s", response.StatusCode, string(body))
	}

	responseStruct := struct {
		Token string `json:"token"`
	}{}

	err = json.NewDecoder(response.Body).Decode(&responseStruct)
	token = responseStruct.Token

	return
}
