package attribute

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"

	"gotest.tools/v3/assert"

	is "gotest.tools/v3/assert/cmp"
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

func TestIssuanceDetails_parseExpiryDate_ShouldParseAllRFC3339Formats(t *testing.T) {
	table := []struct {
		Input    string
		Expected time.Time
	}{
		{
			Input:    "2006-01-02T22:04:05Z",
			Expected: time.Date(2006, 01, 02, 22, 4, 5, 0, time.UTC),
		},
		{
			Input:    "2010-05-20T10:44:25Z",
			Expected: time.Date(2010, 5, 20, 10, 44, 25, 0, time.UTC),
		},
		{
			Input:    "2006-01-02T22:04:05.1Z",
			Expected: time.Date(2006, 1, 2, 22, 4, 5, 100e6, time.UTC),
		},
		{
			Input:    "2012-03-06T04:20:07.5Z",
			Expected: time.Date(2012, 3, 6, 4, 20, 7, 500e6, time.UTC),
		},
		{
			Input:    "2006-01-02T22:04:05.12Z",
			Expected: time.Date(2006, 1, 2, 22, 4, 5, 120e6, time.UTC),
		},
		{
			Input:    "2013-03-04T20:43:55.56Z",
			Expected: time.Date(2013, 3, 4, 20, 43, 55, 560e6, time.UTC),
		},
		{
			Input:    "2006-01-02T22:04:05.123Z",
			Expected: time.Date(2006, 1, 2, 22, 4, 5, 123e6, time.UTC),
		},
		{
			Input:    "2007-04-07T17:34:11.784Z",
			Expected: time.Date(2007, 4, 7, 17, 34, 11, 784e6, time.UTC),
		},
		{
			Input:    "2006-01-02T22:04:05.1234Z",
			Expected: time.Date(2006, 1, 2, 22, 4, 5, 123400e3, time.UTC),
		},
		{
			Input:    "2017-09-14T16:54:30.4784Z",
			Expected: time.Date(2017, 9, 14, 16, 54, 30, 478400e3, time.UTC),
		},
		{
			Input:    "2006-01-02T22:04:05.12345Z",
			Expected: time.Date(2006, 1, 2, 22, 4, 5, 123450e3, time.UTC),
		},
		{
			Input:    "2009-06-07T14:20:30.74622Z",
			Expected: time.Date(2009, 6, 7, 14, 20, 30, 746220e3, time.UTC),
		},
		{
			Input:    "2006-01-02T22:04:05.123456Z",
			Expected: time.Date(2006, 1, 2, 22, 4, 5, 123456e3, time.UTC),
		},
		{
			Input:    "2008-10-25T06:50:55.643562Z",
			Expected: time.Date(2008, 10, 25, 6, 50, 55, 643562e3, time.UTC),
		},
		{
			Input:    "2002-10-02T10:00:00-05:00",
			Expected: time.Date(2002, 10, 2, 10, 0, 0, 0, time.FixedZone("-0500", -5*60*60)),
		},
		{
			Input:    "2002-10-02T10:00:00+11:00",
			Expected: time.Date(2002, 10, 2, 10, 0, 0, 0, time.FixedZone("+1100", 11*60*60)),
		},
		{
			Input:    "1920-03-13T19:50:53.999999Z",
			Expected: time.Date(1920, 3, 13, 19, 50, 53, 999999e3, time.UTC),
		},
		{
			Input:    "1920-03-13T19:50:54.000001Z",
			Expected: time.Date(1920, 3, 13, 19, 50, 54, 1e3, time.UTC),
		},
	}

	for _, row := range table {
		func(input string, expected time.Time) {
			expiryDate, err := parseExpiryDate(input)
			assert.NilError(t, err)
			assert.Equal(t, expiryDate.UTC(), expected.UTC())
		}(row.Input, row.Expected)
	}
}

func TestInvalidProtobufThrowsError(t *testing.T) {
	result, err := ParseIssuanceDetails([]byte("invalid"))

	assert.Assert(t, is.Nil(result))

	assert.Check(t, strings.HasPrefix(err.Error(), "Unable to parse ThirdPartyAttribute value"))
}
