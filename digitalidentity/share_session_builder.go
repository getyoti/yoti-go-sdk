package digitalidentity

import (
	"encoding/json"
)

// ScenarioBuilder builds a dynamic scenario
type ShareSessionBuilder struct {
	shareSession ShareSession
	err          error
}

// Scenario represents a dynamic scenario
type ShareSession struct {
	policy                   *Policy
	redirectUri              string
	extensions               []interface{}
	subject                  *json.RawMessage
	shareSessionNotification *ShareSessionNotification
}

// WithPolicy attaches a DynamicPolicy to the DynamicScenario
func (builder *ShareSessionBuilder) WithPolicy(policy Policy) *ShareSessionBuilder {
	builder.shareSession.policy = &policy
	return builder
}

// WithExtension adds an extension to the scenario
func (builder *ShareSessionBuilder) WithExtension(extension interface{}) *ShareSessionBuilder {
	builder.shareSession.extensions = append(builder.shareSession.extensions, extension)
	return builder
}

// WithCallbackEndpoint sets the callback URL
func (builder *ShareSessionBuilder) WithNotification(notification ShareSessionNotification) *ShareSessionBuilder {
	builder.shareSession.shareSessionNotification = &notification
	return builder
}

// WithRedirectUri
func (builder *ShareSessionBuilder) WithRedirectUri(redirectUri string) *ShareSessionBuilder {
	builder.shareSession.redirectUri = redirectUri
	return builder
}

// WithSubject adds a subject to the ShareSession. Must be valid JSON.
func (builder *ShareSessionBuilder) WithSubject(subject json.RawMessage) *ShareSessionBuilder {
	builder.shareSession.subject = &subject
	return builder
}

// Build constructs the DynamicScenario
func (builder *ShareSessionBuilder) Build() (ShareSession, error) {
	if builder.shareSession.extensions == nil {
		builder.shareSession.extensions = make([]interface{}, 0)
	}
	if builder.shareSession.policy == nil {
		policy, err := (&PolicyBuilder{}).Build()
		if err != nil {
			return builder.shareSession, err
		}
		builder.shareSession.policy = &policy
	}
	return builder.shareSession, builder.err
}

// MarshalJSON returns the JSON encoding
func (shareSesssion ShareSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Policy       Policy                   `json:"policy"`
		Extensions   []interface{}            `json:"extensions"`
		RedirectUri  string                   `json:"redirectUri"`
		Subject      *json.RawMessage         `json:"subject,omitempty"`
		Notification ShareSessionNotification `json:"notification,omitempty"`
	}{
		Policy:       *shareSesssion.policy,
		Extensions:   shareSesssion.extensions,
		RedirectUri:  shareSesssion.redirectUri,
		Subject:      shareSesssion.subject,
		Notification: *shareSesssion.shareSessionNotification,
	})
}
