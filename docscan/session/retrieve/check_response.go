package retrieve

import (
	"time"
)

// CheckResponse represents the attributes of a check, for any given session
type CheckResponse struct {
	ID             string            `json:"id"`
	Type           string            `json:"type"`
	State          string            `json:"state"`
	ResourcesUsed  []string          `json:"resources_used"`
	GeneratedMedia []*GeneratedMedia `json:"generated_media"`
	Report         *ReportResponse   `json:"report"`
	Created        *time.Time        `json:"created"`
	LastUpdated    *time.Time        `json:"last_updated"`
}

// AuthenticityCheckResponse represents a Document Authenticity check for a given session
type AuthenticityCheckResponse struct {
	*CheckResponse
}

// FaceMatchCheckResponse represents a FaceMatch Check for a given session
type FaceMatchCheckResponse struct {
	*CheckResponse
}

// LivenessCheckResponse represents a Liveness Check for a given session
type LivenessCheckResponse struct {
	*CheckResponse
}

// TextDataCheckResponse represents a Text Data check for a given session
type TextDataCheckResponse struct {
	*CheckResponse
}

// IDDocumentComparisonCheckResponse represents a identity document comparison check for a given session
type IDDocumentComparisonCheckResponse struct {
	*CheckResponse
}

// SupplementaryDocumentTextDataCheckResponse represents a supplementary document text data check for a given session
type SupplementaryDocumentTextDataCheckResponse struct {
	*CheckResponse
}

// ThirdPartyCheckResponse represents a check with an external credit reference agency
type ThirdPartyIdentityCheckResponse struct {
	*CheckResponse
}
