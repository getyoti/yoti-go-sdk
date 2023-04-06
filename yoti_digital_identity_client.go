package yoti

import (
	"crypto/rsa"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
)

const DefaultURL = "https://api.yoti.com/api/"

// Client represents a client that can communicate with yoti and return information about Yoti users.
type DigitalIdentityClient struct {
	// SdkID represents the SDK ID and NOT the App ID. This can be found in the integration section of your
	// application hub at https://hub.yoti.com/
	SdkID string

	// Key should be the security key given to you by yoti (see: security keys section of
	// https://hub.yoti.com) for more information about how to load your key from a file see:
	// https://github.com/getyoti/yoti-go-sdk/blob/master/README.md
	Key *rsa.PrivateKey

	apiURL     string
	HTTPClient requests.HttpClient // Mockable HTTP Client Interface
}

// NewClient constructs a Client object
func DigitalIDClient(sdkID string, key []byte) (*Client, error) {
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
func (client *DigitalIdentityClient) OverrideAPIURL(apiURL string) {
	client.apiURL = apiURL
}

func (client *DigitalIdentityClient) getAPIURL() string {
	if client.apiURL != "" {
		return client.apiURL
	}

	if value, exists := os.LookupEnv("YOTI_API_URL"); exists && value != "" {
		return value
	}

	return apiDefaultURL
}

// GetSdkID gets the Client SDK ID attached to this client instance
func (client *DigitalIdentityClient) GetSdkID() string {
	return client.SdkID
}

// CreateShareURL creates a QR code for a specified dynamic scenario
func (client *DigitalIdentityClient) CreateShareURL(shareSession *digitalidentity.ShareSession) (share digitalidentity.ShareURL, err error) {
	return digitalidentity.CreateShareSession(client.HTTPClient, shareSession, client.GetSdkID(), client.getAPIURL(), client.Key)
}