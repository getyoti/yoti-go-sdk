package digitalidentity

import (
	"encoding/json"
)

// ShareSessionNotification specifies the session notification configuration.
type ShareSessionNotification struct {
	url       string
	method    string
	verifyTLS bool
	headers   map[string][]string
}

// ShareSessionNotificationBuilder builds Share Session Notification
type ShareSessionNotificationBuilder struct {
	shareSessionNotification ShareSessionNotification
}

// WithUrl setsUrl  to Share Session Notification
func (b *ShareSessionNotificationBuilder) WithUrl(url string) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.url = url
	return b
}

// WithMethod set method to Share Session Notification
func (b *ShareSessionNotificationBuilder) WithMethod(method string) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.method = method
	return b
}

// WithVerifyTLS sets whether TLS should be verified for notifications.
func (b *ShareSessionNotificationBuilder) WithVerifyTLS(verify bool) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.verifyTLS = verify
	return b
}

// WithHeaders set headers to Share Session Notification
func (b *ShareSessionNotificationBuilder) WithHeaders(headers map[string][]string) *ShareSessionNotificationBuilder {
	b.shareSessionNotification.headers = headers
	return b
}

// Build constructs the Share Session Notification Builder
func (b *ShareSessionNotificationBuilder) Build() (ShareSessionNotification, error) {
	return b.shareSessionNotification, nil
}

// MarshalJSON returns the JSON encoding
func (a *ShareSessionNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Url       string              `json:"url"`
		Method    string              `json:"method"`
		VerifyTls bool                `json:"verifyTls"`
		Headers   map[string][]string `json:"headers"`
	}{
		Url:       a.url,
		Method:    a.method,
		VerifyTls: a.verifyTLS,
		Headers:   a.headers,
	})
}
