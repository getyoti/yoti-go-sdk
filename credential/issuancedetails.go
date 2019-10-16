package credential

import (
	"fmt"
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"
)

// IssuanceDetails contains information about the credential(s) issued by a third party
type IssuanceDetails struct {
	token             string
	expiryDate        *time.Time
	issuingAttributes []string
}

// Token is the issuance token that can be used to retrieve the user's stored details.
// These details will be used to issue attributes on behalf of an organisation to that user.
func (i IssuanceDetails) Token() string {
	return i.token
}

// ExpiryDate is the timestamp at which the request for the attribute value
// from third party will expire.
func (i IssuanceDetails) ExpiryDate() *time.Time {
	return i.expiryDate
}

// IssuingAttributes name of the attributes the third party would like to issue.
func (i IssuanceDetails) IssuingAttributes() []string {
	return i.issuingAttributes
}

// ParseIssuanceDetails takes the Third Party Attribute object and converts it into an IssuanceDetails struct
func ParseIssuanceDetails(thirdPartyAttributeBytes []byte) (*IssuanceDetails, error) {
	thirdPartyAttributeStruct := &yotiprotoshare.ThirdPartyAttribute{}
	if err := proto.Unmarshal(thirdPartyAttributeBytes, thirdPartyAttributeStruct); err != nil {
		return nil, fmt.Errorf("Unable to parse ThirdPartyAttribute value: %q. Error: %q", string(thirdPartyAttributeBytes), err)
	}

	var issuingAttributesProto *yotiprotoshare.IssuingAttributes = thirdPartyAttributeStruct.GetIssuingAttributes()
	var issuingAttributeDefinitions []string = parseIssuingAttributeDefinitions(issuingAttributesProto.GetDefinitions())

	expiryDate, err := parseExpiryDate(issuingAttributesProto.ExpiryDate)

	if err != nil {
		return nil, err
	}

	return &IssuanceDetails{
		token:             string(thirdPartyAttributeStruct.GetIssuanceToken()),
		expiryDate:        expiryDate,
		issuingAttributes: issuingAttributeDefinitions,
	}, nil
}

func parseIssuingAttributeDefinitions(definitions []*yotiprotoshare.Definition) (issuingAttributes []string) {
	for _, definition := range definitions {
		issuingAttributes = append(issuingAttributes, definition.GetName())
	}

	return issuingAttributes
}

func parseExpiryDate(expiryDateString string) (*time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", expiryDateString)
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", expiryDateString, err)
		parsedTime = time.Time{}
	}

	return &parsedTime, err
}
