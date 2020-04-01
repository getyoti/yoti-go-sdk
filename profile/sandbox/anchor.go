package sandbox

import (
	"encoding/json"
	"time"
)

// Initialises an anchor where the type is "SOURCE",
// which has information about how the anchor was sourced.
func SourceAnchor(subtype string, timestamp time.Time, value string) Anchor {
	return Anchor{
		Type:      "SOURCE",
		Value:     value,
		SubType:   subtype,
		Timestamp: timestamp,
	}
}

// Initialises an anchor where the type is "VERIFIER",
// which has information about how the anchor was verified.
func VerifierAnchor(subtype string, timestamp time.Time, value string) Anchor {
	return Anchor{
		Type:      "VERIFIER",
		Value:     value,
		SubType:   subtype,
		Timestamp: timestamp,
	}
}

// Anchor is the metadata associated with an attribute.
// It describes how an attribute has been provided to Yoti
// (SOURCE Anchor) and how it has been verified (VERIFIER Anchor).
type Anchor struct {
	// Type of the Anchor - most likely either SOURCE or VERIFIER, but it's
	// possible that new Anchor types will be added in future.
	Type string
	// Value identifies the provider that either sourced or verified the attribute value.
	// The range of possible values is not limited. For a SOURCE anchor, expect values like
	// PASSPORT, DRIVING_LICENSE. For a VERIFIER anchor expect valuues like YOTI_ADMIN.
	Value string
	// SubType is an indicator of any specific processing method, or subcategory,
	// pertaining to an artifact. For example, for a passport, this would be
	// either "NFC" or "OCR".
	SubType string
	// Timestamp is the time when the anchor was created, i.e. when it was SOURCED or VERIFIED.
	Timestamp time.Time
}

func (anchor *Anchor) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string `json:"type"`
		Value     string `json:"value"`
		SubType   string `json:"sub_type"`
		Timestamp int64  `json:"timestamp"`
	}{
		Type:      anchor.Type,
		Value:     anchor.Value,
		SubType:   anchor.SubType,
		Timestamp: anchor.Timestamp.UnixNano() / 1e6,
	})
}
