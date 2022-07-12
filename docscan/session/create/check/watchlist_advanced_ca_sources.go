package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedCASources is the base type which other CA sources must satisfy
type RequestedCASources interface {
	Type() string
	MarshalJSON() ([]byte, error)
}

type RequestedTypeListSources struct {
	RequestedCASources
	Types []string
}

// Type is the type of the Requested Check
func (c RequestedTypeListSources) Type() string {
	return constants.TypeList
}

// MarshalJSON returns the JSON encoding
func (c RequestedTypeListSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  string   `json:"type"`
		Types []string `json:"types"`
	}{
		Type:  c.Type(),
		Types: c.Types,
	})
}
