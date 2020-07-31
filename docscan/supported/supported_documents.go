package supported

type DocumentsResponse struct {
	SupportedCountries []Country `json:"supported_countries"`
}

type Country struct {
	Code               string     `json:"code"`
	SupportedDocuments []Document `json:"supported_documents"`
}

type Document struct {
	Type string `json:"type"`
}
