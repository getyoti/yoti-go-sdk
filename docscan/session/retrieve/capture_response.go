package retrieve

import (
	"encoding/json"
	"fmt"
)

// RequiredResourceResponse defines interface for required resources.
type RequiredResourceResponse interface {
	GetType() string
	String() string
}

// BaseRequiredResource contains common fields for all required resources.
type BaseRequiredResource struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	State string `json:"state"`
}

func (b *BaseRequiredResource) GetType() string {
	return b.Type
}

func (b *BaseRequiredResource) String() string {
	return fmt.Sprintf("Type: %s, ID: %s, State: %s", b.Type, b.ID, b.State)
}

type RequiredIdDocumentResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredIdDocumentResourceResponse) String() string {
	return fmt.Sprintf("ID Document Resource - %s", r.BaseRequiredResource.String())
}

type RequiredSupplementaryDocumentResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredSupplementaryDocumentResourceResponse) String() string {
	return fmt.Sprintf("Supplementary Document Resource - %s", r.BaseRequiredResource.String())
}

type RequiredZoomLivenessResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredZoomLivenessResourceResponse) String() string {
	return fmt.Sprintf("Zoom Liveness Resource - %s", r.BaseRequiredResource.String())
}

type RequiredFaceCaptureResourceResponse struct {
	BaseRequiredResource
}

func (r *RequiredFaceCaptureResourceResponse) String() string {
	return fmt.Sprintf("Face Capture Resource - %s", r.BaseRequiredResource.String())
}

type UnknownRequiredResourceResponse struct {
	BaseRequiredResource
}

func (r *UnknownRequiredResourceResponse) String() string {
	return fmt.Sprintf("Unknown Resource Type - %s", r.BaseRequiredResource.String())
}

// CaptureResponse holds the biometric consent and polymorphic required resources.
type CaptureResponse struct {
	BiometricConsent  string                     `json:"biometric_consent"`
	RequiredResources []RequiredResourceResponse `json:"-"`
}

// captureResponseAlias used internally for unmarshaling raw JSON resources.
type captureResponseAlias struct {
	BiometricConsent  string            `json:"biometric_consent"`
	RequiredResources []json.RawMessage `json:"required_resources"`
}

// UnmarshalJSON unmarshals CaptureResponse with polymorphic required resources.
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

// filterByType filters resources by the given type T.
func filterByType[T RequiredResourceResponse](resources []RequiredResourceResponse) []T {
	var filtered []T
	for _, r := range resources {
		if typed, ok := r.(T); ok {
			filtered = append(filtered, typed)
		}
	}
	return filtered
}

// GetDocumentResourceRequirements returns ID and supplementary document resources.
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

// GetIdDocumentResourceRequirements returns ID document resources.
func (c *CaptureResponse) GetIdDocumentResourceRequirements() []*RequiredIdDocumentResourceResponse {
	return filterByType[*RequiredIdDocumentResourceResponse](c.RequiredResources)
}

// GetSupplementaryResourceRequirements returns supplementary document resources.
func (c *CaptureResponse) GetSupplementaryResourceRequirements() []*RequiredSupplementaryDocumentResourceResponse {
	return filterByType[*RequiredSupplementaryDocumentResourceResponse](c.RequiredResources)
}

// GetZoomLivenessResourceRequirements returns Zoom liveness resources.
func (c *CaptureResponse) GetZoomLivenessResourceRequirements() []*RequiredZoomLivenessResourceResponse {
	return filterByType[*RequiredZoomLivenessResourceResponse](c.RequiredResources)
}

// GetFaceCaptureResourceRequirements returns face capture resources.
func (c *CaptureResponse) GetFaceCaptureResourceRequirements() []*RequiredFaceCaptureResourceResponse {
	return filterByType[*RequiredFaceCaptureResourceResponse](c.RequiredResources)
}

func (c *CaptureResponse) GetLivenessResourceRequirements() []*RequiredZoomLivenessResourceResponse {
	return filterByType[*RequiredZoomLivenessResourceResponse](c.RequiredResources)
}
