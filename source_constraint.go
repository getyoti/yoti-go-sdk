package yoti

import (
	"encoding/json"
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
}

// New initialises a SourceConstraintBuilder
func (b *SourceConstraintBuilder) New() *SourceConstraintBuilder {
	b.sourceConstraint.anchors = make([]WantedAnchor, 0)
	b.sourceConstraint.softPreference = false
	return b
}

// WithAnchor adds an anchor to the preference list
func (b *SourceConstraintBuilder) WithAnchor(anchor WantedAnchor) *SourceConstraintBuilder {
	b.sourceConstraint.anchors = append(b.sourceConstraint.anchors, anchor)
	return b
}

// WithSoftPreference sets this constraint as a 'soft requirement' if the
// parameter is true, and a hard requirement if it is false.
func (b *SourceConstraintBuilder) WithSoftPreference(soft bool) *SourceConstraintBuilder {
	b.sourceConstraint.softPreference = soft
	return b
}

// Build builds a SourceConstraint
func (b *SourceConstraintBuilder) Build() SourceConstraint {
	return b.sourceConstraint
}

func (constraint *SourceConstraint) isConstraint() bool {
	return true
}

// MarshalJSON ...
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
