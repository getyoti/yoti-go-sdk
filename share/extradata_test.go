package share

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestShouldReturnNilForNoDataEntries(t *testing.T) {
	extraData := DefaultExtraData()

	issuanceDetails := extraData.CredentialIssuanceDetails()

	assert.Assert(t, is.Nil(issuanceDetails))
}

func TestShouldReturnFirstMatchingThirdPartyAttribute(t *testing.T) {
	dataEntries := make([]*yotiprotoshare.DataEntry, 0)

	expiryDate := time.Now().UTC().AddDate(0, 0, 1)
	thirdPartyAttributeDataEntry1 := createThirdPartyAttributeDataEntry(t, expiryDate, []string{"attributeName1"}, "tokenValue1")
	thirdPartyAttributeDataEntry2 := createThirdPartyAttributeDataEntry(t, expiryDate, []string{"attributeName2"}, "tokenValue2")

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry1, &thirdPartyAttributeDataEntry2)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	parsedExtraData := parseProtoExtraData(t, protoExtraData)
	result := parsedExtraData.CredentialIssuanceDetails()

	assert.Equal(t, result.Token(), "tokenValue1")
	assert.Equal(t, result.IssuingAttributes()[0], "attributeName1")
	assert.Equal(t,
		result.ExpiryDate().Format("2006-01-02T15:04:05.000Z"),
		expiryDate.Format("2006-01-02T15:04:05.000Z"))
}

func TestShouldParseMultipleIssuingAttributes(t *testing.T) {
	var base64ExtraData string = test.GetTestFileAsString(t, "testextradata.txt")

	extraData, err := NewExtraData(base64ExtraData)
	assert.Assert(t, is.Nil(err))

	result := extraData.CredentialIssuanceDetails()

	assert.Equal(t, result.Token(), "someIssuanceToken")
	assert.Equal(t,
		result.ExpiryDate().Format("2006-01-02T15:04:05.000Z"),
		time.Date(2019, time.October, 15, 22, 04, 05, 123000000, time.UTC).Format("2006-01-02T15:04:05.000Z"))
	assert.Equal(t, result.IssuingAttributes()[0], "com.thirdparty.id")
	assert.Equal(t, result.IssuingAttributes()[1], "com.thirdparty.other_id")
}

func parseProtoExtraData(t *testing.T, protoExtraData *yotiprotoshare.ExtraData) (parsedExtraData *ExtraData) {
	extraDataMarshalled, err := proto.Marshal(protoExtraData)
	assert.Assert(t, is.Nil(err))

	extraDataBase64 := base64.StdEncoding.EncodeToString(extraDataMarshalled)
	parsedExtraData, err = NewExtraData(extraDataBase64)

	assert.Assert(t, is.Nil(err))
	return parsedExtraData
}

func createThirdPartyAttributeDataEntry(t *testing.T, expiryDate time.Time, stringDefinitions []string, tokenValue string) yotiprotoshare.DataEntry {
	var protoDefinitions []*yotiprotoshare.Definition

	for _, definition := range stringDefinitions {
		protoDefinition := &yotiprotoshare.Definition{
			Name: definition,
		}

		protoDefinitions = append(protoDefinitions, protoDefinition)
	}

	thirdPartyAttribute := &yotiprotoshare.ThirdPartyAttribute{
		IssuanceToken: []byte(tokenValue),
		IssuingAttributes: &yotiprotoshare.IssuingAttributes{
			ExpiryDate:  expiryDate.Format("2006-01-02T15:04:05.000Z"),
			Definitions: protoDefinitions,
		},
	}

	marshalledThirdPartyAttribute, err := proto.Marshal(thirdPartyAttribute)

	assert.Assert(t, is.Nil(err))

	return yotiprotoshare.DataEntry{
		Type:  yotiprotoshare.DataEntry_THIRD_PARTY_ATTRIBUTE,
		Value: marshalledThirdPartyAttribute,
	}
}
