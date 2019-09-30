package attribute

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
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

func TestDocumentDetailsShouldParseDrivingLicenceWithoutExpiry(t *testing.T) {
	drivingLicenceGBR := "PASS_CARD GBR 1234abc - DVLA"

	details := DocumentDetails{}
	err := details.Parse(drivingLicenceGBR)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, details.DocumentType, "PASS_CARD")
	assert.Equal(t, details.DocumentNumber, "1234abc")
	assert.Assert(t, details.ExpirationDate == nil)
	assert.Equal(t, details.IssuingCountry, "GBR")
	assert.Equal(t, details.IssuingAuthority, "DVLA")
}

func TestDocumentDetailsShouldParseDrivingLicenceWithExtraAttribute(t *testing.T) {
	drivingLicenceGBR := "DRIVING_LICENCE GBR 1234abc 2016-05-01 DVLA someThirdAttribute"
	details := DocumentDetails{}
	err := details.Parse(drivingLicenceGBR)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, details.DocumentType, "DRIVING_LICENCE")
	assert.Equal(t, details.DocumentNumber, "1234abc")
	assert.Equal(t, details.ExpirationDate.Format("2006-01-02"), "2016-05-01")
	assert.Equal(t, details.IssuingCountry, "GBR")
	assert.Equal(t, details.IssuingAuthority, "DVLA")
}

func TestDocumentDetailsShouldParseDrivingLicenceWithAllOptionalAttributes(t *testing.T) {
	drivingLicenceGBR := "DRIVING_LICENCE GBR 1234abc 2016-05-01 DVLA"

	details := DocumentDetails{}
	err := details.Parse(drivingLicenceGBR)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, details.DocumentType, "DRIVING_LICENCE")
	assert.Equal(t, details.DocumentNumber, "1234abc")
	assert.Equal(t, details.ExpirationDate.Format("2006-01-02"), "2016-05-01")
	assert.Equal(t, details.IssuingCountry, "GBR")
	assert.Equal(t, details.IssuingAuthority, "DVLA")
}

func TestDocumentDetailsShouldParseAadhaar(t *testing.T) {
	aadhaar := "AADHAAR IND 1234abc 2016-05-01"

	details := DocumentDetails{}
	err := details.Parse(aadhaar)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, details.DocumentType, "AADHAAR")
	assert.Equal(t, details.DocumentNumber, "1234abc")
	assert.Equal(t, details.ExpirationDate.Format("2006-01-02"), "2016-05-01")
	assert.Equal(t, details.IssuingCountry, "IND")
}

func TestDocumentDetailsShouldParsePassportWithMandatoryFieldsOnly(t *testing.T) {
	passportGBR := "PASSPORT GBR 1234abc"

	details := DocumentDetails{}
	err := details.Parse(passportGBR)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, details.DocumentType, "PASSPORT")
	assert.Equal(t, details.DocumentNumber, "1234abc")
	assert.Assert(t, details.ExpirationDate == nil)
	assert.Equal(t, details.IssuingCountry, "GBR")
	assert.Equal(t, details.IssuingAuthority, "")
}

func TestDocumentDetailsShouldErrorOnEmptyString(t *testing.T) {
	empty := ""

	details := DocumentDetails{}
	err := details.Parse(empty)
	assert.ErrorContains(t, err, "Document Details data is invalid")
}

func TestDocumentDetailsShouldErrorIfLessThan3Words(t *testing.T) {
	corrupt := "PASS_CARD GBR"
	details := DocumentDetails{}
	err := details.Parse(corrupt)
	assert.ErrorContains(t, err, "Document Details data is invalid")
}

func TestDocumentDetailsShouldErrorForInvalidCountry(t *testing.T) {
	corrupt := "PASSPORT 13 1234abc 2016-05-01"
	details := DocumentDetails{}
	err := details.Parse(corrupt)
	assert.ErrorContains(t, err, "Document Details data is invalid")
}

func TestDocumentDetailsShouldErrorForInvalidNumber(t *testing.T) {
	corrupt := "PASSPORT GBR $%^$%^£ 2016-05-01"
	details := DocumentDetails{}
	err := details.Parse(corrupt)
	assert.ErrorContains(t, err, "Document Details data is invalid")
}

func TestDocumentDetailsShouldErrorForInvalidExpirationDate(t *testing.T) {
	corrupt := "PASSPORT GBR 1234abc X016-05-01"
	details := DocumentDetails{}
	err := details.Parse(corrupt)
	fmt.Printf("!DEBUG! %v\n", details.ExpirationDate)
	assert.ErrorContains(t, err, "cannot parse")
}
