package attribute

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

var (
	documentDetailsValidateRegexp = regexp.MustCompile("^[A-Za-z_]* [A-Za-z]{3} [A-Za-z0-9]{1}.*$")
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

	expirationDateData time.Time
}

// DocumentDetailsAttribute wraps a document details with anchor data
type DocumentDetailsAttribute struct {
	*yotiprotoattr.Attribute
	value   DocumentDetails
	anchors []*anchor.Anchor
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

// Parse filles a DocumentDetails object from a raw string
func (details *DocumentDetails) Parse(data string) (err error) {
	if !documentDetailsValidateRegexp.MatchString(data) {
		return fmt.Errorf("Document Details data is invalid, %s", data)
	}

	dataSlice := strings.Split(data, " ")

	details.DocumentType = dataSlice[0]
	details.IssuingCountry = dataSlice[1]
	details.DocumentNumber = dataSlice[2]
	if len(dataSlice) > 3 {
		var dateerr error
		details.expirationDateData, dateerr = time.Parse(documentDetailsDateFormatConst, dataSlice[3])
		if dateerr == nil {
			details.ExpirationDate = &details.expirationDateData
		}
	}
	if len(dataSlice) > 4 {
		details.IssuingAuthority = dataSlice[4]
	}
	return
}
