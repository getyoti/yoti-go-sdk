package digitalidentity

import (
	"encoding/json"
)

// ShareSessionRequestBuilder builds a session
type ShareSessionRequestBuilder struct {
	shareSessionRequest ShareSessionRequest
	err                 error
}

// ShareSession represents a sharesession
type ShareSessionRequest struct {
	policy                   Policy
	extensions               []interface{}
	subject                  *json.RawMessage
	shareSessionNotification *ShareSessionNotification
	redirectUri              string
}

// WithPolicy attaches a Policy to the ShareSession
func (builder *ShareSessionRequestBuilder) WithPolicy(policy Policy) *ShareSessionRequestBuilder {
	builder.shareSessionRequest.policy = policy
	return builder
}

// WithExtension adds an extension to the ShareSession
func (builder *ShareSessionRequestBuilder) WithExtension(extension interface{}) *ShareSessionRequestBuilder {
	builder.shareSessionRequest.extensions = append(builder.shareSessionRequest.extensions, extension)
	return builder
}

// WithNotification sets the callback URL
func (builder *ShareSessionRequestBuilder) WithNotification(notification *ShareSessionNotification) *ShareSessionRequestBuilder {
	builder.shareSessionRequest.shareSessionNotification = notification
	return builder
}

// WithRedirectUri sets redirectUri to the ShareSession
func (builder *ShareSessionRequestBuilder) WithRedirectUri(redirectUri string) *ShareSessionRequestBuilder {
	builder.shareSessionRequest.redirectUri = redirectUri
	return builder
}

// WithSubject adds a subject to the ShareSession. Must be valid JSON.
func (builder *ShareSessionRequestBuilder) WithSubject(subject json.RawMessage) *ShareSessionRequestBuilder {
	builder.shareSessionRequest.subject = &subject
	return builder
}

// Build constructs the ShareSession
func (builder *ShareSessionRequestBuilder) Build() (ShareSessionRequest, error) {
	if builder.shareSessionRequest.extensions == nil {
		builder.shareSessionRequest.extensions = make([]interface{}, 0)
	}
	return builder.shareSessionRequest, builder.err
}

// MarshalJSON returns the JSON encoding
func (shareSesssion ShareSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Policy       Policy                    `json:"policy"`
		Extensions   []interface{}             `json:"extensions"`
		RedirectUri  string                    `json:"redirectUri"`
		Subject      *json.RawMessage          `json:"subject,omitempty"`
		Notification *ShareSessionNotification `json:"notification,omitempty"`
	}{
		Policy:       shareSesssion.policy,
		Extensions:   shareSesssion.extensions,
		RedirectUri:  shareSesssion.redirectUri,
		Subject:      shareSesssion.subject,
		Notification: shareSesssion.shareSessionNotification,
	})
}
