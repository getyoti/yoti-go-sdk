package task

const (
	valueProgress       string = "PROGRESS"
	valueMustTryAgain   string = "MUST_TRY_AGAIN"
	valueShouldTryAgain string = "SHOULD_TRY_AGAIN"
)

// TextDataExtractionRecommendation represents a text data extraction reason
type TextDataExtractionRecommendation struct {
	Value  string                    `json:"value"`
	Reason *TextDataExtractionReason `json:"reason,omitempty"`
}

// TextDataExtractionRecommendationBuilder builds a TextDataExtractionRecommendation
type TextDataExtractionRecommendationBuilder struct {
	value  string
	reason *TextDataExtractionReason
}

// NewTextDataExtractionRecommendationBuilder creates a new TextDataExtractionRecommendationBuilder
func NewTextDataExtractionRecommendationBuilder() *TextDataExtractionRecommendationBuilder {
	return &TextDataExtractionRecommendationBuilder{}
}

// ForProgress sets the recommendation value to progress
func (b *TextDataExtractionRecommendationBuilder) ForProgress() *TextDataExtractionRecommendationBuilder {
	b.value = valueProgress
	return b
}

// ForMustTryAgain sets the recommendation value to must try again
func (b *TextDataExtractionRecommendationBuilder) ForMustTryAgain() *TextDataExtractionRecommendationBuilder {
	b.value = valueMustTryAgain
	return b
}

// ForShouldTryAgain sets the recommendation value to should try again
func (b *TextDataExtractionRecommendationBuilder) ForShouldTryAgain() *TextDataExtractionRecommendationBuilder {
	b.value = valueShouldTryAgain
	return b
}

// WithReason sets the recommendation reason
func (b *TextDataExtractionRecommendationBuilder) WithReason(reason *TextDataExtractionReason) *TextDataExtractionRecommendationBuilder {
	b.reason = reason
	return b
}

// Build creates a new TextDataExtractionRecommendation
func (b *TextDataExtractionRecommendationBuilder) Build() (*TextDataExtractionRecommendation, error) {
	return &TextDataExtractionRecommendation{
		Value:  b.value,
		Reason: b.reason,
	}, nil
}
