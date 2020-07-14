package test

import (
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoshare"
	"github.com/golang/protobuf/proto"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

// CreateThirdPartyAttributeDataEntry creates a data entry of type "THIRD_PARTY_ATTRIBUTE", with the specified IssuingAttribute details.
func CreateThirdPartyAttributeDataEntry(t *testing.T, expiryDate *time.Time, stringDefinitions []string, tokenValue string) yotiprotoshare.DataEntry {
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
