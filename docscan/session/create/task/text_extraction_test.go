package task

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedTextExtractionTaskBuilder() {
	check, err := NewRequestedTextExtractionTaskBuilder().
		WithManualCheckAlways().
		WithChipDataIgnore().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_TEXT_DATA_EXTRACTION","config":{"manual_check":"ALWAYS","chip_data":"IGNORE"}}
}
