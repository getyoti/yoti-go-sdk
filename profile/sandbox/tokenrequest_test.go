package sandbox

import (
	"encoding/json"
	"fmt"
	"time"
)

func AnchorList() []Anchor {
	return []Anchor{
		SourceAnchor("", time.Unix(1234567890, 0), ""),
		VerifierAnchor("", time.Unix(1234567890, 0), ""),
	}
}

func ExampleTokenRequest_WithRememberMeID() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithRememberMeID("some-remember-me-id")

	printJson(tokenRequest)
	// Output: {"remember_me_id":"some-remember-me-id","profile_attributes":null}
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
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"AttributeName1","value":"Value","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]},{"name":"AttributeName2","value":"Value","derivation":"","optional":"","anchors":[]}]}
}

func ExampleTokenRequest_WithAttributeStruct() {
	attribute := Attribute{
		Name:    "AttributeName3",
		Value:   "Value3",
		Anchors: AnchorList(),
	}

	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithAttributeStruct(attribute)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"AttributeName3","value":"Value3","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithGivenNames() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithGivenNames(
		"Value",
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"given_names","value":"Value","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithFamilyName() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithFamilyName(
		"Value",
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"family_name","value":"Value","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithFullName() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithFullName(
		"Value",
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"full_name","value":"Value","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithDateOfBirth() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithDateOfBirth(time.Unix(1234567890, 0), AnchorList())
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"date_of_birth","value":"2009-02-13","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithAgeVerification() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithAgeVerification(
		time.Unix(1234567890, 0),
		Derivation{}.AgeOver(18),
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"date_of_birth","value":"2009-02-13","derivation":"age_over:18","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithGender() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithGender("male", AnchorList())
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"gender","value":"male","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithPhoneNumber() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithPhoneNumber("00005550000", AnchorList())
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"phone_number","value":"00005550000","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithNationality() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithNationality("Value", AnchorList())
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"nationality","value":"Value","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithPostalAddress() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithPostalAddress("Value", AnchorList())
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"postal_address","value":"Value","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithStructuredPostalAddress() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithStructuredPostalAddress(
		map[string]interface{}{
			"FormattedAddressLine": "Value",
		},
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"structured_postal_address","value":"{\"FormattedAddressLine\":\"Value\"}","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithSelfie() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithSelfie(
		[]byte{0xDE, 0xAD, 0xBE, 0xEF},
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"selfie","value":"3q2+7w==","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithBase64Selfie() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithBase64Selfie(
		"3q2+7w==",
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"selfie","value":"3q2+7w==","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithEmailAddress() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithEmailAddress("user@example.com", AnchorList())
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"email_address","value":"user@example.com","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithDocumentDetails() {
	time.Local = time.UTC
	tokenRequest := TokenRequest{}.WithDocumentDetails(
		"DRIVING_LICENCE - abc1234",
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"document_details","value":"DRIVING_LICENCE - abc1234","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func ExampleTokenRequest_WithDocumentImages() {
	time.Local = time.UTC

	documentImages := DocumentImages{}.WithPngImage([]byte{0xDE, 0xAD, 0xBE, 0xEF}).WithJpegImage([]byte{0xDE, 0xAD, 0xBE, 0xEF})

	tokenRequest := TokenRequest{}.WithDocumentImages(
		documentImages,
		AnchorList(),
	)
	printJson(tokenRequest)
	// Output: {"remember_me_id":"","profile_attributes":[{"name":"document_images","value":"data:image/png;base64,3q2+7w==\u0026data:image/jpeg;base64,3q2+7w==","derivation":"","optional":"","anchors":[{"type":"SOURCE","value":"","sub_type":"","timestamp":1234567890000},{"type":"VERIFIER","value":"","sub_type":"","timestamp":1234567890000}]}]}
}

func printJson(value interface{}) {
	marshalledJSON, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(marshalledJSON))
}
