package share

import (
	"encoding/base64"

	"github.com/getyoti/yoti-go-sdk/v2/credential"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"
)

// ExtraData represents extra pieces information on the receipt.
// Initialize with NewExtraData or DefaultExtraData
type ExtraData struct {
	credentialIssuanceDetails *credential.IssuanceDetails
}

// DefaultExtraData initialises the ExtraData struct
func DefaultExtraData() (extraData *ExtraData) {
	return &ExtraData{
		credentialIssuanceDetails: nil,
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

	var credentialssuanceDetails *credential.IssuanceDetails

	for _, de := range extraDataProto.GetList() {
		if de.Type == yotiprotoshare.DataEntry_THIRD_PARTY_ATTRIBUTE {
			credentialssuanceDetails, err = credential.ParseIssuanceDetails(de.Value)

			if err == nil {
				return &ExtraData{
					credentialIssuanceDetails: credentialssuanceDetails,
				}, nil
			}
		}
	}

	return extraData, nil
}

// CredentialIssuanceDetails represents the details of credential(s) to be issued by a third party. Will be nil if not provided by Yoti.
func (e ExtraData) CredentialIssuanceDetails() (issuanceDetails *credential.IssuanceDetails) {
	return e.credentialIssuanceDetails
}
