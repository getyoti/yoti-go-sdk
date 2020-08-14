package task

import (
	"encoding/json"
	"fmt"
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
