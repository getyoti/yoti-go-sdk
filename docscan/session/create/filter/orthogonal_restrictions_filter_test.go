package filter

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_WithIncludedCountries() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithIncludedCountries([]string{"KEN", "CIV"}).
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","country_restriction":{"inclusion":"WHITELIST","country_codes":["KEN","CIV"]}}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_WithIncludedDocumentTypes() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithIncludedDocumentTypes([]string{"PASSPORT", "DRIVING_LICENCE"}).
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","type_restriction":{"inclusion":"WHITELIST","document_types":["PASSPORT","DRIVING_LICENCE"]}}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_WithExcludedDocumentTypes() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithExcludedDocumentTypes([]string{"NATIONAL_ID", "PASS_CARD"}).
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","type_restriction":{"inclusion":"BLACKLIST","document_types":["NATIONAL_ID","PASS_CARD"]}}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_WithExcludedCountries() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithExcludedCountries([]string{"CAN", "FJI"}).
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","country_restriction":{"inclusion":"BLACKLIST","country_codes":["CAN","FJI"]}}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withExpiredDocuments() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allow_expired_documents":true}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withDenyExpiredDocuments() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allow_expired_documents":false}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withNonLatinDocuments() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithNonLatinDocuments(true).
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allow_non_latin_documents":true}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withDenyNonLatinDocuments() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
		WithNonLatinDocuments(false).
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allow_non_latin_documents":false}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withAllowDigitalIDs() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allow_digital_ids":true}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withAllowedProviders() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allowed_providers":[{"name":"DIGILOCKER"},{"name":"EPHIL_ID_QR"}]}
}

func ExampleRequestedOrthogonalRestrictionsFilterBuilder_withAllowedProvider() {
	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
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
	// Output: {"type":"ORTHOGONAL_RESTRICTIONS","allowed_providers":[{"name":"DIGILOCKER"},{"name":"EPHIL_ID_QR"}]}
}

// TestRequestedOrthogonalRestrictionsFilterBuilder_WithAllowedProviders_DoesNotAliasCallerSlice guards
// against the built filter changing if the caller mutates (via append) the slice they passed in,
// which can happen silently when the passed-in slice has spare capacity.
func TestRequestedOrthogonalRestrictionsFilterBuilder_WithAllowedProviders_DoesNotAliasCallerSlice(t *testing.T) {
	callerSlice := make([]*DigitalIDProvider, 2, 4)
	callerSlice[0] = &DigitalIDProvider{Name: "A"}
	callerSlice[1] = &DigitalIDProvider{Name: "B"}

	restriction, err := NewRequestedOrthogonalRestrictionsFilterBuilder().
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
