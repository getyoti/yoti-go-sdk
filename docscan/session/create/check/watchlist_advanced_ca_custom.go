package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

type RequestedWatchlistAdvancedCACustomAccountCheck struct {
	config RequestedWatchlistAdvancedCACustomAccountConfig
}

// Type is the type of the requested check
func (c RequestedWatchlistAdvancedCACustomAccountCheck) Type() string {
	return constants.WatchlistAdvancedCA
}

// Config is the configuration of the requested check
func (c RequestedWatchlistAdvancedCACustomAccountCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(c.config)
}

// MarshalJSON returns the JSON encoding
func (c RequestedWatchlistAdvancedCACustomAccountCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

func NewRequestedWatchlistAdvancedCACheckCustomAccountBuilder() *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	return &RequestedWatchlistAdvancedCACheckCustomAccountBuilder{}
}

type RequestedWatchlistAdvancedCACustomAccountConfig struct {
	RequestedWatchlistAdvancedCAConfig
	APIKey     string            `json:"api_key,omitempty"`
	Monitoring bool              `json:"monitoring,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
	ClientRef  string            `json:"client_ref,omitempty"`
}

type RequestedWatchlistAdvancedCACheckCustomAccountBuilder struct {
	removeDeceased              bool
	shareURL                    bool
	requestedCASources          RequestedCASources
	requestedCAMatchingStrategy RequestedCAMatchingStrategy
	apiKey                      string
	monitoring                  bool
	tags                        map[string]string
	clientRef                   string
}

// WithAPIKey sets the API key for the Watchlist Advanced CA check (custom account).
func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithAPIKey(apiKey string) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.apiKey = apiKey
	return b
}

// WithMonitoring sets whether monitoring is used for the Watchlist Advanced CA check (custom account).
func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithMonitoring(monitoring bool) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.monitoring = monitoring
	return b
}

// WithTags sets tags used for custom account Watchlist Advanced CA check.
// Please note this will override any previously set tags
func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithTags(tags map[string]string) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.tags = tags
	return b
}

// WithClientRef sets the client reference for the Watchlist Advanced CA check (custom account).
func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithClientRef(clientRef string) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.clientRef = clientRef
	return b
}

func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithRemoveDeceased(removeDeceased bool) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.removeDeceased = removeDeceased
	return b
}

func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithShareURL(shareURL bool) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.shareURL = shareURL
	return b
}

func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithSources(requestedCASources RequestedCASources) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.requestedCASources = requestedCASources
	return b
}

func (b *RequestedWatchlistAdvancedCACheckCustomAccountBuilder) WithMatchingStrategy(requestedCAMatchingStrategy RequestedCAMatchingStrategy) *RequestedWatchlistAdvancedCACheckCustomAccountBuilder {
	b.requestedCAMatchingStrategy = requestedCAMatchingStrategy
	return b
}

func (b RequestedWatchlistAdvancedCACheckCustomAccountBuilder) Build() (RequestedWatchlistAdvancedCACustomAccountCheck, error) {
	config := RequestedWatchlistAdvancedCACustomAccountConfig{
		RequestedWatchlistAdvancedCAConfig: RequestedWatchlistAdvancedCAConfig{
			Type:             constants.WithCustomAccount,
			RemoveDeceased:   b.removeDeceased,
			ShareUrl:         b.shareURL,
			Sources:          b.requestedCASources,
			MatchingStrategy: b.requestedCAMatchingStrategy,
		},
		APIKey:     b.apiKey,
		Monitoring: b.monitoring,
		Tags:       b.tags,
		ClientRef:  b.clientRef,
	}

	return RequestedWatchlistAdvancedCACustomAccountCheck{
		config: config,
	}, nil
}
