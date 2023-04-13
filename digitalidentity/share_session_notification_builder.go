package digitalidentity

import (
	"encoding/json"
)

// WantedAnchor specifies a preferred anchor for a user's details
type ShareSessionNotification struct {
	url       string
	method    string
	verifyTls bool
	Headers   map[string][]string
}

// ShareSessionNotificationBuilder
type ShareSessionNotificationBuilder struct {
	shareSessionNotification ShareSessionNotification
}

// WithUrl setsUrl
func (b *ShareSessionNotificationBuilder) WithUrl(url string) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.url = url
	return b
}

// WithMethod set method
func (b *ShareSessionNotificationBuilder) WithMethod(method string) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.method = method
	return b
}

// WithVerifyTls set bool value verifyTls
func (b *ShareSessionNotificationBuilder) WithVerifyTls(verifyTls bool) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.verifyTls = verifyTls
	return b
}

// WithHeaders set headers
func (b *ShareSessionNotificationBuilder) WithHeaders(headers map[string][]string) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.Headers = headers
	return b
}

// Build
func (b *ShareSessionNotificationBuilder) Build() (ShareSessionNotification, error) {
	return b.shareSessionNotification, nil
}

// MarshalJSON ...
func (a *ShareSessionNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Url       string              `json:"url"`
		Method    string              `json:"method"`
		VerifyTls bool                `json:"verifyTls"`
		Headers   map[string][]string `json:"headers"`
	}{
		Url:       a.url,
		Method:    a.method,
		VerifyTls: a.verifyTls,
		Headers:   a.Headers,
	})
}
