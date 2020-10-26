package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// SupplementaryDocumentResourceResponse represents an supplementary document resource for a given session
type SupplementaryDocumentResourceResponse struct {
	*ResourceResponse
	// DocumentType is the supplementary document type, e.g. "UTILITY_BILL"
	DocumentType string `json:"document_type"`
	// IssuingCountry is the issuing country of the supplementary document
	IssuingCountry string `json:"issuing_country"`
	// Pages are the individual pages of the supplementary document
	Pages []*PageResponse `json:"pages"`
	// DocumentFields are the associated document fields of a document
	DocumentFields *DocumentFieldsResponse `json:"document_fields"`
	// DocumentFile is the associated document file
	DocumentFile        *FileResponse `json:"file"`
	textExtractionTasks []*SupplementaryDocumentTextExtractionTaskResponse
}

// TextExtractionTasks returns a slice of text extraction tasks associated with the supplementary document
func (i *SupplementaryDocumentResourceResponse) TextExtractionTasks() []*SupplementaryDocumentTextExtractionTaskResponse {
	return i.textExtractionTasks
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (i *SupplementaryDocumentResourceResponse) UnmarshalJSON(data []byte) error {
	type result SupplementaryDocumentResourceResponse // declared as "type" to prevent recursive unmarshalling
	if err := json.Unmarshal(data, (*result)(i)); err != nil {
		return err
	}

	for _, task := range i.Tasks {
		switch task.Type {
		case constants.SupplementaryDocumentTextDataExtraction:
			i.textExtractionTasks = append(
				i.textExtractionTasks,
				&SupplementaryDocumentTextExtractionTaskResponse{
					TaskResponse: task,
				},
			)
		}
	}

	return nil
}
