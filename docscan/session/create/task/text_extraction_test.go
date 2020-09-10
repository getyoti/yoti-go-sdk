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

	data, _ := json.Marshal(task)
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
