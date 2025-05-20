package retrieve

import (
	"encoding/json"
	"fmt"
)

type CaptureResponse struct {
	BiometricConsent  string                     `json:"biometric_consent"`
	RequiredResources []RequiredResourceResponse `json:"-"`
}

type captureResponseAlias struct {
	BiometricConsent  string            `json:"biometric_consent"`
	RequiredResources []json.RawMessage `json:"required_resources"`
}

func (c *CaptureResponse) UnmarshalJSON(data []byte) error {
	aux := captureResponseAlias{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("failed to unmarshal CaptureResponse: %w", err)
	}

	c.BiometricConsent = aux.BiometricConsent
	c.RequiredResources = make([]RequiredResourceResponse, 0, len(aux.RequiredResources))

	for _, raw := range aux.RequiredResources {
		var base BaseRequiredResource
		if err := json.Unmarshal(raw, &base); err != nil {
			return fmt.Errorf("failed to unmarshal base resource: %w", err)
		}

		var resource RequiredResourceResponse

		switch base.Type {
		case "ID_DOCUMENT":
			var r RequiredIdDocumentResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return fmt.Errorf("failed to unmarshal ID_DOCUMENT resource: %w", err)
			}
			resource = &r
		case "SUPPLEMENTARY_DOCUMENT":
			var r RequiredSupplementaryDocumentResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return fmt.Errorf("failed to unmarshal SUPPLEMENTARY_DOCUMENT resource: %w", err)
			}
			resource = &r
		case "LIVENESS":
			var r RequiredZoomLivenessResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return fmt.Errorf("failed to unmarshal LIVENESS resource: %w", err)
			}
			resource = &r
		case "FACE_CAPTURE":
			var r RequiredFaceCaptureResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return fmt.Errorf("failed to unmarshal FACE_CAPTURE resource: %w", err)
			}
			resource = &r
		default:
			var r UnknownRequiredResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return fmt.Errorf("failed to unmarshal unknown resource type: %w", err)
			}
			resource = &r
		}

		c.RequiredResources = append(c.RequiredResources, resource)
	}

	return nil
}

// Helper generic filter function for typed filtering
func filterByType[T RequiredResourceResponse](resources []RequiredResourceResponse) []T {
	var filtered []T
	for _, r := range resources {
		if typed, ok := r.(T); ok {
			filtered = append(filtered, typed)
		}
	}
	return filtered
}

func (c *CaptureResponse) GetDocumentResourceRequirements() []RequiredResourceResponse {
	var docs []RequiredResourceResponse
	for _, r := range c.RequiredResources {
		switch r.(type) {
		case *RequiredIdDocumentResourceResponse, *RequiredSupplementaryDocumentResourceResponse:
			docs = append(docs, r)
		}
	}
	return docs
}

func (c *CaptureResponse) GetIdDocumentResourceRequirements() []*RequiredIdDocumentResourceResponse {
	return filterByType[*RequiredIdDocumentResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetSupplementaryResourceRequirements() []*RequiredSupplementaryDocumentResourceResponse {
	return filterByType[*RequiredSupplementaryDocumentResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetZoomLivenessResourceRequirements() []*RequiredZoomLivenessResourceResponse {
	return filterByType[*RequiredZoomLivenessResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetFaceCaptureResourceRequirements() []*RequiredFaceCaptureResourceResponse {
	return filterByType[*RequiredFaceCaptureResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetLivenessResourceRequirements() []*RequiredZoomLivenessResourceResponse {
	return filterByType[*RequiredZoomLivenessResourceResponse](c.RequiredResources)
}
