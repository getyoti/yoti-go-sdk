package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedTextExtractionTask requests creation of a Text Extraction Task
type RequestedTextExtractionTask struct {
	RequestedTask
	Config RequestedTextExtractionTaskConfig `json:"config"`
}

// NewRequestedTextExtractionTask creates a new Document Authenticity Check
func NewRequestedTextExtractionTask(config RequestedTextExtractionTaskConfig) *RequestedTextExtractionTask {
	return &RequestedTextExtractionTask{
		RequestedTask{
			Type: constants.IDDocumentTextDataExtraction,
		},
		config,
	}
}

// RequestedTextExtractionTaskConfig is the configuration applied when creating a Text Extraction Task
type RequestedTextExtractionTaskConfig struct {
	RequestedTaskConfig
	ManualCheck string `json:"manual_check"`
	ChipData    string `json:"chip_data"`
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
		RequestedTaskConfig{},
		builder.manualCheck,
		builder.chipData,
	}

	return NewRequestedTextExtractionTask(config), nil
}
