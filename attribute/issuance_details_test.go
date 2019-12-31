package attribute

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"

	"gotest.tools/assert"

	is "gotest.tools/assert/cmp"
)

func TestShouldParseThirdPartyAttributeCorrectly(t *testing.T) {
	var thirdPartyAttributeBytes []byte = test.GetTestFileBytes(t, "../test/fixtures/test_third_party_issuance_details.txt")
	issuanceDetails, err := ParseIssuanceDetails(thirdPartyAttributeBytes)

	assert.Assert(t, is.Nil(err))
	assert.Equal(t, issuanceDetails.Attributes()[0].Name(), "com.thirdparty.id")
	assert.Equal(t, issuanceDetails.Token(), "c29tZUlzc3VhbmNlVG9rZW4=")
	assert.Equal(t,
		issuanceDetails.ExpiryDate().Format("2006-01-02T15:04:05.000Z"),
		"2019-10-15T22:04:05.123Z")
}

func TestShouldLogWarningIfErrorInParsingExpiryDate(t *testing.T) {
	var tokenValue string = "41548a175dfaw"
	thirdPartyAttribute := &yotiprotoshare.ThirdPartyAttribute{
		IssuanceToken: []byte(tokenValue),
		IssuingAttributes: &yotiprotoshare.IssuingAttributes{
			ExpiryDate: "2006-13-02T15:04:05.000Z",
		},
	}

	marshalled, err := proto.Marshal(thirdPartyAttribute)

	assert.Assert(t, is.Nil(err))

	var tokenBytes []byte = []byte(tokenValue)
	var expectedBase64Token string = base64.StdEncoding.EncodeToString(tokenBytes)

	result, err := ParseIssuanceDetails(marshalled)
	assert.Equal(t, expectedBase64Token, result.Token())
	assert.Assert(t, is.Nil(result.ExpiryDate()))
	assert.Equal(t, "parsing time \"2006-13-02T15:04:05.000Z\": month out of range", err.Error())
}

func TestInvalidProtobufThrowsError(t *testing.T) {
	result, err := ParseIssuanceDetails([]byte("invalid"))

	assert.Assert(t, is.Nil(result))

	assert.Check(t, strings.HasPrefix(err.Error(), "Unable to parse ThirdPartyAttribute value"))
}
