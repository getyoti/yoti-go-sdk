package retrieve

import (
	"encoding/json"
)

// Define RequiredResourceResponse interface (polymorphic)
type RequiredResourceResponse interface {
	GetType() string
}

// Base struct for common fields
type BaseRequiredResource struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	State string `json:"state"`
}

func (b *BaseRequiredResource) GetType() string {
	return b.Type
}

// Concrete resource structs embedding BaseRequiredResource
type RequiredIdDocumentResourceResponse struct {
	BaseRequiredResource
	// add other fields if needed
}

type RequiredSupplementaryDocumentResourceResponse struct {
	BaseRequiredResource
	// add other fields if needed
}

type RequiredZoomLivenessResourceResponse struct {
	BaseRequiredResource
	// add other fields if needed
}

type RequiredFaceCaptureResourceResponse struct {
	BaseRequiredResource
	// add other fields if needed
}

type UnknownRequiredResourceResponse struct {
	BaseRequiredResource
}

// CaptureResponse struct with polymorphic RequiredResources slice
type CaptureResponse struct {
	BiometricConsent  string                     `json:"biometric_consent"`
	RequiredResources []RequiredResourceResponse `json:"-"`
}

// Raw struct for unmarshaling RequiredResources as raw JSON
type captureResponseAlias struct {
	BiometricConsent  string            `json:"biometric_consent"`
	RequiredResources []json.RawMessage `json:"required_resources"`
}

func (c *CaptureResponse) UnmarshalJSON(data []byte) error {
	aux := captureResponseAlias{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	c.BiometricConsent = aux.BiometricConsent
	c.RequiredResources = make([]RequiredResourceResponse, 0, len(aux.RequiredResources))

	for _, raw := range aux.RequiredResources {
		// Peek type field
		var base BaseRequiredResource
		if err := json.Unmarshal(raw, &base); err != nil {
			return err
		}

		var resource RequiredResourceResponse

		switch base.Type {
		case "ID_DOCUMENT":
			var r RequiredIdDocumentResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return err
			}
			resource = &r
		case "SUPPLEMENTARY_DOCUMENT":
			var r RequiredSupplementaryDocumentResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return err
			}
			resource = &r
		case "LIVENESS":
			var r RequiredZoomLivenessResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return err
			}
			resource = &r
		case "FACE_CAPTURE":
			var r RequiredFaceCaptureResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return err
			}
			resource = &r
		default:
			var r UnknownRequiredResourceResponse
			if err := json.Unmarshal(raw, &r); err != nil {
				return err
			}
			resource = &r
		}

		c.RequiredResources = append(c.RequiredResources, resource)
	}

	return nil
}

// Helper generic filter function
func filterByType[T RequiredResourceResponse](resources []RequiredResourceResponse) []T {
	var filtered []T
	for _, r := range resources {
		if typed, ok := r.(T); ok {
			filtered = append(filtered, typed)
		}
	}
	return filtered
}

// Document resource requirements (ID + supplementary)
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

func (c *CaptureResponse) GetLivenessResourceRequirements() []*RequiredZoomLivenessResourceResponse {
	return filterByType[*RequiredZoomLivenessResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetZoomLivenessResourceRequirements() []*RequiredZoomLivenessResourceResponse {
	return filterByType[*RequiredZoomLivenessResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetFaceCaptureResourceRequirements() []*RequiredFaceCaptureResourceResponse {
	return filterByType[*RequiredFaceCaptureResourceResponse](c.RequiredResources)
}
