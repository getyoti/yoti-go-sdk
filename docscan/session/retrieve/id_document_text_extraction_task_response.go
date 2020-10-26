package retrieve

// TextExtractionTaskResponse represents a Text Extraction task response
type TextExtractionTaskResponse struct {
	*TaskResponse
}

// GeneratedTextDataChecks filters the checks, returning only text data checks
func (t *TextExtractionTaskResponse) GeneratedTextDataChecks() []*GeneratedTextDataCheckResponse {
	return t.generatedTextDataChecks
}
