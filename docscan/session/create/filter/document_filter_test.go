package filter

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedDocumentFilterBuilder() {
	filter, err := NewRequestedDocumentFilterBuilder().
		WithCountryCodes([]string{"AUS", "GBR", "USA", "UKR"}).
		WithDocumentType("DRIVING_LICENCE").
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(filter)
	fmt.Println(string(data))
	// Output: {"document_types":["DRIVING_LICENCE"],"country_codes":["AUS","GBR","USA","UKR"]}
}
