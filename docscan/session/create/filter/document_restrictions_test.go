package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedDocumentRestrictionBuilder() {
	restriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentType("PASSPORT").
		WithCountryCode("FRA").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(restriction)
	fmt.Println(string(data))
	// Output: {"document_types":["PASSPORT"],"country_codes":["FRA"]}
}
