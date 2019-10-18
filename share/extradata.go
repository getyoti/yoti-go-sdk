package share

import (
	"encoding/base64"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"
)

// ExtraData represents extra pieces information on the receipt.
// Initialize with NewExtraData or DefaultExtraData
type ExtraData struct {
	attributeIssuanceDetails *attribute.IssuanceDetails
}

// DefaultExtraData initialises the ExtraData struct
func DefaultExtraData() (extraData *ExtraData) {
	return &ExtraData{
		attributeIssuanceDetails: nil,
	}
}

// NewExtraData takes a base64 encoded string and parses it into ExtraData
func NewExtraData(extraDataEncodedString string) (*ExtraData, error) {
	var extraDataBytes []byte
	var err error
	var extraData *ExtraData = DefaultExtraData()

	if extraDataBytes, err = base64.StdEncoding.DecodeString(extraDataEncodedString); err != nil {
		return extraData, err
	}

	extraDataProto := &yotiprotoshare.ExtraData{}
	if err = proto.Unmarshal(extraDataBytes, extraDataProto); err != nil {
		return extraData, err
	}

	var attributeIssuanceDetails *attribute.IssuanceDetails

	for _, de := range extraDataProto.GetList() {
		if de.Type == yotiprotoshare.DataEntry_THIRD_PARTY_ATTRIBUTE {
			attributeIssuanceDetails, err = attribute.ParseIssuanceDetails(de.Value)

			return &ExtraData{
				attributeIssuanceDetails: attributeIssuanceDetails,
			}, err
		}
	}

	return extraData, nil
}

// AttributeIssuanceDetails represents the details of attribute(s) to be issued by a third party. Will be nil if not provided by Yoti.
func (e ExtraData) AttributeIssuanceDetails() (issuanceDetails *attribute.IssuanceDetails) {
	return e.attributeIssuanceDetails
}
