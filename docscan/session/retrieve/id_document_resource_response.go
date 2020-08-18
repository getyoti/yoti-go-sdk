package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// IDDocumentResourceResponse represents an Identity Document resource for a given session
type IDDocumentResourceResponse struct {
	*ResourceResponse
	// DocumentType is the identity document type, e.g. "PASSPORT"
	DocumentType string `json:"document_type"`
	// IssuingCountry is the issuing country of the identity document
	IssuingCountry string `json:"issuing_country"`
	// Pages are the individual pages of the identity document
	Pages []*PageResponse `json:"pages"`
	// DocumentFields are the associated document fields of a document
	DocumentFields      *DocumentFieldsResponse  `json:"document_fields"`
	DocumentIDPhoto     *DocumentIDPhotoResponse `json:"document_id_photo"`
	textExtractionTasks []*TextExtractionTaskResponse
}

// TextExtractionTasks returns a list of text extraction tasks associated with the ID document
func (i *IDDocumentResourceResponse) TextExtractionTasks() []*TextExtractionTaskResponse {
	return i.textExtractionTasks
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (i *IDDocumentResourceResponse) UnmarshalJSON(data []byte) error {
	type result IDDocumentResourceResponse // declared as "type" to prevent recursive unmarshalling
	if err := json.Unmarshal(data, (*result)(i)); err != nil {
		return err
	}

	for _, task := range i.Tasks {
		switch task.Type {
		case constants.IDDocumentTextDataExtraction:
			i.textExtractionTasks = append(i.textExtractionTasks, &TextExtractionTaskResponse{TaskResponse: task})
		}
	}

	return nil
}
