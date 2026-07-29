// Package requests provides authentication strategies and HTTP request building
// for the Yoti digital identity API.
package requests

import (
	"bytes"
	"errors"
	"net/http"
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
)

// AuthStrategy defines the interface for authentication mechanisms used when
// making requests to the Yoti API. Implementations can provide either
// signed-request authentication or bearer token (central auth) authentication.
type AuthStrategy interface {
	// CreateAuthHeaders returns the headers required for authentication.
	// The httpMethod, endpoint, and body are provided to support signing-based
	// strategies that need to compute a digest.
	CreateAuthHeaders(httpMethod, endpoint string, body []byte) (map[string][]string, error)

	// CreateQueryParams returns any query parameters required by the auth strategy.
	// For signed requests this includes nonce, timestamp, and sdkID.
	// For bearer token auth this returns an empty map.
	CreateQueryParams() (map[string]string, error)
}

// BuildAuthRequest constructs an http.Request using the provided AuthStrategy
// for authentication. It applies the strategy's headers and query params.
func BuildAuthRequest(strategy AuthStrategy, httpMethod, baseURL, endpoint string, headers map[string][]string, body []byte) (*http.Request, error) {
	if strategy == nil {
		return nil, errors.New("auth strategy must not be nil")
	}

	queryParams, err := strategy.CreateQueryParams()
	if err != nil {
		return nil, err
	}

	// Build endpoint with query params
	endpointWithParams := endpoint
	if len(queryParams) > 0 {
		separator := "?"
		if strings.Contains(endpoint, "?") {
			separator = "&"
		}
		for param, value := range queryParams {
			endpointWithParams += separator + param + "=" + value
			separator = "&"
		}
	}

	authHeaders, err := strategy.CreateAuthHeaders(httpMethod, endpointWithParams, body)
	if err != nil {
		return nil, err
	}

	// Construct the HTTP request
	var bodyReader *bytes.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	request, err := http.NewRequest(httpMethod, baseURL+endpointWithParams, bodyReader)
	if err != nil {
		return nil, err
	}

	// Apply auth headers
	for key, values := range authHeaders {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	// Apply custom headers
	for key, values := range headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	// Add SDK identification headers
	request.Header.Add("X-Yoti-SDK", consts.SDKIdentifier)
	request.Header.Add("X-Yoti-SDK-Version", consts.SDKVersionIdentifier)

	return request, nil
}
