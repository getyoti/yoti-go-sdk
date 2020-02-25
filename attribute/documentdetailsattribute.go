package attribute

import (
	"fmt"
	"strings"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
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
	*yotiprotoattr.Attribute
	value   DocumentDetails
	anchors []*anchor.Anchor
}

// Value returns the document details struct attached to this attribute
func (attr *DocumentDetailsAttribute) Value() DocumentDetails {
	return attr.value
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (attr *DocumentDetailsAttribute) Anchors() []*anchor.Anchor {
	return attr.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (attr *DocumentDetailsAttribute) Sources() []*anchor.Anchor {
	return anchor.GetSources(attr.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (attr *DocumentDetailsAttribute) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(attr.anchors)
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
		Attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:   details,
		anchors: parsedAnchors,
	}, nil
}

// Parse fills a DocumentDetails object from a raw string
func (details *DocumentDetails) Parse(data string) (err error) {
	dataSlice := strings.Split(data, " ")

	if len(dataSlice) < 3 {
		return fmt.Errorf("Document Details data is invalid")
	}
	for _, section := range dataSlice {
		if section == "" {
			return fmt.Errorf("Document Details data is invalid")
		}
	}

	details.DocumentType = dataSlice[0]
	details.IssuingCountry = dataSlice[1]
	details.DocumentNumber = dataSlice[2]
	if len(dataSlice) > 3 && dataSlice[3] != "-" {
		var dateerr error
		expirationDateData, dateerr := time.Parse(documentDetailsDateFormatConst, dataSlice[3])

		if dateerr == nil {
			details.ExpirationDate = &expirationDateData
		} else {
			return dateerr
		}
	}
	if len(dataSlice) > 4 {
		details.IssuingAuthority = dataSlice[4]
	}
	return
}
