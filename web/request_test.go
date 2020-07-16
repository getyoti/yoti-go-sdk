package web

import (
	"net/http"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestYotiClient_NilHTTPClientShouldUse10sTimeout(t *testing.T) {
	result := ensureHttpClientTimeout(nil).(*http.Client)

	assert.Equal(t, 10*time.Second, result.Timeout)
}

func TestYotiClient_ShouldSetHTTPClientTimeout(t *testing.T) {
	httpClient := &http.Client{
		Timeout: time.Minute * 12,
	}
	result := ensureHttpClientTimeout(httpClient).(*http.Client)

	assert.Equal(t, 12*time.Minute, result.Timeout)
}
