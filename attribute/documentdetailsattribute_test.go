package attribute

import (
	"fmt"
)

func ExampleDocumentDetails_Parse() {
	raw := "PASSPORT GBR 1234567 2022-09-12"
	details := DocumentDetails{}
	err := details.Parse(raw)
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"Document Type: %s, Issuing Country: %s, Document Number: %s, Expiration Date: %s",
		details.DocumentType,
		details.IssuingCountry,
		details.DocumentNumber,
		details.ExpirationDate,
	)
	// Output: Document Type: PASSPORT, Issuing Country: GBR, Document Number: 1234567, Expiration Date: 2022-09-12 00:00:00 +0000 UTC
}
