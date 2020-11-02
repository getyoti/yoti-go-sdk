package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleRequiredIDDocument_MarshalJSON() {
	docRestriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentTypes([]string{"PASSPORT"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	var docFilter *RequestedDocumentRestrictionsFilter
	docFilter, err = NewRequestedDocumentRestrictionsFilterBuilder().
		ForIncludeList().
		WithDocumentRestriction(docRestriction).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	var requiredIDDocument *RequiredIDDocument
	requiredIDDocument, err = NewRequiredIDDocumentBuilder().
		WithFilter(docFilter).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(requiredIDDocument)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT","filter":{"type":"DOCUMENT_RESTRICTIONS","inclusion":"WHITELIST","documents":[{"document_types":["PASSPORT"]}]}}
}
