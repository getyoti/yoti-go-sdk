package dynamic_sharing_service

import (
	"encoding/json"
)

// WantedAnchor specifies a preferred anchor for a user's details
type WantedAnchor struct {
	name    string
	subType string
}

// WantedAnchorBuilder describes a desired anchor for user profile data
type WantedAnchorBuilder struct {
	wantedAnchor WantedAnchor
}

// WithValue sets the anchor's name
func (b *WantedAnchorBuilder) WithValue(name string) *WantedAnchorBuilder {
	b.wantedAnchor.name = name
	return b
}

// WithSubType sets the anchors subtype
func (b *WantedAnchorBuilder) WithSubType(subType string) *WantedAnchorBuilder {
	b.wantedAnchor.subType = subType
	return b
}

// Build constructs the anchor from the builder's specification
func (b *WantedAnchorBuilder) Build() (WantedAnchor, error) {
	return b.wantedAnchor, nil
}

// MarshalJSON ...
func (a *WantedAnchor) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name    string `json:"name"`
		SubType string `json:"sub_type"`
	}{
		Name:    a.name,
		SubType: a.subType,
	})
}
