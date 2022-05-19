package attribute

import (
	"fmt"
	"strings"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

const (
	documentDetailsDateFormatConst = "2006-01-02"
)

// DocumentDetails represents information extracted from a document provided by the user
type DocumentDetails struct {
	DocumentType     string
	IssuingCountry   string
	DocumentNumber   string
	ExpirationDate   *time.Time
	IssuingAuthority string
}

// DocumentDetailsAttribute wraps a document details with anchor data
type DocumentDetailsAttribute struct {
	attributeDetails
	value DocumentDetails
}

// Value returns the document details struct attached to this attribute
func (attr *DocumentDetailsAttribute) Value() DocumentDetails {
	return attr.value
}

// NewDocumentDetails creates a DocumentDetailsAttribute which wraps a
// DocumentDetails with anchor data
func NewDocumentDetails(a *yotiprotoattr.Attribute) (*DocumentDetailsAttribute, error) {
	parsedAnchors := anchor.ParseAnchors(a.Anchors)
	details := DocumentDetails{}
	err := details.Parse(string(a.Value))
	if err != nil {
		return nil, err
	}

	return &DocumentDetailsAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
			id:          &a.EphemeralId,
		},
		value: details,
	}, nil
}

// Parse fills a DocumentDetails object from a raw string
func (details *DocumentDetails) Parse(data string) error {
	dataSlice := strings.Split(data, " ")

	if len(dataSlice) < 3 {
		return fmt.Errorf("Document Details data is invalid, %s", data)
	}
	for _, section := range dataSlice {
		if section == "" {
			return fmt.Errorf("Document Details data is invalid %s", data)
		}
	}

	details.DocumentType = dataSlice[0]
	details.IssuingCountry = dataSlice[1]
	details.DocumentNumber = dataSlice[2]
	if len(dataSlice) > 3 && dataSlice[3] != "-" {
		expirationDateData, dateErr := time.Parse(documentDetailsDateFormatConst, dataSlice[3])

		if dateErr == nil {
			details.ExpirationDate = &expirationDateData
		} else {
			return dateErr
		}
	}
	if len(dataSlice) > 4 {
		details.IssuingAuthority = dataSlice[4]
	}

	return nil
}
