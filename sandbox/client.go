package sandbox

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	yotirequest "github.com/getyoti/yoti-go-sdk/v2/requests"
)

// Client is responsible for setting up test data in the sandbox instance
type Client struct {
	AppID      string
	Key        *rsa.PrivateKey
	BaseURL    string
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
func (client *Client) SetupSharingProfile(profile Profile) (token string, err error) {
	requestEndpoint := "/apps/" + client.AppID + "/tokens"
	requestBody, err := json.Marshal(profile)
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
