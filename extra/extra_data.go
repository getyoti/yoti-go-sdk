package extra

import (
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoshare"
	"github.com/golang/protobuf/proto"
)

// Data represents extra pieces information on the receipt.
// Initialize with NewExtraData or DefaultExtraData
type Data struct {
	attributeIssuanceDetails *attribute.IssuanceDetails
}

// DefaultExtraData initialises the ExtraData struct
func DefaultExtraData() (extraData *Data) {
	return &Data{
		attributeIssuanceDetails: nil,
	}
}

// NewExtraData takes a base64 encoded string and parses it into ExtraData
func NewExtraData(extraDataBytes []byte) (*Data, error) {
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

			return &Data{
				attributeIssuanceDetails: attributeIssuanceDetails,
			}, err
		}
	}

	return extraData, nil
}

// AttributeIssuanceDetails represents the details of attribute(s) to be issued by a third party. Will be nil if not provided by Yoti.
func (e Data) AttributeIssuanceDetails() (issuanceDetails *attribute.IssuanceDetails) {
	return e.attributeIssuanceDetails
}
