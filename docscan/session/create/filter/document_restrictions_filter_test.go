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

	data, err := json.Marshal(docFilter)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

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

	data, err := json.Marshal(docFilter)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"BLACKLIST","documents":[{"document_types":["PASSPORT"]}]}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_withExpiredDocuments() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithExpiredDocuments(true).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(restriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allow_expired_documents":true}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_withDenyExpiredDocuments() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithExpiredDocuments(false).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(restriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allow_expired_documents":false}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_withAllowNonLatinDocuments() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithAllowNonLatinDocuments(true).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(restriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allow_non_latin_documents":true}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_withDenyNonLatinDocuments() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithAllowNonLatinDocuments(false).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(restriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allow_non_latin_documents":false}
}
