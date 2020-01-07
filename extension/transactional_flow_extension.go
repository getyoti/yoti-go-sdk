package extension

import (
	"encoding/json"
)

const (
	transactionalFlowExtensionTypeConst = "TRANSACTIONAL_FLOW"
)

// TransactionalFlowExtension represents a type of extension in a dynamic share
type TransactionalFlowExtension struct {
	content interface{}
}

// TransactionalFlowExtensionBuilder constructs a TransactionalFlowExtension
type TransactionalFlowExtensionBuilder struct {
	extension TransactionalFlowExtension
}

// WithContent sets the payload data for a TransactionalFlowExtension. The
// content must implement JSON serialization
func (builder *TransactionalFlowExtensionBuilder) WithContent(content interface{}) *TransactionalFlowExtensionBuilder {
	builder.extension.content = content
	return builder
}

// Build constructs a TransactionalFlowExtension
func (builder *TransactionalFlowExtensionBuilder) Build() (TransactionalFlowExtension, error) {
	return builder.extension, nil
}

// MarshalJSON returns the JSON encoding
func (extension TransactionalFlowExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string      `json:"type"`
		Content interface{} `json:"content"`
	}{
		Type:    transactionalFlowExtensionTypeConst,
		Content: extension.content,
	})
}
