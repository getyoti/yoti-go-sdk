package sandbox

import (
	"fmt"
	"time"
)

func AnchorList() []Anchor {
	return []Anchor{
		SourceAnchor("", time.Unix(1234567890, 0), ""),
		VerifierAnchor("", time.Unix(1234567890, 0), ""),
	}
}

func ExampleTokenRequest_WithAttribute() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithAttribute(
		"AttributeName1",
		"Value",
		AnchorList(),
	).WithAttribute(
		"AttributeName2",
		"Value",
		nil,
	)
	fmt.Println(tokenRequest)
	// Output: { [{AttributeName1 Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]} {AttributeName2 Value   []}]}
}

func ExampleTokenRequest_WithGivenNames() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithGivenNames(
		"Value",
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{given_names Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithFamilyName() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithFamilyName(
		"Value",
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{family_name Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithFullName() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithFullName(
		"Value",
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{full_name Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithDateOfBirth() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithDateOfBirth(time.Unix(1234567890, 0), AnchorList())
	fmt.Println(tokenRequest)
	// Output: { [{date_of_birth 2009-02-13   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithAgeVerification() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithAgeVerification(
		time.Unix(1234567890, 0),
		Derivation{}.AgeOver(18),
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{date_of_birth 2009-02-13 age_over:18  [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithGender() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithGender("male", AnchorList())
	fmt.Println(tokenRequest)
	// Output: { [{gender male   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithPhoneNumber() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithPhoneNumber("00005550000", AnchorList())
	fmt.Println(tokenRequest)
	// Output: { [{phone_number 00005550000   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithNationality() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithNationality("Value", AnchorList())
	fmt.Println(tokenRequest)
	// Output: { [{nationality Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithPostalAddress() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithPostalAddress("Value", AnchorList())
	fmt.Println(tokenRequest)
	// Output: { [{postal_address Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithStructuredPostalAddress() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithStructuredPostalAddress(
		map[string]interface{}{
			"FormattedAddressLine": "Value",
		},
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{structured_postal_address {"FormattedAddressLine":"Value"}   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithSelfie() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithSelfie(
		[]byte{0xDE, 0xAD, 0xBE, 0xEF},
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{selfie 3q2+7w==   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithBase64Selfie() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithBase64Selfie(
		"3q2+7w==",
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{selfie 3q2+7w==   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithEmailAddress() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithEmailAddress("user@example.com", AnchorList())
	fmt.Println(tokenRequest)
	// Output: { [{email_address user@example.com   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleTokenRequest_WithDocumentDetails() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithDocumentDetails(
		"DRIVING_LICENCE - abc1234",
		AnchorList(),
	)
	fmt.Println(tokenRequest)
	// Output: { [{document_details DRIVING_LICENCE - abc1234   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}
