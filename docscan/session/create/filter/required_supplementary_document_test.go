package filter

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/objective"
)

func ExampleRequiredSupplementaryDocument() {
	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err := NewRequiredSupplementaryDocumentBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT"}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithFilter() {
	docRestriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentTypes([]string{"UTILITY_BILL"}).
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

	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err = NewRequiredSupplementaryDocumentBuilder().
		WithFilter(docFilter).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT","filter":{"type":"DOCUMENT_RESTRICTIONS","inclusion":"WHITELIST","documents":[{"document_types":["UTILITY_BILL"]}]}}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithCountryCodes() {
	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err := NewRequiredSupplementaryDocumentBuilder().
		WithCountryCodes([]string{"SOME_COUNTRY"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT","country_codes":["SOME_COUNTRY"]}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithCountryCodes_empty() {
	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err := NewRequiredSupplementaryDocumentBuilder().
		WithCountryCodes([]string{}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT"}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithDocumentTypes() {
	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err := NewRequiredSupplementaryDocumentBuilder().
		WithDocumentTypes([]string{"SOME_DOCUMENT_TYPE"}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT","document_types":["SOME_DOCUMENT_TYPE"]}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithDocumentTypes_empty() {
	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err := NewRequiredSupplementaryDocumentBuilder().
		WithDocumentTypes([]string{}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT"}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithObjective() {
	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err := NewRequiredSupplementaryDocumentBuilder().
		WithObjective(&mockObjective{}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT","objective":{"type":"SOME_OBJECTIVE"}}
}

func ExampleRequiredSupplementaryDocumentBuilder_WithObjective_proofOfAddress() {
	var proofOfAddress objective.Objective
	proofOfAddress, err := objective.NewProofOfAddressObjectiveBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	var requiredSupplementaryDocument *RequiredSupplementaryDocument
	requiredSupplementaryDocument, err = NewRequiredSupplementaryDocumentBuilder().
		WithObjective(proofOfAddress).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(requiredSupplementaryDocument)
	fmt.Println(string(data))
	// Output: {"type":"SUPPLEMENTARY_DOCUMENT","objective":{"type":"PROOF_OF_ADDRESS"}}
}

type mockObjective struct{}

func (o *mockObjective) Type() string {
	return "SOME_OBJECTIVE"
}

// MarshalJSON returns the JSON encoding
func (o *mockObjective) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
	}{
		Type: o.Type(),
	})
}
