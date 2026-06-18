package centralauth

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/test"
	"gotest.tools/v3/assert"
)

type mockHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.do(req)
}

func TestBuilder_MissingSdkId(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	_, err := NewBuilder().WithKey(key).Build()
	assert.ErrorContains(t, err, "SDK ID must not be empty")
}

func TestBuilder_MissingKey(t *testing.T) {
	_, err := NewBuilder().WithSdkID("test-sdk-id").Build()
	assert.ErrorContains(t, err, "private key must not be nil")
}

func TestBuilder_Success(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	gen, err := NewBuilder().WithSdkID("test-sdk-id").WithKey(key).Build()
	assert.NilError(t, err)
	assert.Check(t, gen != nil)
}

func TestBuilder_CustomAuthURL(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	gen, err := NewBuilder().
		WithSdkID("test-sdk-id").
		WithKey(key).
		WithAuthAPIURL("https://custom.auth/v1/token").
		Build()
	assert.NilError(t, err)
	assert.Equal(t, gen.authAPIURL, "https://custom.auth/v1/token")
}

func TestBuilder_DefaultAuthURL(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	gen, err := NewBuilder().WithSdkID("test-sdk-id").WithKey(key).Build()
	assert.NilError(t, err)
	assert.Equal(t, gen.authAPIURL, DefaultAuthURL)
}

func TestBuilder_CustomJwtIdSupplier(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	gen, err := NewBuilder().
		WithSdkID("test-sdk-id").
		WithKey(key).
		WithJwtIDSupplier(func() string { return "fixed-id" }).
		Build()
	assert.NilError(t, err)
	assert.Equal(t, gen.jwtIDSupplier(), "fixed-id")
}

func TestGenerate_EmptyScopes(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	gen, _ := NewBuilder().WithSdkID("test-sdk-id").WithKey(key).Build()
	_, err := gen.Generate([]string{})
	assert.ErrorContains(t, err, "scopes must not be empty")
}

func TestGenerate_NilScopes(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")
	gen, _ := NewBuilder().WithSdkID("test-sdk-id").WithKey(key).Build()
	_, err := gen.Generate(nil)
	assert.ErrorContains(t, err, "scopes must not be empty")
}

func TestGenerate_Success(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			// Verify request method and content type
			assert.Equal(t, req.Method, http.MethodPost)
			assert.Equal(t, req.Header.Get("Content-Type"), "application/x-www-form-urlencoded")

			// Verify form body
			body, _ := io.ReadAll(req.Body)
			bodyStr := string(body)
			assert.Check(t, strings.Contains(bodyStr, "grant_type=client_credentials"))
			assert.Check(t, strings.Contains(bodyStr, "scope=scope1+scope2"))
			assert.Check(t, strings.Contains(bodyStr, "client_assertion_type="))
			assert.Check(t, strings.Contains(bodyStr, "client_assertion="))

			// Verify JWT structure (header.claims.signature)
			// Extract client_assertion from form data
			formValues, _ := io.ReadAll(strings.NewReader(bodyStr))
			parsedForm, _ := parseFormValues(string(formValues))
			jwtParts := strings.Split(parsedForm["client_assertion"], ".")
			assert.Equal(t, len(jwtParts), 3)

			return &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`{
					"access_token": "test-token-123",
					"expires_in": 3600,
					"token_type": "Bearer",
					"scope": "scope1 scope2"
				}`)),
			}, nil
		},
	}

	gen, _ := NewBuilder().
		WithSdkID("test-sdk-id").
		WithKey(key).
		WithAuthAPIURL("https://auth.test/v1/oauth/token").
		WithHTTPClient(mockClient).
		WithJwtIDSupplier(func() string { return "test-jwt-id" }).
		Build()

	resp, err := gen.Generate([]string{"scope1", "scope2"})
	assert.NilError(t, err)
	assert.Equal(t, resp.AccessToken, "test-token-123")
	assert.Equal(t, resp.ExpiresIn, 3600)
	assert.Equal(t, resp.TokenType, "Bearer")
	assert.Equal(t, resp.Scope, "scope1 scope2")
}

func TestGenerate_HTTPError(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 401,
				Body:       io.NopCloser(strings.NewReader(`{"error":"invalid_client"}`)),
			}, nil
		},
	}

	gen, _ := NewBuilder().
		WithSdkID("test-sdk-id").
		WithKey(key).
		WithAuthAPIURL("https://auth.test/v1/oauth/token").
		WithHTTPClient(mockClient).
		Build()

	_, err := gen.Generate([]string{"scope1"})
	assert.ErrorContains(t, err, "auth token request failed with HTTP 401")
}

func TestGenerate_NetworkError(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			return nil, io.ErrUnexpectedEOF
		},
	}

	gen, _ := NewBuilder().
		WithSdkID("test-sdk-id").
		WithKey(key).
		WithAuthAPIURL("https://auth.test/v1/oauth/token").
		WithHTTPClient(mockClient).
		Build()

	_, err := gen.Generate([]string{"scope1"})
	assert.ErrorContains(t, err, "auth token request failed")
}

func TestGenerate_InvalidResponseJSON(t *testing.T) {
	key := test.GetValidKey("../../test/test-key.pem")

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`not json`)),
			}, nil
		},
	}

	gen, _ := NewBuilder().
		WithSdkID("test-sdk-id").
		WithKey(key).
		WithAuthAPIURL("https://auth.test/v1/oauth/token").
		WithHTTPClient(mockClient).
		Build()

	_, err := gen.Generate([]string{"scope1"})
	assert.ErrorContains(t, err, "failed to deserialize authentication token response")
}

// parseFormValues is a test helper to parse URL-encoded form data
func parseFormValues(body string) (map[string]string, error) {
	result := make(map[string]string)
	pairs := strings.Split(body, "&")
	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			key, _ := strings.CutPrefix(parts[0], "")
			val, _ := strings.CutPrefix(parts[1], "")
			result[key] = val
		}
	}
	return result, nil
}
