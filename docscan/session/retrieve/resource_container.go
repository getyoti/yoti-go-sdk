package retrieve

// ResourceContainer contains different resources that are part of the Yoti Doc Scan session
type ResourceContainer struct {
	IDDocuments     []IDDocumentResourceResponse `json:"id_documents"`
	LivenessCapture []LivenessResourceResponse   `json:"liveness_capture"`
}
