package anchor

import (
	"crypto/x509/pkix"
	"io/ioutil"
	"log"
	"math/big"
	"testing"
	"time"

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

func createAnchorSliceFromTestFile(t *testing.T, filename string) []*yotiprotoattr.Anchor {
	anchorBytes := test.DecodeTestFile(t, filename)

	protoAnchor := &yotiprotoattr.Anchor{}
	err2 := proto.Unmarshal(anchorBytes, protoAnchor)
	assert.Assert(t, is.Nil(err2))

	protoAnchors := append([]*yotiprotoattr.Anchor{}, protoAnchor)

	return protoAnchors
}

func TestAnchorParser_parseExtesion_ShouldErrorForInvalidExtension(t *testing.T) {
	invalidExt := pkix.Extension{
		Id: sourceOID,
	}

	_, _, err := parseExtension(invalidExt)

	assert.Check(t, err != nil)
	assert.Error(t, err, "unable to unmarshal extension: asn1: syntax error: sequence truncated")
}

func TestAnchorParser_Passport(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	anchorSlice := createAnchorSliceFromTestFile(t, "../test/fixtures/test_anchor_passport.txt")

	parsedAnchors := ParseAnchors(anchorSlice)

	actualAnchor := parsedAnchors[0]

	assert.Equal(t, actualAnchor.Type(), AnchorTypeSource)

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
	anchorSlice := createAnchorSliceFromTestFile(t, "../test/fixtures/test_anchor_driving_license.txt")

	parsedAnchors := ParseAnchors(anchorSlice)
	resultAnchor := parsedAnchors[0]

	assert.Equal(t, resultAnchor.Type(), AnchorTypeSource)

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
	anchorSlice := createAnchorSliceFromTestFile(t, "../test/fixtures/test_anchor_unknown.txt")

	resultAnchor := ParseAnchors(anchorSlice)[0]

	expectedDate := time.Date(2019, time.March, 5, 10, 45, 11, 840037e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := "TEST UNKNOWN SUB TYPE"
	expectedType := AnchorTypeUnknown
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)
	assert.Equal(t, resultAnchor.Type(), expectedType)
	assert.Equal(t, len(resultAnchor.Value()), 0)
}

func TestAnchorParser_YotiAdmin(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "../test/fixtures/test_anchor_yoti_admin.txt")

	resultAnchor := ParseAnchors(anchorSlice)[0]

	assert.Equal(t, resultAnchor.Type(), AnchorTypeVerifier)

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
	anchorSlice := []*Anchor{}

	sources := GetSources(anchorSlice)
	assert.Equal(t, len(sources), 0, "GetSources should not return anything with empty anchors")

	verifiers := GetVerifiers(anchorSlice)
	assert.Equal(t, len(verifiers), 0, "GetVerifiers should not return anything with empty anchors")
}
