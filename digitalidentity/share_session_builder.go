package digitalidentity

import (
	"encoding/json"
)

// ShareSessionBuilder builds a session
type ShareSessionBuilder struct {
	shareSessionRequest ShareSessionRequest
	err                 error
}

// ShareSession represents a sharesession
type ShareSessionRequest struct {
	policy                   *Policy
	redirectUri              string
	extensions               []interface{}
	subject                  *json.RawMessage
	shareSessionNotification ShareSessionNotification
}

// WithPolicy attaches a Policy to the ShareSession
func (builder *ShareSessionBuilder) WithPolicy(policy Policy) *ShareSessionBuilder {
	builder.shareSessionRequest.policy = &policy
	return builder
}

// WithExtension adds an extension to the ShareSession
func (builder *ShareSessionBuilder) WithExtension(extension interface{}) *ShareSessionBuilder {
	builder.shareSessionRequest.extensions = append(builder.shareSessionRequest.extensions, extension)
	return builder
}

// WithNotification sets the callback URL
func (builder *ShareSessionBuilder) WithNotification(notification ShareSessionNotification) *ShareSessionBuilder {
	builder.shareSessionRequest.shareSessionNotification = notification
	return builder
}

// WithRedirectUri sets redirectUri to the ShareSession
func (builder *ShareSessionBuilder) WithRedirectUri(redirectUri string) *ShareSessionBuilder {
	builder.shareSessionRequest.redirectUri = redirectUri
	return builder
}

// WithSubject adds a subject to the ShareSession. Must be valid JSON.
func (builder *ShareSessionBuilder) WithSubject(subject json.RawMessage) *ShareSessionBuilder {
	builder.shareSessionRequest.subject = &subject
	return builder
}

// Build constructs the ShareSession
func (builder *ShareSessionBuilder) Build() (ShareSessionRequest, error) {
	if builder.shareSessionRequest.extensions == nil {
		builder.shareSessionRequest.extensions = make([]interface{}, 0)
	}
	if builder.shareSessionRequest.policy == nil {
		policy, err := (&PolicyBuilder{}).Build()
		if err != nil {
			return builder.shareSessionRequest, err
		}
		builder.shareSessionRequest.policy = &policy
	}
	return builder.shareSessionRequest, builder.err
}

// MarshalJSON returns the JSON encoding
func (shareSesssion ShareSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Policy       Policy                   `json:"policy"`
		Extensions   []interface{}            `json:"extensions"`
		RedirectUri  string                   `json:"redirectUri"`
		Subject      *json.RawMessage         `json:"subject,omitempty"`
		Notification ShareSessionNotification `json:"notification"`
	}{
		Policy:       *shareSesssion.policy,
		Extensions:   shareSesssion.extensions,
		RedirectUri:  shareSesssion.redirectUri,
		Subject:      shareSesssion.subject,
		Notification: shareSesssion.shareSessionNotification,
	})
}
