package filter

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRequestedDocumentRestrictionBuilder() {
	restriction, err := NewRequestedDocumentRestrictionBuilder().
		WithDocumentType("PASSPORT").
		WithCountryCode("FRA").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(restriction)
	fmt.Println(string(data))
	// Output: {"document_types":["PASSPORT"],"country_codes":["FRA"]}
}
