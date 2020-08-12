package task

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedTextExtractionTask requests creation of a Text Extraction Task
type RequestedTextExtractionTask struct {
	config RequestedTextExtractionTaskConfig
}

// Type is the type of the Requested Check
func (t RequestedTextExtractionTask) Type() string {
	return constants.IDDocumentTextDataExtraction
}

// Config is the configuration of the Requested Check
func (t RequestedTextExtractionTask) Config() RequestedTaskConfig {
	return t.config
}

func (t RequestedTextExtractionTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string              `json:"type"`
		Config RequestedTaskConfig `json:"config,omitempty"`
	}{
		Type:   t.Type(),
		Config: t.Config(),
	})
}

// NewRequestedTextExtractionTask creates a new Document Authenticity Check
func NewRequestedTextExtractionTask(config RequestedTextExtractionTaskConfig) *RequestedTextExtractionTask {
	return &RequestedTextExtractionTask{config}
}

// RequestedTextExtractionTaskConfig is the configuration applied when creating a Text Extraction Task
type RequestedTextExtractionTaskConfig struct {
	ManualCheck string `json:"manual_check,omitempty"`
	ChipData    string `json:"chip_data,omitempty"`
}

// NewRequestedTextExtractionTaskBuilder creates a new RequestedTextExtractionTaskBuilder
func NewRequestedTextExtractionTaskBuilder() *RequestedTextExtractionTaskBuilder {
	return &RequestedTextExtractionTaskBuilder{}
}

// RequestedTextExtractionTaskBuilder builds a RequestedTextExtractionTask
type RequestedTextExtractionTaskBuilder struct {
	manualCheck string
	chipData    string
}

// WithManualCheckAlways sets the value of manual check to "ALWAYS"
func (builder *RequestedTextExtractionTaskBuilder) WithManualCheckAlways() *RequestedTextExtractionTaskBuilder {
	builder.manualCheck = constants.Always
	return builder
}

// WithManualCheckFallback sets the value of manual check to "FALLBACK"
func (builder *RequestedTextExtractionTaskBuilder) WithManualCheckFallback() *RequestedTextExtractionTaskBuilder {
	builder.manualCheck = constants.Fallback
	return builder
}

// WithManualCheckNever sets the value of manual check to "NEVER"
func (builder *RequestedTextExtractionTaskBuilder) WithManualCheckNever() *RequestedTextExtractionTaskBuilder {
	builder.manualCheck = constants.Never
	return builder
}

// WithChipDataDesired sets the value of chip data to "DESIRED"
func (builder *RequestedTextExtractionTaskBuilder) WithChipDataDesired() *RequestedTextExtractionTaskBuilder {
	builder.chipData = chipDataDesired
	return builder
}

// WithChipDataIgnore sets the value of chip data to "IGNORE"
func (builder *RequestedTextExtractionTaskBuilder) WithChipDataIgnore() *RequestedTextExtractionTaskBuilder {
	builder.chipData = chipDataIgnore
	return builder
}

// Build builds the RequestedTextExtractionTask
func (builder *RequestedTextExtractionTaskBuilder) Build() (*RequestedTextExtractionTask, error) {
	config := RequestedTextExtractionTaskConfig{
		builder.manualCheck,
		builder.chipData,
	}

	return NewRequestedTextExtractionTask(config), nil
}
