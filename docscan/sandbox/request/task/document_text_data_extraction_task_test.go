package task

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleDocumentTextDataExtractionTaskBuilder() {
	filter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		return
	}

	check, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFilter(filter).
		WithDocumentField("some", "field").
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some":"field"}}}
}
