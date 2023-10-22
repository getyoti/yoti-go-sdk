package digitalidentity

import (
	"encoding/json"
)

// ShareSessionBuilder builds a session
type CreateShareQrCodeBuilder struct {
	createShareQrCode CreateShareQrCodeResult
	err               error
}

// ShareSession represents a sharesession
type CreateShareQrCodeResult struct {
	id  string
	uri string
}

// Build constructs the ShareSession
func (builder *CreateShareQrCodeBuilder) Build() (CreateShareQrCodeResult, error) {
	return builder.createShareQrCode, builder.err
}

// MarshalJSON returns the JSON encoding
func (createShareQrCode CreateShareQrCodeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id  string `json:"id"`
		Uri string `json:"uri"`
	}{
		Id:  createShareQrCode.id,
		Uri: createShareQrCode.uri,
	})
}
