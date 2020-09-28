package check

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedFaceMatchCheckBuilder() {
	check, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_FACE_MATCH","config":{"manual_check":"NEVER"}}
}

func TestRequestedFaceMatchCheckBuilder_WithManualCheckAlways(t *testing.T) {
	task, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckAlways().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedFaceMatchConfig)
	assert.Equal(t, "ALWAYS", config.ManualCheck)
}

func TestRequestedFaceMatchCheckBuilder_WithManualCheckFallback(t *testing.T) {
	task, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedFaceMatchConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedFaceMatchCheckBuilder_WithManualCheckNever(t *testing.T) {
	task, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedFaceMatchConfig)
	assert.Equal(t, "NEVER", config.ManualCheck)
}
