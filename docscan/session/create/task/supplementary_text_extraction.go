package task

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedSupplementaryDocTextExtractionTask requests creation of a Text Extraction Task
type RequestedSupplementaryDocTextExtractionTask struct {
	config RequestedSupplementaryDocTextExtractionTaskConfig
}

// Type is the type of the Requested Task
func (t *RequestedSupplementaryDocTextExtractionTask) Type() string {
	return constants.SupplementaryDocumentTextDataExtraction
}

// Config is the configuration of the Requested Task
func (t *RequestedSupplementaryDocTextExtractionTask) Config() RequestedTaskConfig {
	return t.config
}

// MarshalJSON marshals the RequestedSupplementaryDocTextExtractionTask to JSON
func (t *RequestedSupplementaryDocTextExtractionTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string              `json:"type"`
		Config RequestedTaskConfig `json:"config,omitempty"`
	}{
		Type:   t.Type(),
		Config: t.Config(),
	})
}

// NewRequestedSupplementaryDocTextExtractionTask creates a new supplementary document text extraction task
func NewRequestedSupplementaryDocTextExtractionTask(config RequestedSupplementaryDocTextExtractionTaskConfig) *RequestedSupplementaryDocTextExtractionTask {
	return &RequestedSupplementaryDocTextExtractionTask{config}
}

// RequestedSupplementaryDocTextExtractionTaskConfig is the configuration applied when creating a Text Extraction Task
type RequestedSupplementaryDocTextExtractionTaskConfig struct {
	ManualCheck string `json:"manual_check,omitempty"`
}

// NewRequestedSupplementaryDocTextExtractionTaskBuilder creates a new RequestedSupplementaryDocTextExtractionTaskBuilder
func NewRequestedSupplementaryDocTextExtractionTaskBuilder() *RequestedSupplementaryDocTextExtractionTaskBuilder {
	return &RequestedSupplementaryDocTextExtractionTaskBuilder{}
}

// RequestedSupplementaryDocTextExtractionTaskBuilder builds a RequestedSupplementaryDocTextExtractionTask
type RequestedSupplementaryDocTextExtractionTaskBuilder struct {
	manualCheck string
}

// WithManualCheckAlways sets the value of manual check to "ALWAYS"
func (builder *RequestedSupplementaryDocTextExtractionTaskBuilder) WithManualCheckAlways() *RequestedSupplementaryDocTextExtractionTaskBuilder {
	builder.manualCheck = constants.Always
	return builder
}

// WithManualCheckFallback sets the value of manual check to "FALLBACK"
func (builder *RequestedSupplementaryDocTextExtractionTaskBuilder) WithManualCheckFallback() *RequestedSupplementaryDocTextExtractionTaskBuilder {
	builder.manualCheck = constants.Fallback
	return builder
}

// WithManualCheckNever sets the value of manual check to "NEVER"
func (builder *RequestedSupplementaryDocTextExtractionTaskBuilder) WithManualCheckNever() *RequestedSupplementaryDocTextExtractionTaskBuilder {
	builder.manualCheck = constants.Never
	return builder
}

// Build builds the RequestedSupplementaryDocTextExtractionTask
func (builder *RequestedSupplementaryDocTextExtractionTaskBuilder) Build() (*RequestedSupplementaryDocTextExtractionTask, error) {
	config := RequestedSupplementaryDocTextExtractionTaskConfig{
		builder.manualCheck,
	}

	return NewRequestedSupplementaryDocTextExtractionTask(config), nil
}
