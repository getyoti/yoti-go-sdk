package sandbox

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	httpClient yoti.HttpClient
}

// SetupSharingProfile creates a user profile in the sandbox instance
func (client *Client) SetupSharingProfile(tokenRequest TokenRequest) (token string, err error) {
	if client.BaseURL == "" {
		if os.Getenv("YOTI_API_URL") != "" {
			client.BaseURL = os.Getenv("YOTI_API_URL")
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

	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}

	response, err := (client.httpClient).Do(request)
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