package create

import "github.com/getyoti/yoti-go-sdk/v3/docscan/constants"

// NotificationConfig represents the configuration properties for notifications within the Doc Scan (IDV) system.
// Notifications can be configured within a Doc Scan (IDV) session to allow your backend to be
// notified of certain events, without having to constantly poll for the state of a session.
type NotificationConfig struct {
	AuthToken string   `json:"auth_token,omitempty"`
	Endpoint  string   `json:"endpoint,omitempty"`
	Topics    []string `json:"topics,omitempty"`
}

// NewNotificationConfigBuilder creates a new NotificationConfigBuilder
func NewNotificationConfigBuilder() *NotificationConfigBuilder {
	return &NotificationConfigBuilder{}
}

// NotificationConfigBuilder builds the NotificationConfig struct
type NotificationConfigBuilder struct {
	authToken string
	endpoint  string
	topics    []string
}

// WithAuthToken sets the authorization token to be included in call-back messages
func (b *NotificationConfigBuilder) WithAuthToken(token string) *NotificationConfigBuilder {
	b.authToken = token
	return b
}

// WithEndpoint sets the endpoint that notifications should be sent to
func (b *NotificationConfigBuilder) WithEndpoint(endpoint string) *NotificationConfigBuilder {
	b.endpoint = endpoint
	return b
}

// WithTopic adds a topic to the slice of topics that trigger notification messages
func (b *NotificationConfigBuilder) WithTopic(topic string) *NotificationConfigBuilder {
	b.topics = append(b.topics, topic)
	return b
}

// ForResourceUpdate Adds "RESOURCE_UPDATE" to the slice of topics that trigger notification messages
func (b *NotificationConfigBuilder) ForResourceUpdate() *NotificationConfigBuilder {
	b.topics = append(b.topics, constants.ResourceUpdate)
	return b
}

// ForTaskCompletion Adds "TASK_COMPLETION" to the slice of topics that trigger notification messages
func (b *NotificationConfigBuilder) ForTaskCompletion() *NotificationConfigBuilder {
	b.topics = append(b.topics, constants.TaskCompletion)
	return b
}

// ForSessionCompletion Adds "SESSION_COMPLETION" to the slice of topics that trigger notification messages
func (b *NotificationConfigBuilder) ForSessionCompletion() *NotificationConfigBuilder {
	b.topics = append(b.topics, constants.SessionCompletion)
	return b
}

// ForCheckCompletion Adds "CHECK_COMPLETION" to the slice of topics that trigger notification messages
func (b *NotificationConfigBuilder) ForCheckCompletion() *NotificationConfigBuilder {
	b.topics = append(b.topics, constants.CheckCompletion)
	return b
}

// Build builds the NotificationConfig struct using the supplied values
func (b *NotificationConfigBuilder) Build() (*NotificationConfig, error) {
	return &NotificationConfig{
		b.authToken,
		b.endpoint,
		b.topics,
	}, nil
}
