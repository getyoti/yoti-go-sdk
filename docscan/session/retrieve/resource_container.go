package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ResourceContainer contains different resources that are part of the Yoti Doc Scan session
type ResourceContainer struct {
	IDDocuments           []*IDDocumentResourceResponse `json:"id_documents"`
	LivenessCapture       []*LivenessResourceResponse   `json:"liveness_capture"`
	zoomLivenessResources []*ZoomLivenessResourceResponse
}

// ZoomLivenessResources  filters the liveness resources, returning only the "Zoom" liveness resources
func (r ResourceContainer) ZoomLivenessResources() []*ZoomLivenessResourceResponse {
	return r.zoomLivenessResources
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (r *ResourceContainer) UnmarshalJSON(data []byte) error {
	type result ResourceContainer // declared as "type" to prevent recursive unmarshalling
	if err := json.Unmarshal(data, (*result)(r)); err != nil {
		return err
	}

	for _, resource := range r.LivenessCapture {
		if resource.LivenessType == constants.Zoom {
			r.zoomLivenessResources = append(r.zoomLivenessResources, &ZoomLivenessResourceResponse{LivenessResourceResponse: resource})
		}
	}

	return nil
}
