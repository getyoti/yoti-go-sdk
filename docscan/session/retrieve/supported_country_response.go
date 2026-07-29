package retrieve

// SupportedCountryResponse represents a country and its supported documents for a required resource
type SupportedCountryResponse struct {
	Code               string                       `json:"code"`
	SupportedDocuments []*SupportedDocumentResponse `json:"supported_documents"`
}

// SupportedDocumentResponse represents a supported document type and the providers that can fulfil it
type SupportedDocumentResponse struct {
	Type      string   `json:"type"`
	Providers []string `json:"providers"`
}
