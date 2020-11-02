package task

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedTextExtractionTaskBuilder() {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithManualCheckAlways().
		WithChipDataIgnore().
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
	// Output: {"type":"ID_DOCUMENT_TEXT_DATA_EXTRACTION","config":{"manual_check":"ALWAYS","chip_data":"IGNORE"}}
}

func TestRequestedTextExtractionTaskBuilder_Build_UsesLastManualCheck(t *testing.T) {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithManualCheckAlways().
		WithManualCheckNever().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedTextExtractionTaskConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedTextExtractionTaskBuilder_WithManualCheckAlways(t *testing.T) {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithManualCheckAlways().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedTextExtractionTaskConfig)
	assert.Equal(t, "ALWAYS", config.ManualCheck)
}

func TestRequestedTextExtractionTaskBuilder_WithManualCheckFallback(t *testing.T) {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithManualCheckFallback().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedTextExtractionTaskConfig)
	assert.Equal(t, "FALLBACK", config.ManualCheck)
}

func TestRequestedTextExtractionTaskBuilder_WithManualCheckNever(t *testing.T) {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedTextExtractionTaskConfig)
	assert.Equal(t, "NEVER", config.ManualCheck)
}

func TestRequestedTextExtractionTaskBuilder_WithChipDataDesired(t *testing.T) {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithChipDataDesired().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedTextExtractionTaskConfig)
	assert.Equal(t, "DESIRED", config.ChipData)
}

func TestRequestedTextExtractionTaskBuilder_WithChipDataIgnore(t *testing.T) {
	task, err := NewRequestedTextExtractionTaskBuilder().
		WithChipDataIgnore().
		Build()
	if err != nil {
		t.Fail()
	}

	config := task.Config().(RequestedTextExtractionTaskConfig)
	assert.Equal(t, "IGNORE", config.ChipData)
}
