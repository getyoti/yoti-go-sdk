package task

import (
	"encoding/json"
	"fmt"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox"
)

func Example_documentTextDataExtractionTaskBuilder() {
	filter, _ := sandbox.NewDocumentFilterBuilder().Build()

	check, _ := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFilter(filter).
		WithDocumentField("some", "field").
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":null,"country_codes":null},"result":{"document_fields":{"some":"field"}}}
}
