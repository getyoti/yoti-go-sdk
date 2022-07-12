package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

type RequestedWatchlistAdvancedCAYotiAccountCheck struct {
	config RequestedWatchlistAdvancedCAYotiAccountConfig
}

// Type is the type of the requested check
func (c RequestedWatchlistAdvancedCAYotiAccountCheck) Type() string {
	return constants.WatchlistAdvancedCA
}

// Config is the configuration of the requested check
func (c RequestedWatchlistAdvancedCAYotiAccountCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(c.config)
}

// MarshalJSON returns the JSON encoding
func (c RequestedWatchlistAdvancedCAYotiAccountCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

type RequestedWatchlistAdvancedCACheckYotiAccountBuilder struct {
	removeDeceased              bool
	shareURL                    bool
	requestedCASources          RequestedCASources
	requestedCAMatchingStrategy RequestedCAMatchingStrategy
}

type RequestedWatchlistAdvancedCAYotiAccountConfig struct {
	RequestedWatchlistAdvancedCAConfig
}

// NewRequestedWatchlistAdvancedCACheckYotiAccountBuilder creates a new builder for RequestedWatchlistAdvancedCACheckYotiAccountBuilder
func NewRequestedWatchlistAdvancedCACheckYotiAccountBuilder() *RequestedWatchlistAdvancedCACheckYotiAccountBuilder {
	return &RequestedWatchlistAdvancedCACheckYotiAccountBuilder{}
}

func (b *RequestedWatchlistAdvancedCACheckYotiAccountBuilder) WithRemoveDeceased(removeDeceased bool) *RequestedWatchlistAdvancedCACheckYotiAccountBuilder {
	b.removeDeceased = removeDeceased
	return b
}

func (b *RequestedWatchlistAdvancedCACheckYotiAccountBuilder) WithShareURL(shareURL bool) *RequestedWatchlistAdvancedCACheckYotiAccountBuilder {
	b.shareURL = shareURL
	return b
}

func (b *RequestedWatchlistAdvancedCACheckYotiAccountBuilder) WithSources(requestedCASources RequestedCASources) *RequestedWatchlistAdvancedCACheckYotiAccountBuilder {
	b.requestedCASources = requestedCASources
	return b
}

func (b *RequestedWatchlistAdvancedCACheckYotiAccountBuilder) WithMatchingStrategy(requestedCAMatchingStrategy RequestedCAMatchingStrategy) *RequestedWatchlistAdvancedCACheckYotiAccountBuilder {
	b.requestedCAMatchingStrategy = requestedCAMatchingStrategy
	return b
}

func (b RequestedWatchlistAdvancedCACheckYotiAccountBuilder) Build() (RequestedWatchlistAdvancedCAYotiAccountCheck, error) {
	config := RequestedWatchlistAdvancedCAYotiAccountConfig{
		RequestedWatchlistAdvancedCAConfig{
			Type:             constants.WithYotiAccounts,
			RemoveDeceased:   b.removeDeceased,
			ShareUrl:         b.shareURL,
			Sources:          b.requestedCASources,
			MatchingStrategy: b.requestedCAMatchingStrategy,
		},
	}

	return RequestedWatchlistAdvancedCAYotiAccountCheck{
		config: config,
	}, nil
}
