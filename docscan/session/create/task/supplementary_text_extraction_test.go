package task

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedSupplementaryDocTextExtractionTaskBuilder() {
	task, err := NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckAlways().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_EXTRACTION","config":{"manual_check":"ALWAYS"}}
}

func TestRequestedSupplementaryDocTextExtractionTaskBuilder_Build_UsesLastManualCheck(t *testing.T) {
	task, err := NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckAlways().
		WithManualCheckNever().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedSupplementaryDocTextExtractionTaskConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedSupplementaryDocTextExtractionTaskBuilder_WithManualCheckAlways(t *testing.T) {
	task, err := NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckAlways().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedSupplementaryDocTextExtractionTaskConfig)
	assert.Equal(t, "ALWAYS", config.ManualCheck)
}

func TestRequestedSupplementaryDocTextExtractionTaskBuilder_WithManualCheckFallback(t *testing.T) {
	task, err := NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedSupplementaryDocTextExtractionTaskConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedSupplementaryDocTextExtractionTaskBuilder_WithManualCheckNever(t *testing.T) {
	task, err := NewRequestedSupplementaryDocTextExtractionTaskBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedSupplementaryDocTextExtractionTaskConfig)
	assert.Equal(t, "NEVER", config.ManualCheck)
}
