package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedDocumentRestriction() {
	docRestriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentTypes([]string{"PASSPORT"}).
		WithCountryCodes([]string{"GBR"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(docRestriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_types":["PASSPORT"],"country_codes":["GBR"]}
}

func ExampleRequestedDocumentRestrictionBuilder_WithDocumentTypes() {
	docRestriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentTypes([]string{"PASSPORT"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(docRestriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_types":["PASSPORT"]}
}

func ExampleRequestedDocumentRestrictionBuilder_WithCountryCodes() {
	docRestriction, err := NewRequestedDocumentRestrictionBuilder().
		WithCountryCodes([]string{"GBR"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(docRestriction)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"country_codes":["GBR"]}
}
