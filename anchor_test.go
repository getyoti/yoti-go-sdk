package yoti

import (
	"io/ioutil"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func assertServerCertSerialNo(t *testing.T, expectedSerialNo string, actualSerialNo *big.Int) {
	expectedSerialNoBigInt := new(big.Int)
	expectedSerialNoBigInt, ok := expectedSerialNoBigInt.SetString(expectedSerialNo, 10)
	assert.Assert(t, ok, "Unexpected error when setting string as big int")

	assert.Equal(t, expectedSerialNoBigInt.Cmp(actualSerialNo), 0) //0 == equivalent
}

func createAttributeFromTestFile(t *testing.T, filename string) *yotiprotoattr.Attribute {
	attributeBytes := test.DecodeTestFile(t, filename)

	attributeStruct := &yotiprotoattr.Attribute{}

	err2 := proto.Unmarshal(attributeBytes, attributeStruct)

	assert.Assert(t, is.Nil(err2))

	return attributeStruct
}

func createAnchorSliceFromTestFile(t *testing.T, filename string) []*yotiprotoattr.Anchor {
	anchorBytes := test.DecodeTestFile(t, filename)

	protoAnchor := &yotiprotoattr.Anchor{}
	err2 := proto.Unmarshal(anchorBytes, protoAnchor)
	assert.Assert(t, is.Nil(err2))

	protoAnchors := append([]*yotiprotoattr.Anchor{}, protoAnchor)

	return protoAnchors
}

func TestAnchorParser_Passport(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_passport.txt")

	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A"
	}`)

	a := &yotiprotoattr.Attribute{
		Name:        AttrConstStructuredPostalAddress,
		Value:       structuredAddressBytes,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(a)

	var actualStructuredPostalAddress *attribute.JSONAttribute

	actualStructuredPostalAddress, err := result.StructuredPostalAddress()

	assert.Assert(t, is.Nil(err))

	actualAnchor := actualStructuredPostalAddress.Anchors()[0]

	assert.Equal(t, actualAnchor, actualStructuredPostalAddress.Sources()[0], "Anchors and Sources should be the same when there is only one Source")
	assert.Equal(t, actualAnchor.Type(), anchor.AnchorTypeSource)

	expectedDate := time.Date(2018, time.April, 12, 13, 14, 32, 835537e3, time.UTC)
	actualDate := actualAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := "OCR"
	assert.Equal(t, actualAnchor.SubType(), expectedSubType)

	expectedValue := "PASSPORT"
	assert.Equal(t, actualAnchor.Value()[0], expectedValue)

	actualSerialNo := actualAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "277870515583559162487099305254898397834", actualSerialNo)
}

func TestAnchorParser_DrivingLicense(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_driving_license.txt")

	attribute := &yotiprotoattr.Attribute{
		Name:        AttrConstGender,
		Value:       []byte("value"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attribute)

	genderAttribute := result.Gender()
	resultAnchor := genderAttribute.Anchors()[0]

	assert.Equal(t, resultAnchor, genderAttribute.Sources()[0], "Anchors and Sources should be the same when there is only one Source")
	assert.Equal(t, resultAnchor.Type(), anchor.AnchorTypeSource)

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 3, 923537e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := ""
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)

	expectedValue := "DRIVING_LICENCE"
	assert.Equal(t, resultAnchor.Value()[0], expectedValue)

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "46131813624213904216516051554755262812", actualSerialNo)
}

func TestAnchorParser_UnknownAnchor(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_unknown.txt")

	attr := &yotiprotoattr.Attribute{
		Name:        AttrConstDateOfBirth,
		Value:       []byte("1999-01-01"),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attr)

	DoB, err := result.DateOfBirth()

	assert.Assert(t, is.Nil(err))
	resultAnchor := DoB.Anchors()[0]

	expectedDate := time.Date(2019, time.March, 5, 10, 45, 11, 840037e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := "TEST UNKNOWN SUB TYPE"
	expectedType := anchor.AnchorTypeUnknown
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)
	assert.Equal(t, resultAnchor.Type(), expectedType)
	assert.Equal(t, len(resultAnchor.Value()), 0)
}

func TestAnchorParser_YotiAdmin(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_yoti_admin.txt")

	attr := &yotiprotoattr.Attribute{
		Name:        AttrConstDateOfBirth,
		Value:       []byte("1999-01-01"),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attr)

	DoB, err := result.DateOfBirth()

	assert.Assert(t, is.Nil(err))

	resultAnchor := DoB.Anchors()[0]

	assert.Equal(t, resultAnchor, DoB.Verifiers()[0])

	assert.Equal(t, resultAnchor.Type(), anchor.AnchorTypeVerifier)

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 4, 95238e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := ""
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)

	expectedValue := "YOTI_ADMIN"
	assert.Equal(t, resultAnchor.Value()[0], expectedValue)

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "256616937783084706710155170893983549581", actualSerialNo)
}

func TestAnchors_None(t *testing.T) {
	anchorSlice := []*anchor.Anchor{}

	sources := anchor.GetSources(anchorSlice)
	assert.Equal(t, len(sources), 0, "GetSources should not return anything with empty anchors")

	verifiers := anchor.GetVerifiers(anchorSlice)
	assert.Equal(t, len(verifiers), 0, "GetVerifiers should not return anything with empty anchors")
}
