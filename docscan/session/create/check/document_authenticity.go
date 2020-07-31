package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedDocumentAuthenticityCheck requests creation of a Document Authenticity Check
type RequestedDocumentAuthenticityCheck struct {
	RequestedCheck
	Config RequestedDocumentAuthenticityConfig `json:"config"`
}

// RequestedDocumentAuthenticityConfig is the configuration applied when creating a Document Authenticity Check
type RequestedDocumentAuthenticityConfig struct {
	RequestedCheckConfig
}

// NewRequestedDocumentAuthenticityCheck creates a new Document Authenticity Check
func NewRequestedDocumentAuthenticityCheck(config RequestedDocumentAuthenticityConfig) *RequestedDocumentAuthenticityCheck {
	return &RequestedDocumentAuthenticityCheck{
		RequestedCheck: RequestedCheck{
			Type: constants.IDDocumentAuthenticity,
		},
		Config: config,
	}
}

// RequestedDocumentAuthenticityCheckBuilder builds a RequestedDocumentAuthenticityCheck
type RequestedDocumentAuthenticityCheckBuilder struct {
	config RequestedDocumentAuthenticityConfig
}

// NewRequestedDocumentAuthenticityCheckBuilder creates a new DocumentAuthenticityCheckBuilder
func NewRequestedDocumentAuthenticityCheckBuilder() *RequestedDocumentAuthenticityCheckBuilder {
	return &RequestedDocumentAuthenticityCheckBuilder{}
}

// Build creates a new DocumentAuthenticityCheck
func (b *RequestedDocumentAuthenticityCheckBuilder) Build() (*RequestedDocumentAuthenticityCheck, error) {
	return NewRequestedDocumentAuthenticityCheck(b.config), nil
}
