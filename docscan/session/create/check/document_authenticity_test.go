package check

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedDocumentAuthenticityCheckBuilder() {
	docAuthCheck, err := NewRequestedDocumentAuthenticityCheckBuilder().WithManualCheckFallback().Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(docAuthCheck)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_AUTHENTICITY","config":{"manual_check":"FALLBACK"}}
}

func ExampleRequestedDocumentAuthenticityCheckBuilder_Build() {
	docAuthCheck, err := NewRequestedDocumentAuthenticityCheckBuilder().Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(docAuthCheck)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_AUTHENTICITY","config":{}}
}

func TestRequestedDocumentAuthenticityCheckBuilder_WithManualCheckAlways(t *testing.T) {
	docAuthCheck, err := NewRequestedDocumentAuthenticityCheckBuilder().
		WithManualCheckAlways().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	result := docAuthCheck.Config().(RequestedDocumentAuthenticityConfig)
	assert.Equal(t, "ALWAYS", result.ManualCheck)
}

func TestRequestedDocumentAuthenticityCheckBuilder_WithManualCheckFallback(t *testing.T) {
	docAuthCheck, err := NewRequestedDocumentAuthenticityCheckBuilder().
		WithManualCheckFallback().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	result := docAuthCheck.Config().(RequestedDocumentAuthenticityConfig)
	assert.Equal(t, "FALLBACK", result.ManualCheck)
}

func TestRequestedDocumentAuthenticityCheckBuilder_WithManualCheckNever(t *testing.T) {
	docAuthCheck, err := NewRequestedDocumentAuthenticityCheckBuilder().
		WithManualCheckNever().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	result := docAuthCheck.Config().(RequestedDocumentAuthenticityConfig)
	assert.Equal(t, "NEVER", result.ManualCheck)
}

func TestRequestedDocumentAuthenticityCheckBuilder_UsesLastValue(t *testing.T) {
	docAuthCheck, err := NewRequestedDocumentAuthenticityCheckBuilder().
		WithManualCheckFallback().
		WithManualCheckNever().
		WithManualCheckAlways().
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	result := docAuthCheck.Config().(RequestedDocumentAuthenticityConfig)
	assert.Equal(t, "ALWAYS", result.ManualCheck)
}
