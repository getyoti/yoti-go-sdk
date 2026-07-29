package retrieve

import "time"

// DigitalIDShareResponse represents a digital ID share performed in a session
type DigitalIDShareResponse struct {
	ID             string                       `json:"id"`
	DocumentType   string                       `json:"document_type"`
	IssuingCountry string                       `json:"issuing_country"`
	Provider       string                       `json:"provider"`
	CreatedAt      *time.Time                   `json:"created_at"`
	LastUpdated    *time.Time                   `json:"last_updated"`
	ResourceID     string                       `json:"resource_id"`
	Error          *DigitalIDShareErrorResponse `json:"error"`
}

// DigitalIDShareErrorResponse represents the error of a failed digital ID share
type DigitalIDShareErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
