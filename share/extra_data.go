package share

import (
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoshare"
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
func NewExtraData(extraDataBytes []byte) (*ExtraData, error) {
	var err error
	var extraData = DefaultExtraData()

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
