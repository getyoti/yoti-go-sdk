package share

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestAttributeIssuanceDetailsShouldReturnNilWhenNoDataEntries(t *testing.T) {
	extraData := DefaultExtraData()

	issuanceDetails := extraData.AttributeIssuanceDetails()

	assert.Assert(t, is.Nil(issuanceDetails))
}

func TestShouldReturnFirstMatchingThirdPartyAttribute(t *testing.T) {
	dataEntries := make([]*yotiprotoshare.DataEntry, 0)

	expiryDate := time.Now().UTC().AddDate(0, 0, 1)
	var tokenValue1 string = "tokenValue1"

	thirdPartyAttributeDataEntry1 := test.CreateThirdPartyAttributeDataEntry(t, &expiryDate, []string{"attributeName1"}, tokenValue1)
	thirdPartyAttributeDataEntry2 := test.CreateThirdPartyAttributeDataEntry(t, &expiryDate, []string{"attributeName2"}, "tokenValue2")

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry1, &thirdPartyAttributeDataEntry2)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	parsedExtraData, err := parseProtoExtraData(t, protoExtraData)
	assert.Assert(t, is.Nil(err))

	result := parsedExtraData.AttributeIssuanceDetails()

	var tokenBytes []byte = []byte(tokenValue1)
	var base64EncodedToken string = base64.StdEncoding.EncodeToString(tokenBytes)

	assert.Equal(t, result.Token(), base64EncodedToken)
	assert.Equal(t, result.Attributes()[0].Name(), "attributeName1")
	assert.Equal(t,
		result.ExpiryDate().Format("2006-01-02T15:04:05.000Z"),
		expiryDate.Format("2006-01-02T15:04:05.000Z"))
}

func TestShouldParseMultipleIssuingAttributes(t *testing.T) {
	var base64ExtraData string = test.GetTestFileAsString(t, "testextradata.txt")

	extraData, err := NewExtraData(base64ExtraData)
	assert.Assert(t, is.Nil(err))

	result := extraData.AttributeIssuanceDetails()

	assert.Equal(t, result.Token(), "c29tZUlzc3VhbmNlVG9rZW4=")
	assert.Equal(t,
		result.ExpiryDate().Format("2006-01-02T15:04:05.000Z"),
		time.Date(2019, time.October, 15, 22, 04, 05, 123000000, time.UTC).Format("2006-01-02T15:04:05.000Z"))
	assert.Equal(t, result.Attributes()[0].Name(), "com.thirdparty.id")
	assert.Equal(t, result.Attributes()[1].Name(), "com.thirdparty.other_id")
}

func TestShouldHandleNoExpiryDate(t *testing.T) {
	var protoDefinitions []*yotiprotoshare.Definition

	protoDefinitions = append(protoDefinitions, &yotiprotoshare.Definition{Name: "attribute.name"})

	thirdPartyAttribute := &yotiprotoshare.ThirdPartyAttribute{
		IssuanceToken: []byte("tokenValue"),
		IssuingAttributes: &yotiprotoshare.IssuingAttributes{
			ExpiryDate:  "",
			Definitions: protoDefinitions,
		},
	}

	marshalledThirdPartyAttribute, err := proto.Marshal(thirdPartyAttribute)
	assert.Assert(t, is.Nil(err))

	result, _ := processThirdPartyAttribute(t, marshalledThirdPartyAttribute)

	assert.Assert(t, is.Nil(result.ExpiryDate()))
}

func TestShouldHandleNoIssuingAttributes(t *testing.T) {
	var tokenValueBytes []byte = []byte("token")
	thirdPartyAttribute := &yotiprotoshare.ThirdPartyAttribute{
		IssuanceToken:     tokenValueBytes,
		IssuingAttributes: &yotiprotoshare.IssuingAttributes{},
	}

	marshalledThirdPartyAttribute, err := proto.Marshal(thirdPartyAttribute)
	assert.Assert(t, is.Nil(err))

	result, err := processThirdPartyAttribute(t, marshalledThirdPartyAttribute)

	assert.Assert(t, is.Nil(err))
	assert.Equal(t, base64.StdEncoding.EncodeToString(tokenValueBytes), result.Token())
}

func TestShouldHandleNoIssuingAttributeDefinitions(t *testing.T) {
	var tokenValueBytes []byte = []byte("token")

	thirdPartyAttribute := &yotiprotoshare.ThirdPartyAttribute{
		IssuanceToken: tokenValueBytes,
		IssuingAttributes: &yotiprotoshare.IssuingAttributes{
			ExpiryDate:  time.Now().UTC().AddDate(0, 0, 1).Format("2006-01-02T15:04:05.000Z"),
			Definitions: []*yotiprotoshare.Definition{},
		},
	}

	marshalledThirdPartyAttribute, err := proto.Marshal(thirdPartyAttribute)
	assert.Assert(t, is.Nil(err))

	result, err := processThirdPartyAttribute(t, marshalledThirdPartyAttribute)

	assert.Assert(t, is.Nil(err))
	assert.Equal(t, base64.StdEncoding.EncodeToString(tokenValueBytes), result.Token())
}

func processThirdPartyAttribute(t *testing.T, marshalledThirdPartyAttribute []byte) (*attribute.IssuanceDetails, error) {
	dataEntries := make([]*yotiprotoshare.DataEntry, 0)

	thirdPartyAttributeDataEntry := yotiprotoshare.DataEntry{
		Type:  yotiprotoshare.DataEntry_THIRD_PARTY_ATTRIBUTE,
		Value: marshalledThirdPartyAttribute,
	}

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	parsedExtraData, err := parseProtoExtraData(t, protoExtraData)

	return parsedExtraData.AttributeIssuanceDetails(), err
}

func parseProtoExtraData(t *testing.T, protoExtraData *yotiprotoshare.ExtraData) (*ExtraData, error) {
	extraDataMarshalled, err := proto.Marshal(protoExtraData)
	assert.Assert(t, is.Nil(err))

	extraDataBase64 := base64.StdEncoding.EncodeToString(extraDataMarshalled)
	return NewExtraData(extraDataBase64)
}
