package attribute

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoshare"
	"google.golang.org/protobuf/proto"
)

// IssuanceDetails contains information about the attribute(s) issued by a third party
type IssuanceDetails struct {
	token      string
	expiryDate *time.Time
	attributes []Definition
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
func (i IssuanceDetails) Attributes() []Definition {
	return i.attributes
}

// ParseIssuanceDetails takes the Third Party Attribute object and converts it into an IssuanceDetails struct
func ParseIssuanceDetails(thirdPartyAttributeBytes []byte) (*IssuanceDetails, error) {
	thirdPartyAttributeStruct := &yotiprotoshare.ThirdPartyAttribute{}
	if err := proto.Unmarshal(thirdPartyAttributeBytes, thirdPartyAttributeStruct); err != nil {
		return nil, fmt.Errorf("unable to parse ThirdPartyAttribute value: %q. Error: %q", string(thirdPartyAttributeBytes), err)
	}

	var issuingAttributesProto = thirdPartyAttributeStruct.GetIssuingAttributes()
	var issuingAttributeDefinitions = parseIssuingAttributeDefinitions(issuingAttributesProto.GetDefinitions())

	expiryDate, dateParseErr := parseExpiryDate(issuingAttributesProto.ExpiryDate)

	var issuanceTokenBytes = thirdPartyAttributeStruct.GetIssuanceToken()

	if len(issuanceTokenBytes) == 0 {
		return nil, errors.New("Issuance Token is invalid")
	}

	base64EncodedToken := base64.StdEncoding.EncodeToString(issuanceTokenBytes)

	return &IssuanceDetails{
		token:      base64EncodedToken,
		expiryDate: expiryDate,
		attributes: issuingAttributeDefinitions,
	}, dateParseErr
}

func parseIssuingAttributeDefinitions(definitions []*yotiprotoshare.Definition) (issuingAttributes []Definition) {
	for _, definition := range definitions {
		attributeDefinition := Definition{
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

	parsedTime, err := time.Parse(time.RFC3339Nano, expiryDateString)
	if err != nil {
		return nil, err
	}

	return &parsedTime, err
}
