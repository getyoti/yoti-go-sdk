package yotierror

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

var (
	expectedErr = Error{
		Id:        "8f6a9dfe72128de20909af0d476769b6",
		Status:    401,
		ErrorCode: "INVALID_REQUEST_SIGNATURE",
		Message:   "Invalid request signature",
	}
)

func TestError_ShouldReturnFormattedError(t *testing.T) {
	jsonBytes := json.RawMessage(`{"id":"8f6a9dfe72128de20909af0d476769b6","status":401,"error":"INVALID_REQUEST_SIGNATURE","message":"Invalid request signature"}`)

	err := NewResponseError(
		&http.Response{
			StatusCode: 401,
			Body:       io.NopCloser(bytes.NewReader(jsonBytes)),
		},
	)

	assert.ErrorIs(t, *err, expectedErr)
}

func TestError_ShouldReturnFormattedError_ReturnWrappedErrorWhenInvalidJSON(t *testing.T) {
	response := &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader("some invalid JSON")),
	}
	err := NewResponseError(
		response,
	)

	assert.ErrorContains(t, err, "unknown HTTP error")
}

func TestError_ShouldReturnTemporaryForServerError(t *testing.T) {
	response := &http.Response{
		StatusCode: 500,
	}
	err := NewResponseError(
		response,
	)

	assert.Check(t, err.Temporary())
}

func TestError_ShouldNotReturnTemporaryForClientError(t *testing.T) {
	response := &http.Response{
		StatusCode: 400,
	}
	err := NewResponseError(
		response,
	)

	assert.Check(t, !err.Temporary())
}
