package docscanerr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// DataObject maps from Doc Scan JSON error responses
type DataObject struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Error   []ErrorItemDO `json:"error,omitempty"`
}

// ErrorItemDO maps from Doc Scan JSON error items
type ErrorItemDO struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

// Error indicates errors related to the Doc Scan API.
type Error struct {
	message  string
	Err      error
	Response *http.Response
}

// New creates a new Doc Scan Error
func New(err error, response *http.Response) *Error {
	return &Error{
		message:  formatResponseMessage(err, response),
		Err:      err,
		Response: response,
	}
}

// Error return the error message
func (e Error) Error() string {
	return e.message
}

// Unwrap returns the internal error for debugging
func (e Error) Unwrap() error {
	return e.Err
}

func formatResponseMessage(err error, response *http.Response) string {
	if response == nil || response.Body == nil {
		return err.Error()
	}

	body, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var errorDO DataObject
	jsonError := json.Unmarshal(body, &errorDO)

	if jsonError != nil || errorDO.Code == "" || errorDO.Message == "" {
		return err.Error()
	}

	formattedCodeMessage := fmt.Sprintf("%d: %s - %s", response.StatusCode, errorDO.Code, errorDO.Message)

	formattedItems := []string{}
	for _, item := range errorDO.Error {
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
