package check

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedFaceComparisonCheckBuilder() {
	check, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(check)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"FACE_COMPARISON","config":{"manual_check":"NEVER"}}
}

func TestRequestedFaceComparisonCheckBuilder_WithManualCheckAlways(t *testing.T) {
	task, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckAlways().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedFaceMatchConfig)
	assert.Equal(t, "ALWAYS", config.ManualCheck)
}

func TestRequestedFaceComparisonCheckBuilder_WithManualCheckFallback(t *testing.T) {
	task, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedFaceMatchConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedFaceComparisonCheckBuilder_WithManualCheckNever(t *testing.T) {
	task, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedFaceMatchConfig)
	assert.Equal(t, "NEVER", config.ManualCheck)
}
