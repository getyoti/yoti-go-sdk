package task

const (
	valueQuality   string = "QUALITY"
	valueUserError string = "USER_ERROR"
)

// TextDataExtractionReason represents a text data extraction reason
type TextDataExtractionReason struct {
	Value  string `json:"value"`
	Detail string `json:"detail,omitempty"`
}

// TextDataExtractionReasonBuilder builds a TextDataExtractionReason
type TextDataExtractionReasonBuilder struct {
	value  string
	detail string
}

// NewTextDataExtractionReasonBuilder creates a new TextDataExtractionReasonBuilder
func NewTextDataExtractionReasonBuilder() *TextDataExtractionReasonBuilder {
	return &TextDataExtractionReasonBuilder{}
}

// ForQuality sets the reason to quality
func (b *TextDataExtractionReasonBuilder) ForQuality() *TextDataExtractionReasonBuilder {
	b.value = valueQuality
	return b
}

// ForUserError sets the reason to user error
func (b *TextDataExtractionReasonBuilder) ForUserError() *TextDataExtractionReasonBuilder {
	b.value = valueUserError
	return b
}

// WithDetail sets the reason detail
func (b *TextDataExtractionReasonBuilder) WithDetail(detail string) *TextDataExtractionReasonBuilder {
	b.detail = detail
	return b
}

// Build creates a new TextDataExtractionReason
func (b *TextDataExtractionReasonBuilder) Build() (*TextDataExtractionReason, error) {
	return &TextDataExtractionReason{
		Detail: b.detail,
		Value:  b.value,
	}, nil
}
