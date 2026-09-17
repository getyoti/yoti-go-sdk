package filter

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
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

func ExampleRequestedDocumentRestrictionsFilterBuilder_withAllowDigitalIDs() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithAllowDigitalIDs(true).
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
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allow_digital_ids":true}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_withAllowedProviders() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithAllowedProviders([]*DigitalIDProvider{
			{Name: "DIGILOCKER"},
			{Name: "EPHIL_ID_QR"},
		}).
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
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allowed_providers":[{"name":"DIGILOCKER"},{"name":"EPHIL_ID_QR"}]}
}

func ExampleRequestedDocumentRestrictionsFilterBuilder_withAllowedProvider() {
	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithAllowedProvider(&DigitalIDProvider{Name: "DIGILOCKER"}).
		WithAllowedProvider(&DigitalIDProvider{Name: "EPHIL_ID_QR"}).
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
	// Output: {"type":"DOCUMENT_RESTRICTIONS","inclusion":"","documents":[],"allowed_providers":[{"name":"DIGILOCKER"},{"name":"EPHIL_ID_QR"}]}
}

// TestRequestedDocumentRestrictionsFilterBuilder_WithAllowedProviders_DoesNotAliasCallerSlice guards
// against the built filter changing if the caller mutates (via append) the slice they passed in,
// which can happen silently when the passed-in slice has spare capacity.
func TestRequestedDocumentRestrictionsFilterBuilder_WithAllowedProviders_DoesNotAliasCallerSlice(t *testing.T) {
	callerSlice := make([]*DigitalIDProvider, 2, 4)
	callerSlice[0] = &DigitalIDProvider{Name: "A"}
	callerSlice[1] = &DigitalIDProvider{Name: "B"}

	restriction, err := NewRequestedDocumentRestrictionsFilterBuilder().
		WithAllowedProviders(callerSlice).
		WithAllowedProvider(&DigitalIDProvider{Name: "C"}).
		Build()
	assert.NilError(t, err)

	before, err := json.Marshal(restriction)
	assert.NilError(t, err)

	// Mutating the caller's own slice must not affect the already-built filter.
	_ = append(callerSlice, &DigitalIDProvider{Name: "should-not-appear"})

	after, err := json.Marshal(restriction)
	assert.NilError(t, err)

	assert.Equal(t, string(before), string(after))
}
