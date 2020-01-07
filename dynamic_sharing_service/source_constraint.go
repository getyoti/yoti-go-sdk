package dynamic_sharing_service

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v2/yotierror"
)

// Anchor name constants
const (
	AnchorDrivingLicenceConst = "DRIVING_LICENCE"
	AnchorPassportConst       = "PASSPORT"
	AnchorNationalIDConst     = "NATIONAL_ID"
	AnchorPassCardConst       = "PASS_CARD"
)

// SourceConstraint describes a requirement or preference for a particular set
// of anchors
type SourceConstraint struct {
	anchors        []WantedAnchor
	softPreference bool
}

// SourceConstraintBuilder builds a source constraint
type SourceConstraintBuilder struct {
	sourceConstraint SourceConstraint
	err              error
}

// WithAnchorByValue is a helper method which builds an anchor and adds it to
// the source constraint
func (b *SourceConstraintBuilder) WithAnchorByValue(value, subtype string) *SourceConstraintBuilder {
	anchor, err := (&WantedAnchorBuilder{}).
		WithValue(value).
		WithSubType(subtype).
		Build()
	if err != nil {
		b.err = yotierror.MultiError{This: err, Next: b.err}
	}

	return b.WithAnchor(anchor)
}

// WithAnchor adds an anchor to the preference list
func (b *SourceConstraintBuilder) WithAnchor(anchor WantedAnchor) *SourceConstraintBuilder {
	b.sourceConstraint.anchors = append(b.sourceConstraint.anchors, anchor)
	return b
}

// WithPassport adds a passport anchor
func (b *SourceConstraintBuilder) WithPassport(subtype string) *SourceConstraintBuilder {
	return b.WithAnchorByValue(AnchorPassportConst, subtype)
}

// WithDrivingLicence adds a Driving Licence anchor
func (b *SourceConstraintBuilder) WithDrivingLicence(subtype string) *SourceConstraintBuilder {
	return b.WithAnchorByValue(AnchorDrivingLicenceConst, subtype)
}

// WithNationalID adds a national ID anchor
func (b *SourceConstraintBuilder) WithNationalID(subtype string) *SourceConstraintBuilder {
	return b.WithAnchorByValue(AnchorNationalIDConst, subtype)
}

// WithPasscard adds a passcard anchor
func (b *SourceConstraintBuilder) WithPasscard(subtype string) *SourceConstraintBuilder {
	return b.WithAnchorByValue(AnchorPassCardConst, subtype)
}

// WithSoftPreference sets this constraint as a 'soft requirement' if the
// parameter is true, and a hard requirement if it is false.
func (b *SourceConstraintBuilder) WithSoftPreference(soft bool) *SourceConstraintBuilder {
	b.sourceConstraint.softPreference = soft
	return b
}

// Build builds a SourceConstraint
func (b *SourceConstraintBuilder) Build() (SourceConstraint, error) {
	if b.sourceConstraint.anchors == nil {
		b.sourceConstraint.anchors = make([]WantedAnchor, 0)
	}
	return b.sourceConstraint, b.err
}

func (constraint *SourceConstraint) isConstraint() bool {
	return true
}

// MarshalJSON returns the JSON encoding
func (constraint *SourceConstraint) MarshalJSON() ([]byte, error) {
	type PreferenceList struct {
		Anchors        []WantedAnchor `json:"anchors"`
		SoftPreference bool           `json:"soft_preference"`
	}
	return json.Marshal(&struct {
		Type             string         `json:"type"`
		PreferredSources PreferenceList `json:"preferred_sources"`
	}{
		Type: "SOURCE",
		PreferredSources: PreferenceList{
			Anchors:        constraint.anchors,
			SoftPreference: constraint.softPreference,
		},
	})
}
