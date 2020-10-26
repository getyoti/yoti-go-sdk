package retrieve

// SupplementaryDocumentTextExtractionTaskResponse represents a Supplementary Text Extraction task response
type SupplementaryDocumentTextExtractionTaskResponse struct {
	*TaskResponse
}

// GeneratedTextDataChecks filters the checks, returning only text data checks
func (t *SupplementaryDocumentTextExtractionTaskResponse) GeneratedTextDataChecks() []*GeneratedSupplementaryTextDataCheckResponse {
	return t.generatedSupplementaryTextDataChecks
}
