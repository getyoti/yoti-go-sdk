package filter

import (
	"encoding/json"
	"fmt"
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
