package facecapture

type CreateFaceCaptureResourcePayload struct {
	RequirementID string `json:"requirement_id"`
}

// NewCreateFaceCaptureResourcePayload creates a new payload with the given requirement ID.
func NewCreateFaceCaptureResourcePayload(requirementID string) *CreateFaceCaptureResourcePayload {
	return &CreateFaceCaptureResourcePayload{
		RequirementID: requirementID,
	}
}
