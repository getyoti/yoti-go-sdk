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

func ExampleProfile_WithAttribute() {
	time.Local = time.UTC
	profile := Profile{}.WithAttribute(
		"AttributeName1",
		"Value",
		AnchorList(),
	).WithAttribute(
		"AttributeName2",
		"Value",
		nil,
	)
	fmt.Println(profile)
	// Output: { [{AttributeName1 Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]} {AttributeName2 Value   []}]}
}

func ExampleProfile_WithGivenNames() {
	time.Local = time.UTC
	profile := Profile{}.WithGivenNames(
		"Value",
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{given_names Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithFamilyName() {
	time.Local = time.UTC
	profile := Profile{}.WithFamilyName(
		"Value",
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{family_name Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithFullName() {
	time.Local = time.UTC
	profile := Profile{}.WithFullName(
		"Value",
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{full_name Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithDateOfBirth() {
	time.Local = time.UTC
	profile := Profile{}.WithDateOfBirth(time.Unix(1234567890, 0), AnchorList())
	fmt.Println(profile)
	// Output: { [{date_of_birth 2009-02-13   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithAgeVerification() {
	time.Local = time.UTC
	profile := Profile{}.WithAgeVerification(
		time.Unix(1234567890, 0),
		Derivation{}.AgeOver(18),
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{date_of_birth 2009-02-13 age_over:18  [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithGender() {
	time.Local = time.UTC
	profile := Profile{}.WithGender("male", AnchorList())
	fmt.Println(profile)
	// Output: { [{gender male   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithPhoneNumber() {
	time.Local = time.UTC
	profile := Profile{}.WithPhoneNumber("00005550000", AnchorList())
	fmt.Println(profile)
	// Output: { [{phone_number 00005550000   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithNationality() {
	time.Local = time.UTC
	profile := Profile{}.WithNationality("Value", AnchorList())
	fmt.Println(profile)
	// Output: { [{nationality Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithPostalAddress() {
	time.Local = time.UTC
	profile := Profile{}.WithPostalAddress("Value", AnchorList())
	fmt.Println(profile)
	// Output: { [{postal_address Value   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithStructuredPostalAddress() {
	time.Local = time.UTC
	profile := Profile{}.WithStructuredPostalAddress(
		map[string]string{
			"FormattedAddressLine": "Value",
		},
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{structured_postal_address {"FormattedAddressLine":"Value"}   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithSelfie() {
	time.Local = time.UTC
	profile := Profile{}.WithSelfie(
		[]byte{0xDE, 0xAD, 0xBE, 0xEF},
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{selfie 3q2+7w==   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithEmailAddress() {
	time.Local = time.UTC
	profile := Profile{}.WithEmailAddress("user@example.com", AnchorList())
	fmt.Println(profile)
	// Output: { [{email_address user@example.com   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithDocumentDetails() {
	time.Local = time.UTC
	profile := Profile{}.WithDocumentDetails(
		"DRIVING_LICENCE - abc1234",
		AnchorList(),
	)
	fmt.Println(profile)
	// Output: { [{document_details DRIVING_LICENCE - abc1234   [{SOURCE   2009-02-13 23:31:30 +0000 UTC} {VERIFIER   2009-02-13 23:31:30 +0000 UTC}]}]}
}

func ExampleProfile_WithoutAttributes() {
	profile := Profile{}.WithoutAttributes()
	fmt.Println(profile)
	// Output: { [{unused unused   []}]}
}
