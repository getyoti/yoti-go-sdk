package task

import (
	"encoding/base64"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentTextDataExtractionTask represents a document text data extraction task
type DocumentTextDataExtractionTask struct {
	*documentTask
	Result documentTextDataExtractionTaskResult `json:"result"`
	Config documentTextDataExtractionConfig     `json:"config,omitempty"`
}

type documentTextDataExtractionConfig struct {
	ManualCheck string `json:"manual_check,omitempty"`
}

// DocumentTextDataExtractionTaskBuilder builds a DocumentTextDataExtractionTask
type DocumentTextDataExtractionTaskBuilder struct {
	documentTaskBuilder
	documentFields  map[string]interface{}
	documentIDPhoto *documentIDPhoto
	detectedCountry string
	recommendation  *TextDataExtractionRecommendation
	manualCheck     string
}

type documentTextDataExtractionTaskResult struct {
	DocumentFields  map[string]interface{}            `json:"document_fields,omitempty"`
	DocumentIDPhoto *documentIDPhoto                  `json:"document_id_photo,omitempty"`
	DetectedCountry string                            `json:"detected_country,omitempty"`
	Recommendation  *TextDataExtractionRecommendation `json:"recommendation,omitempty"`
}

type documentIDPhoto struct {
	ContentType string `json:"content_type"`
	Data        string `json:"data"`
}

// NewDocumentTextDataExtractionTaskBuilder creates a new DocumentTextDataExtractionTaskBuilder
func NewDocumentTextDataExtractionTaskBuilder() *DocumentTextDataExtractionTaskBuilder {
	return &DocumentTextDataExtractionTaskBuilder{}
}

// WithDocumentFilter adds a document filter to the task
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentFilter(filter *filter.DocumentFilter) *DocumentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the task
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentField(key string, value interface{}) *DocumentTextDataExtractionTaskBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]interface{})
	}
	b.documentFields[key] = value
	return b
}

// WithDocumentFields sets document fields
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentFields(documentFields map[string]interface{}) *DocumentTextDataExtractionTaskBuilder {
	b.documentFields = documentFields
	return b
}

// WithDocumentIDPhoto sets the document ID photo
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentIDPhoto(contentType string, data []byte) *DocumentTextDataExtractionTaskBuilder {
	b.documentIDPhoto = &documentIDPhoto{
		ContentType: contentType,
		Data:        base64.StdEncoding.EncodeToString(data),
	}
	return b
}

// WithDetectedCountry sets the detected country
func (b *DocumentTextDataExtractionTaskBuilder) WithDetectedCountry(detectedCountry string) *DocumentTextDataExtractionTaskBuilder {
	b.detectedCountry = detectedCountry
	return b
}

// WithRecommendation sets the recommendation
func (b *DocumentTextDataExtractionTaskBuilder) WithRecommendation(recommendation *TextDataExtractionRecommendation) *DocumentTextDataExtractionTaskBuilder {
	b.recommendation = recommendation
	return b
}

// WithManualCheckNever sets the manual check config to never
func (b *DocumentTextDataExtractionTaskBuilder) WithManualCheckNever() {
	b.manualCheck = constants.Never
}

// WithManualCheckAlways sets the manual check config to always
func (b *DocumentTextDataExtractionTaskBuilder) WithManualCheckAlways() {
	b.manualCheck = constants.Always
}

// WithManualCheckFallback sets the manual check config to fallback
func (b *DocumentTextDataExtractionTaskBuilder) WithManualCheckFallback() {
	b.manualCheck = constants.Fallback
}

// Build creates a new DocumentTextDataExtractionTask
func (b *DocumentTextDataExtractionTaskBuilder) Build() (*DocumentTextDataExtractionTask, error) {
	return &DocumentTextDataExtractionTask{
		documentTask: b.documentTaskBuilder.build(),
		Result: documentTextDataExtractionTaskResult{
			DocumentFields:  b.documentFields,
			DocumentIDPhoto: b.documentIDPhoto,
			DetectedCountry: b.detectedCountry,
			Recommendation:  b.recommendation,
		},
		Config: documentTextDataExtractionConfig{
			ManualCheck: b.manualCheck,
		},
	}, nil
}
