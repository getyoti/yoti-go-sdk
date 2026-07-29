package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ResourceContainer contains different resources that are part of the Yoti IDV session
type ResourceContainer struct {
	IDDocuments             []*IDDocumentResourceResponse            `json:"id_documents"`
	SupplementaryDocuments  []*SupplementaryDocumentResourceResponse `json:"supplementary_documents"`
	ShareCodes              []*ShareCodeResourceResponse             `json:"share_codes"`
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

// filterForCheck returns a new ResourceContainer holding only the resources used by the given check.
func (r *ResourceContainer) filterForCheck(check *CheckResponse) *ResourceContainer {
	filtered := &ResourceContainer{}
	if r == nil || check == nil {
		return filtered
	}

	ids := make(map[string]struct{}, len(check.ResourcesUsed))
	for _, id := range check.ResourcesUsed {
		ids[id] = struct{}{}
	}

	filtered.IDDocuments = filterResourcesByIDs(r.IDDocuments, ids)
	filtered.SupplementaryDocuments = filterResourcesByIDs(r.SupplementaryDocuments, ids)
	filtered.ShareCodes = filterResourcesByIDs(r.ShareCodes, ids)
	filtered.zoomLivenessResources = filterResourcesByIDs(r.zoomLivenessResources, ids)
	filtered.staticLivenessResources = filterResourcesByIDs(r.staticLivenessResources, ids)

	// LivenessCapture and RawLivenessCapture are index-parallel (see UnmarshalJSON),
	// so filter them together to keep the returned container internally consistent.
	for i, liveness := range r.LivenessCapture {
		if _, ok := ids[liveness.GetID()]; !ok {
			continue
		}
		filtered.LivenessCapture = append(filtered.LivenessCapture, liveness)
		if i < len(r.RawLivenessCapture) {
			filtered.RawLivenessCapture = append(filtered.RawLivenessCapture, r.RawLivenessCapture[i])
		}
	}

	return filtered
}

func filterResourcesByIDs[T interface{ GetID() string }](resources []T, ids map[string]struct{}) []T {
	var filtered []T
	for _, resource := range resources {
		if _, ok := ids[resource.GetID()]; ok {
			filtered = append(filtered, resource)
		}
	}
	return filtered
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
