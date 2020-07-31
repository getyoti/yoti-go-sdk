package retrieve

import (
	"github.com/getyoti/yoti-go-sdk/v3/util"
)

// IDDocumentResourceResponse represents an Identity Document resource for a given session
type IDDocumentResourceResponse struct {
	ResourceResponse
	// DocumentType is the identity document type, e.g. "PASSPORT"
	DocumentType string `json:"document_type"`
	// IssuingCountry is the issuing country of the identity document
	IssuingCountry string `json:"issuing_country"`
	// Pages are the individual pages of the identity document
	Pages []PageResponse `json:"pages"`
	// DocumentFields are the associated document fields of a document
	DocumentFields  DocumentFieldsResponse  `json:"document_fields"`
	DocumentIDPhoto DocumentIDPhotoResponse `json:"document_id_photo"`
}

// GetTextExtractionTasks returns a slice of text extraction tasks associated with the ID document
func (id IDDocumentResourceResponse) GetTextExtractionTasks() []TextExtractionTaskResponse {
	filteredTasks := util.Filter(id.Tasks, func(val interface{}) bool {
		_, isTextExtractionTaskResponse := val.(TextExtractionTaskResponse)
		return isTextExtractionTaskResponse
	})

	return filteredTasks.([]TextExtractionTaskResponse)
}
