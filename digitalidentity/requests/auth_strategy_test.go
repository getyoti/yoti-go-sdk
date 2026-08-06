package requests

import (
	"net/http"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/test"
	"gotest.tools/v3/assert"
)

func TestBearerTokenStrategy_RejectsEmptyToken(t *testing.T) {
	_, err := NewBearerTokenStrategy("")
	assert.ErrorContains(t, err, "authentication token must not be empty")
}

func TestBearerTokenStrategy_CreateAuthHeaders(t *testing.T) {
	strategy, err := NewBearerTokenStrategy("my-test-token")
	assert.NilError(t, err)

	headers, err := strategy.CreateAuthHeaders(http.MethodGet, "/v2/sessions", nil)
	assert.NilError(t, err)

	assert.DeepEqual(t, headers, map[string][]string{
		"Authorization": {"Bearer my-test-token"},
	})
}

func TestBearerTokenStrategy_CreateQueryParams_Empty(t *testing.T) {
	strategy, err := NewBearerTokenStrategy("my-test-token")
	assert.NilError(t, err)

	params, err := strategy.CreateQueryParams()
	assert.NilError(t, err)
	assert.Equal(t, len(params), 0)
}

func TestSignedRequestStrategy_RejectsNilKey(t *testing.T) {
	_, err := NewSignedRequestStrategy(nil, "sdkId")
	assert.ErrorContains(t, err, "private key must not be nil")
}

func TestSignedRequestStrategy_RejectsEmptySdkId(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	_, err := NewSignedRequestStrategy(key, "")
	assert.ErrorContains(t, err, "client SDK ID must not be empty")
}

func TestSignedRequestStrategy_CreateAuthHeaders(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	strategy, err := NewSignedRequestStrategy(key, "test-sdk-id")
	assert.NilError(t, err)

	headers, err := strategy.CreateAuthHeaders(http.MethodGet, "/v2/sessions?nonce=abc&timestamp=123", nil)
	assert.NilError(t, err)

	// Should have X-Yoti-Auth-Digest and X-Yoti-Auth-Id
	assert.Equal(t, len(headers["X-Yoti-Auth-Digest"]), 1)
	assert.Equal(t, headers["X-Yoti-Auth-Id"][0], "test-sdk-id")

	// Digest should be non-empty base64
	assert.Check(t, len(headers["X-Yoti-Auth-Digest"][0]) > 0)
}

func TestSignedRequestStrategy_CreateQueryParams(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	strategy, err := NewSignedRequestStrategy(key, "test-sdk-id")
	assert.NilError(t, err)

	params, err := strategy.CreateQueryParams()
	assert.NilError(t, err)

	// Should contain nonce, timestamp, sdkID
	assert.Check(t, params["nonce"] != "")
	assert.Check(t, params["timestamp"] != "")
	assert.Equal(t, params["sdkID"], "test-sdk-id")
}

func TestBuildAuthRequest_BearerToken(t *testing.T) {
	strategy, err := NewBearerTokenStrategy("test-bearer-token")
	assert.NilError(t, err)

	request, err := BuildAuthRequest(strategy, http.MethodGet, "https://api.yoti.com", "/v2/sessions", JSONHeaders(), nil)
	assert.NilError(t, err)

	// Should have Authorization header with bearer token
	assert.Equal(t, request.Header.Get("Authorization"), "Bearer test-bearer-token")

	// Should NOT have X-Yoti-Auth-Digest
	assert.Equal(t, request.Header.Get("X-Yoti-Auth-Digest"), "")

	// Should have SDK headers
	assert.Equal(t, request.Header.Get("X-Yoti-SDK"), "Go")
	assert.Check(t, request.Header.Get("X-Yoti-SDK-Version") != "")

	// Should have JSON headers
	assert.Equal(t, request.Header.Get("Content-Type"), "application/json")

	// URL should have no query params (bearer doesn't add them)
	assert.Equal(t, request.URL.RawQuery, "")
	assert.Equal(t, request.URL.Path, "/v2/sessions")
}

func TestBuildAuthRequest_SignedRequest(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	strategy, err := NewSignedRequestStrategy(key, "test-sdk-id")
	assert.NilError(t, err)

	request, err := BuildAuthRequest(strategy, http.MethodPost, "https://api.yoti.com", "/v2/sessions", JSONHeaders(), []byte(`{"policy":{}}`))
	assert.NilError(t, err)

	// Should have X-Yoti-Auth-Digest
	assert.Check(t, request.Header.Get("X-Yoti-Auth-Digest") != "")

	// Should have X-Yoti-Auth-Id
	assert.Equal(t, request.Header.Get("X-Yoti-Auth-Id"), "test-sdk-id")

	// Should NOT have Authorization bearer header
	assert.Equal(t, request.Header.Get("Authorization"), "")

	// URL should have query params (nonce, timestamp, sdkID)
	query := request.URL.Query()
	assert.Check(t, query.Get("nonce") != "")
	assert.Check(t, query.Get("timestamp") != "")
	assert.Equal(t, query.Get("sdkID"), "test-sdk-id")
}
