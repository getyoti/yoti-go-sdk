package objective

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ProofOfAddressObjective requests creation of a proof of address objective
type ProofOfAddressObjective struct {
}

// Type is the objective type
func (o *ProofOfAddressObjective) Type() string {
	return constants.ProofOfAddress
}

// MarshalJSON returns the JSON encoding
func (o *ProofOfAddressObjective) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
	}{
		Type: o.Type(),
	})
}

// NewProofOfAddressObjectiveBuilder creates a new ProofOfAddressObjectiveBuilder
func NewProofOfAddressObjectiveBuilder() *ProofOfAddressObjectiveBuilder {
	return &ProofOfAddressObjectiveBuilder{}
}

// ProofOfAddressObjectiveBuilder builds a ProofOfAddress
type ProofOfAddressObjectiveBuilder struct {
}

// Build builds the ProofOfAddressObjective
func (builder *ProofOfAddressObjectiveBuilder) Build() (*ProofOfAddressObjective, error) {
	return &ProofOfAddressObjective{}, nil
}
