package attribute

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"
)

// IssuanceDetails contains information about the attribute(s) issued by a third party
type IssuanceDetails struct {
	token      string
	expiryDate *time.Time
	attributes []AttributeDefinition
}

// Token is the issuance token that can be used to retrieve the user's stored details.
// These details will be used to issue attributes on behalf of an organisation to that user.
func (i IssuanceDetails) Token() string {
	return i.token
}

// ExpiryDate is the timestamp at which the request for the attribute value
// from third party will expire. Will be nil if not provided.
func (i IssuanceDetails) ExpiryDate() *time.Time {
	return i.expiryDate
}

// Attributes information about the attributes the third party would like to issue.
func (i IssuanceDetails) Attributes() []AttributeDefinition {
	return i.attributes
}

// ParseIssuanceDetails takes the Third Party Attribute object and converts it into an IssuanceDetails struct
func ParseIssuanceDetails(thirdPartyAttributeBytes []byte) (*IssuanceDetails, error) {
	thirdPartyAttributeStruct := &yotiprotoshare.ThirdPartyAttribute{}
	if err := proto.Unmarshal(thirdPartyAttributeBytes, thirdPartyAttributeStruct); err != nil {
		return nil, fmt.Errorf("Unable to parse ThirdPartyAttribute value: %q. Error: %q", string(thirdPartyAttributeBytes), err)
	}

	var issuingAttributesProto *yotiprotoshare.IssuingAttributes = thirdPartyAttributeStruct.GetIssuingAttributes()
	var issuingAttributeDefinitions []AttributeDefinition = parseIssuingAttributeDefinitions(issuingAttributesProto.GetDefinitions())

	expiryDate, dateParseErr := parseExpiryDate(issuingAttributesProto.ExpiryDate)

	var issuanceToken string = string(thirdPartyAttributeStruct.GetIssuanceToken())

	if issuanceToken == "" {
		return nil, errors.New("Issuance Token is invalid")
	}

	return &IssuanceDetails{
		token:      issuanceToken,
		expiryDate: expiryDate,
		attributes: issuingAttributeDefinitions,
	}, dateParseErr
}

func parseIssuingAttributeDefinitions(definitions []*yotiprotoshare.Definition) (issuingAttributes []AttributeDefinition) {
	for _, definition := range definitions {
		attributeDefinition := AttributeDefinition{
			name: definition.Name,
		}
		issuingAttributes = append(issuingAttributes, attributeDefinition)
	}

	return issuingAttributes
}

func parseExpiryDate(expiryDateString string) (*time.Time, error) {
	if expiryDateString == "" {
		return nil, nil
	}

	parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", expiryDateString)
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", expiryDateString, err)
		return nil, err
	}

	return &parsedTime, err
}
