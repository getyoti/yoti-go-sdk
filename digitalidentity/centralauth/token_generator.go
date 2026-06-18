// Package centralauth provides utilities for generating central auth tokens
// using the Yoti OAuth2 client_credentials grant.
package centralauth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity/requests"
)

const (
	// DefaultAuthURL is the default Yoti authentication API URL.
	DefaultAuthURL = "https://auth.api.yoti.com/v1/oauth/token"

	// EnvAuthURL is the environment variable name for overriding the auth API URL.
	EnvAuthURL = "YOTI_AUTH_URL"
)

// AuthenticationTokenResponse represents the response from the Yoti auth API
// when generating a new authentication token.
type AuthenticationTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// AuthenticationTokenGenerator generates authentication tokens by performing an
// OAuth2 client_credentials grant using a PS384-signed JWT as the client assertion.
// The generated token can then be used with BearerTokenStrategy for authorized requests.
type AuthenticationTokenGenerator struct {
	sdkID         string
	key           *rsa.PrivateKey
	jwtIDSupplier func() string
	authAPIURL    string
	httpClient    requests.HttpClient
}

// Generate creates a new authentication token with the specified scopes.
// Scopes must not be empty.
func (g *AuthenticationTokenGenerator) Generate(scopes []string) (*AuthenticationTokenResponse, error) {
	if len(scopes) == 0 {
		return nil, fmt.Errorf("scopes must not be empty")
	}

	jwt, err := g.createSignedJWT()
	if err != nil {
		return nil, fmt.Errorf("failed to create signed JWT: %v", err)
	}

	formData := url.Values{
		"grant_type":            {"client_credentials"},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"scope":                 {strings.Join(scopes, " ")},
		"client_assertion":      {jwt},
	}

	return g.performFormRequest(formData)
}

func (g *AuthenticationTokenGenerator) createSignedJWT() (string, error) {
	sdkIDProperty := "sdk:" + g.sdkID
	now := time.Now().Unix()
	jwtID := g.jwtIDSupplier()

	header := map[string]string{
		"alg": "PS384",
		"typ": "JWT",
	}

	claims := map[string]interface{}{
		"iss": sdkIDProperty,
		"sub": sdkIDProperty,
		"jti": jwtID,
		"aud": g.authAPIURL,
		"exp": now + 300,
		"iat": now,
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JWT header: %v", err)
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JWT claims: %v", err)
	}

	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)
	signingInput := headerEncoded + "." + claimsEncoded

	hash := sha512.Sum384([]byte(signingInput))
	signature, err := rsa.SignPSS(rand.Reader, g.key, crypto.SHA384, hash[:], &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthEqualsHash,
	})
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %v", err)
	}

	return signingInput + "." + base64.RawURLEncoding.EncodeToString(signature), nil
}

func (g *AuthenticationTokenGenerator) performFormRequest(formData url.Values) (*AuthenticationTokenResponse, error) {
	req, err := http.NewRequest(http.MethodPost, g.authAPIURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create auth token request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("auth token request failed: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read auth token response: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("auth token request failed with HTTP %d: %s", resp.StatusCode, string(responseBody))
	}

	var tokenResponse AuthenticationTokenResponse
	if err := json.Unmarshal(responseBody, &tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to deserialize authentication token response: %v", err)
	}

	return &tokenResponse, nil
}

// Builder provides a fluent API for constructing an AuthenticationTokenGenerator.
type Builder struct {
	sdkID         string
	key           *rsa.PrivateKey
	jwtIDSupplier func() string
	authAPIURL    string
	httpClient    requests.HttpClient
}

// NewBuilder creates a new Builder for AuthenticationTokenGenerator.
func NewBuilder() *Builder {
	return &Builder{}
}

// WithSdkID sets the SDK ID for the generator.
func (b *Builder) WithSdkID(sdkID string) *Builder {
	b.sdkID = sdkID
	return b
}

// WithKey sets the RSA private key for signing JWTs.
func (b *Builder) WithKey(key *rsa.PrivateKey) *Builder {
	b.key = key
	return b
}

// WithJwtIDSupplier sets a custom function for generating JWT IDs.
// By default, a UUID-like random string is used.
func (b *Builder) WithJwtIDSupplier(supplier func() string) *Builder {
	b.jwtIDSupplier = supplier
	return b
}

// WithAuthAPIURL sets the authentication API URL.
// By default, the YOTI_AUTH_URL environment variable is checked,
// then falls back to DefaultAuthURL.
func (b *Builder) WithAuthAPIURL(authAPIURL string) *Builder {
	b.authAPIURL = authAPIURL
	return b
}

// WithHTTPClient sets a custom HTTP client for making token requests.
func (b *Builder) WithHTTPClient(httpClient requests.HttpClient) *Builder {
	b.httpClient = httpClient
	return b
}

func defaultJwtIDSupplier() string {
	buf := make([]byte, 16)
	_, _ = rand.Read(buf)
	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}

// Build creates the AuthenticationTokenGenerator with the configured settings.
func (b *Builder) Build() (*AuthenticationTokenGenerator, error) {
	if b.sdkID == "" {
		return nil, fmt.Errorf("SDK ID must not be empty")
	}
	if b.key == nil {
		return nil, fmt.Errorf("private key must not be nil")
	}

	jwtIDSupplier := b.jwtIDSupplier
	if jwtIDSupplier == nil {
		jwtIDSupplier = defaultJwtIDSupplier
	}

	authAPIURL := b.authAPIURL
	if authAPIURL == "" {
		if envURL, exists := os.LookupEnv(EnvAuthURL); exists && envURL != "" {
			authAPIURL = envURL
		} else {
			authAPIURL = DefaultAuthURL
		}
	}

	httpClient := b.httpClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	return &AuthenticationTokenGenerator{
		sdkID:         b.sdkID,
		key:           b.key,
		jwtIDSupplier: jwtIDSupplier,
		authAPIURL:    authAPIURL,
		httpClient:    httpClient,
	}, nil
}
