package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ResourceContainer contains different resources that are part of the Yoti IDV session
type ResourceContainer struct {
	IDDocuments             []*IDDocumentResourceResponse            `json:"id_documents"`
	SupplementaryDocuments  []*SupplementaryDocumentResourceResponse `json:"supplementary_documents"`
	LivenessCapture         []*LivenessResourceResponse
	RawLivenessCapture      []json.RawMessage `json:"liveness_capture"`
	zoomLivenessResources   []*ZoomLivenessResourceResponse
	staticLivenessResources []*StaticLivenessResourceResponse
}

// ZoomLivenessResources  filters the liveness resources, returning only the "Zoom" liveness resources
func (r *ResourceContainer) ZoomLivenessResources() []*ZoomLivenessResourceResponse {
	return r.zoomLivenessResources
}

// ZoomLivenessResources  filters the liveness resources, returning only the "Zoom" liveness resources
func (r *ResourceContainer) StaticLivenessResources() []*StaticLivenessResourceResponse {
	return r.staticLivenessResources
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (r *ResourceContainer) UnmarshalJSON(data []byte) error {
	type resourceContainer ResourceContainer
	err := json.Unmarshal(data, (*resourceContainer)(r))
	if err != nil {
		return err
	}

	for _, raw := range r.RawLivenessCapture {
		var v LivenessResourceResponse
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return err
		}

		switch v.LivenessType {
		case constants.Zoom:
			var zoom ZoomLivenessResourceResponse
			err = json.Unmarshal(raw, &zoom)
			if err != nil {
				return err
			}
			r.zoomLivenessResources = append(r.zoomLivenessResources, &zoom)
		case constants.Static:
			var static StaticLivenessResourceResponse
			err = json.Unmarshal(raw, &static)
			if err != nil {
				return err
			}
			r.staticLivenessResources = append(r.staticLivenessResources, &static)
		default:
			err = json.Unmarshal(raw, &LivenessResourceResponse{})
			if err != nil {
				return err
			}
		}

		r.LivenessCapture = append(r.LivenessCapture, &v)
	}

	return nil
}
