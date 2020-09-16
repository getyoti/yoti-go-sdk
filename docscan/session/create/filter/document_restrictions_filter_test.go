package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedDocumentRestrictionsFilterBuilder_ForIncludeList() {
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

	data, _ := json.Marshal(docFilter)
	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"WHITELIST","documents":[{"document_types":["PASSPORT"]}]}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_ForExcludeList() {
	docRestriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentTypes([]string{"PASSPORT"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	var docFilter *RequestedDocumentRestrictionsFilter
	docFilter, err = NewRequestedDocumentRestrictionsFilterBuilder().
		ForExcludeList().
		WithDocumentRestriction(docRestriction).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(docFilter)
	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"BLACKLIST","documents":[{"document_types":["PASSPORT"]}]}
}
