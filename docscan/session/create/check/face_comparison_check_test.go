package check

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedFaceComparisonCheckBuilder() {
	check, err := NewRequestedFaceComparisonCheckBuilder().
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
	check, err := NewRequestedFaceComparisonCheckBuilder().
		WithManualCheckAlways().
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	config := check.Config().(RequestedFaceComparisonConfig)
	assert.Equal(t, "ALWAYS", config.ManualCheck)
}

func TestRequestedFaceComparisonCheckBuilder_WithManualCheckFallback(t *testing.T) {
	check, err := NewRequestedFaceComparisonCheckBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	config := check.Config().(RequestedFaceComparisonConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedFaceComparisonCheckBuilder_WithManualCheckNever(t *testing.T) {
	check, err := NewRequestedFaceComparisonCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	config := check.Config().(RequestedFaceComparisonConfig)
	assert.Equal(t, "NEVER", config.ManualCheck)
}

func TestRequestedFaceComparisonCheckBuilder_DefaultManualCheckEmpty(t *testing.T) {
	builder := NewRequestedFaceComparisonCheckBuilder()

	check, err := builder.Build()
	assert.NilError(t, err)
	assert.Equal(t, check.config.ManualCheck, "") // default is empty string
}
