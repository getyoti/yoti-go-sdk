package extension

import (
	"encoding/json"
)

// Extension is a generic type of extension that can be used where a more
// specialised Extension type is not available
type Extension struct {
	extensionType string
	content       interface{}
}

// ExtensionBuilder is used to construct an Extension object
type ExtensionBuilder struct {
	extension Extension
}

// New initializes an ExtensionBuilder
func (builder *ExtensionBuilder) New() *ExtensionBuilder {
	builder.extension.extensionType = ""
	builder.extension.content = nil
	return builder
}

// WithType sets the extension type string
func (builder *ExtensionBuilder) WithType(extensionType string) *ExtensionBuilder {
	builder.extension.extensionType = extensionType
	return builder
}

// WithContent attaches data to the Extension. The content must implement JSON
// serialization
func (builder *ExtensionBuilder) WithContent(content interface{}) *ExtensionBuilder {
	builder.extension.content = content
	return builder
}

// Build constructs the Extension
func (builder *ExtensionBuilder) Build() Extension {
	return builder.extension
}

// MarshalJSON ...
func (extension *Extension) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    string      `json:"type"`
		Content interface{} `json:"content"`
	}{
		Type:    extension.extensionType,
		Content: extension.content,
	})
}
