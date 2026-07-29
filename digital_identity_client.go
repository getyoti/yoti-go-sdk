package yoti

import (
	"crypto/rsa"
	"errors"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
	direquests "github.com/getyoti/yoti-go-sdk/v3/digitalidentity/requests"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
)

const DefaultURL = "https://api.yoti.com/share"

// DigitalIdentityClient represents a client that can communicate with yoti and return information about Yoti users.
type DigitalIdentityClient struct {
	// SdkID represents the SDK ID and NOT the App ID. This can be found in the integration section of your
	// application hub at https://hub.yoti.com/
	SdkID string

	// Key should be the security key given to you by yoti (see: security keys section of
	// https://hub.yoti.com) for more information about how to load your key from a file see:
	// https://github.com/getyoti/yoti-go-sdk/blob/master/README.md
	Key *rsa.PrivateKey

	// AuthToken is the central auth bearer token. When set, the client uses
	// bearer token authentication instead of signed-request authentication.
	// This is mutually exclusive with SdkID/Key.
	AuthToken string `json:"-"`

	apiURL     string
	HTTPClient requests.HttpClient // Mockable HTTP Client Interface
}

// NewDigitalIdentityClient constructs a Client object using signed-request authentication.
func NewDigitalIdentityClient(sdkID string, key []byte) (*DigitalIdentityClient, error) {
	decodedKey, err := cryptoutil.ParseRSAKey(key)

	if err != nil {
		return nil, err
	}

	return &DigitalIdentityClient{
		SdkID: sdkID,
		Key:   decodedKey,
	}, err
}

// NewDigitalIdentityClientWithToken constructs a Client object using central auth
// bearer token authentication. The token is provided by the relying business and
// will be sent as an Authorization: Bearer header on all API requests.
func NewDigitalIdentityClientWithToken(authToken string) (*DigitalIdentityClient, error) {
	if authToken == "" {
		return nil, errors.New("authentication token must not be empty")
	}
	return &DigitalIdentityClient{
		AuthToken: authToken,
	}, nil
}

// GetAuthStrategy returns the appropriate AuthStrategy based on the client configuration.
func (client *DigitalIdentityClient) GetAuthStrategy() (direquests.AuthStrategy, error) {
	if client.AuthToken != "" {
		if client.Key != nil || client.SdkID != "" {
			return nil, errors.New("must not supply both authentication token and SDK credentials (SdkID/Key)")
		}
		return direquests.NewBearerTokenStrategy(client.AuthToken)
	}

	if client.Key == nil {
		return nil, errors.New("missing private key for signed-request authentication")
	}
	if client.SdkID == "" {
		return nil, errors.New("missing SDK ID for signed-request authentication")
	}
	return direquests.NewSignedRequestStrategy(client.Key, client.SdkID)
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

	return DefaultURL
}

// GetSdkID gets the Client SDK ID attached to this client instance
func (client *DigitalIdentityClient) GetSdkID() string {
	return client.SdkID
}

// CreateShareSession creates a sharing session to initiate a sharing process based on a policy
func (client *DigitalIdentityClient) CreateShareSession(shareSessionRequest *digitalidentity.ShareSessionRequest) (shareSession *digitalidentity.ShareSession, err error) {
	strategy, err := client.GetAuthStrategy()
	if err != nil {
		return nil, err
	}
	return digitalidentity.CreateShareSession(client.HTTPClient, shareSessionRequest, client.getAPIURL(), strategy)
}

// GetShareSession retrieves the sharing session.
func (client *DigitalIdentityClient) GetShareSession(sessionID string) (*digitalidentity.ShareSession, error) {
	strategy, err := client.GetAuthStrategy()
	if err != nil {
		return nil, err
	}
	return digitalidentity.GetShareSession(client.HTTPClient, sessionID, client.getAPIURL(), strategy)
}

// CreateShareQrCode generates a sharing session QR code to initiate a sharing process based on session ID
func (client *DigitalIdentityClient) CreateShareQrCode(sessionID string) (share *digitalidentity.QrCode, err error) {
	strategy, err := client.GetAuthStrategy()
	if err != nil {
		return nil, err
	}
	return digitalidentity.CreateShareQrCode(client.HTTPClient, sessionID, client.getAPIURL(), strategy)
}

// Get session QR code based on generated Qr ID
func (client *DigitalIdentityClient) GetQrCode(qrCodeId string) (share digitalidentity.ShareSessionQrCode, err error) {
	strategy, err := client.GetAuthStrategy()
	if err != nil {
		return share, err
	}
	return digitalidentity.GetShareSessionQrCode(client.HTTPClient, qrCodeId, client.getAPIURL(), strategy)
}

// GetShareReceipt fetches the receipt of the share given a receipt id.
// Requires a private key for receipt decryption; not available with bearer token authentication.
func (client *DigitalIdentityClient) GetShareReceipt(receiptId string) (share digitalidentity.SharedReceiptResponse, err error) {
	if client.Key == nil {
		return share, errors.New("GetShareReceipt requires a private key for receipt decryption and cannot be used with bearer token authentication")
	}
	strategy, err := client.GetAuthStrategy()
	if err != nil {
		return share, err
	}
	return digitalidentity.GetShareReceipt(client.HTTPClient, receiptId, client.getAPIURL(), strategy, client.Key)
}
