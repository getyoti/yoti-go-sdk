package supported

// DocumentsResponse details the supported countries and associated documents
type DocumentsResponse struct {
	SupportedCountries []*Country `json:"supported_countries"`
}

// Country details the supported documents for a particular country
type Country struct {
	Code               string      `json:"code"`
	SupportedDocuments []*Document `json:"supported_documents"`
}

// Document is the document type that is supported
type Document struct {
	Type            string `json:"type"`
	IsStrictlyLatin bool   `json:"is_strictly_latin"`
}
