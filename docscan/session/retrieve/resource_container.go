package retrieve

// ResourceContainer TODO: add
type ResourceContainer struct {
	IDDocuments     []IDDocumentResourceResponse `json:"id_documents"`
	LivenessCapture []LivenessResourceResponse   `json:"liveness_capture"`
}

//TODO: add ZoomLivenessResources filter
