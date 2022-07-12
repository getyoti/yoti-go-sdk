package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedCAMatchingStrategy is the base type which other CA matching strategies must satisfy
type RequestedCAMatchingStrategy interface {
	Type() string
}

func NewRequestedFuzzyMatchingStrategy() *RequestedFuzzyMatchingStrategy {
	return &RequestedFuzzyMatchingStrategy{}
}

type RequestedFuzzyMatchingStrategy struct {
	RequestedCAMatchingStrategy
	Fuzziness float64
}

// MarshalJSON returns the JSON encoding
func (c RequestedFuzzyMatchingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string  `json:"type"`
		Fuzziness float64 `json:"fuzziness"`
	}{
		Type:      constants.Fuzzy,
		Fuzziness: c.Fuzziness,
	})
}

type RequestedExactMatchingStrategy struct {
	RequestedCAMatchingStrategy
	ExactMatch bool
}

// MarshalJSON returns the JSON encoding
func (c RequestedExactMatchingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type       string `json:"type"`
		ExactMatch bool   `json:"exact_match"`
	}{
		Type:       constants.Exact,
		ExactMatch: c.ExactMatch,
	})
}
