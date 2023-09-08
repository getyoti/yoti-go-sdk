package yotierror

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	defaultUnknownErrorMessageConst = "unknown HTTP error"

	// DefaultHTTPErrorMessages maps HTTP error status codes to format strings
	// to create useful error messages. -1 is used to specify a default message
	// that can be used if an error code is not explicitly defined
	DefaultHTTPErrorMessages = map[int]string{
		-1: defaultUnknownErrorMessageConst,
	}
)

// DataObject maps from JSON error responses
type DataObject struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Errors  []ItemDataObject `json:"errors,omitempty"`
}

// ItemDataObject maps from JSON error items
type ItemDataObject struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

// Error indicates errors related to the Yoti API.
type Error struct {
	message  string
	Err      error
	Response *http.Response
	id       string
	status   int
}

// NewResponseError creates a new Error
func NewResponseError(response *http.Response, httpErrorMessages ...map[int]string) *Error {
	return &Error{
		message:  formatResponseMessage(response, httpErrorMessages...),
		Response: response,
	}
}

// Error return the error message
func (e Error) Error() string {
	return e.message
}

// Temporary indicates this error is a temporary error
func (e Error) Temporary() bool {
	return e.Response != nil && e.Response.StatusCode >= 500
}

func formatResponseMessage(response *http.Response, httpErrorMessages ...map[int]string) string {
	defaultMessage := handleHTTPError(response, httpErrorMessages...)

	if response == nil || response.Body == nil {
		return defaultMessage
	}

	body, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(bytes.NewBuffer(body))

	var errorDO DataObject
	jsonError := json.Unmarshal(body, &errorDO)

	if jsonError != nil || errorDO.Code == "" || errorDO.Message == "" {
		return defaultMessage
	}

	formattedCodeMessage := fmt.Sprintf("%d: %s - %s", response.StatusCode, errorDO.Code, errorDO.Message)

	var formattedItems []string
	for _, item := range errorDO.Errors {
		if item.Message != "" && item.Property != "" {
			formattedItems = append(
				formattedItems,
				fmt.Sprintf("%s: `%s`", item.Property, item.Message),
			)
		}
	}

	if len(formattedItems) > 0 {
		return fmt.Sprintf("%s: %s", formattedCodeMessage, strings.Join(formattedItems, ", "))
	}

	return formattedCodeMessage
}

func formatHTTPError(message string, statusCode int, body []byte) string {
	if len(body) == 0 {
		return fmt.Sprintf("%d: %s", statusCode, message)
	}
	return fmt.Sprintf("%d: %s - %s", statusCode, message, body)
}

func handleHTTPError(response *http.Response, errorMessages ...map[int]string) string {
	var body []byte
	if response.Body != nil {
		body, _ = io.ReadAll(response.Body)
		response.Body = io.NopCloser(bytes.NewBuffer(body))
	} else {
		body = make([]byte, 0)
	}
	for _, handler := range errorMessages {
		for code, message := range handler {
			if code == response.StatusCode {
				return formatHTTPError(
					message,
					response.StatusCode,
					body,
				)
			}

		}
		if defaultMessage, ok := handler[-1]; ok {
			return formatHTTPError(
				defaultMessage,
				response.StatusCode,
				body,
			)
		}

	}
	return formatHTTPError(
		defaultUnknownErrorMessageConst,
		response.StatusCode,
		body,
	)
}
