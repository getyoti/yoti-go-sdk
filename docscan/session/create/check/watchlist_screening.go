package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedWatchlistScreeningCheck requests creation of a Watchlist Screening Check.
// To request a RequestedWatchlistScreeningCheck you must request task.RequestedTextExtractionTask as a minimum
type RequestedWatchlistScreeningCheck struct {
	config RequestedWatchlistScreeningCheckConfig
}

// Type is the type of the requested check
func (c *RequestedWatchlistScreeningCheck) Type() string {
	return constants.WatchlistScreening
}

// Config is the configuration of the requested check
func (c *RequestedWatchlistScreeningCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(c.config)
}

// MarshalJSON returns the JSON encoding
func (c *RequestedWatchlistScreeningCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedWatchlistScreeningCheckConfig is the configuration applied when creating
// a watchlist screening check
type RequestedWatchlistScreeningCheckConfig struct {
	Categories []string `json:"categories"`
}

// RequestedWatchlistScreeningCheckBuilder builds a RequestedWatchlistScreeningCheck
type RequestedWatchlistScreeningCheckBuilder struct {
	config     RequestedWatchlistScreeningCheckConfig
	categories []string
}

// NewRequestedWatchlistScreeningCheckBuilder creates a new builder for RequestedWatchlistScreeningCheck
func NewRequestedWatchlistScreeningCheckBuilder() *RequestedWatchlistScreeningCheckBuilder {
	return &RequestedWatchlistScreeningCheckBuilder{}
}

// WithCategory adds a category to the list of categories used for watchlist screening
func (b *RequestedWatchlistScreeningCheckBuilder) WithCategory(category string) *RequestedWatchlistScreeningCheckBuilder {
	b.categories = append(b.categories, category)
	return b
}

// WithAdverseMediaCategory adds ADVERSE_MEDIA to the list of categories used for watchlist screening
func (b *RequestedWatchlistScreeningCheckBuilder) WithAdverseMediaCategory() *RequestedWatchlistScreeningCheckBuilder {
	return b.WithCategory(constants.AdverseMedia)
}

// WithSanctionsCategory adds SANCTIONS to the list of categories used for watchlist screening
func (b *RequestedWatchlistScreeningCheckBuilder) WithSanctionsCategory() *RequestedWatchlistScreeningCheckBuilder {
	return b.WithCategory(constants.AdverseMedia)
}

// Build builds the RequestedWatchlistScreeningCheck
func (b *RequestedWatchlistScreeningCheckBuilder) Build() (*RequestedWatchlistScreeningCheck, error) {
	return &RequestedWatchlistScreeningCheck{
		config: b.config,
	}, nil
}
