package retrieve

import "github.com/getyoti/yoti-go-sdk/v3/util"

// ResourceContainer contains different resources that are part of the Yoti Doc Scan session
type ResourceContainer struct {
	IDDocuments     []IDDocumentResourceResponse `json:"id_documents"`
	LivenessCapture []LivenessResourceResponse   `json:"liveness_capture"`
}

// GetZoomLivenessResources returns a filtered slice of Zoom liveness capture resources
func (rc ResourceContainer) GetZoomLivenessResources() []ZoomLivenessResourceResponse {
	filteredResources := util.Filter(rc.LivenessCapture, func(val interface{}) bool {
		_, isZoomLivenessResource := val.(TextExtractionTaskResponse)
		return isZoomLivenessResource
	})

	return filteredResources.([]ZoomLivenessResourceResponse)
}
