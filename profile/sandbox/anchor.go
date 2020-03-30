package sandbox

import (
	"encoding/json"
	"time"
)

func SourceAnchor(subtype string, timestamp time.Time, value string) Anchor {
	return Anchor{
		Type:      "SOURCE",
		Value:     value,
		SubType:   subtype,
		Timestamp: timestamp,
	}
}

func VerifierAnchor(subtype string, timestamp time.Time, value string) Anchor {
	return Anchor{
		Type:      "VERIFIER",
		Value:     value,
		SubType:   subtype,
		Timestamp: timestamp,
	}
}

// Anchor describes an anchor on a Sandbox Attribute
type Anchor struct {
	Type      string
	Value     string
	SubType   string
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
