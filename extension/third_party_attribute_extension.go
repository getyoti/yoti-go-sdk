package extension

import (
	"encoding/json"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
)

const (
	thirdPartyAttributeExtensionTypeConst = "THIRD_PARTY_ATTRIBUTE"
)

// ThirdPartyAttributeExtensionBuilder is used to construct a ThirdPartyAttributeExtension
type ThirdPartyAttributeExtensionBuilder struct {
	extension ThirdPartyAttributeExtension
}

// ThirdPartyAttributeExtension is an extension representing the issuance of a third party attribute
type ThirdPartyAttributeExtension struct {
	expiryDate  *time.Time
	definitions []attribute.AttributeDefinition
}

// WithExpiryDate sets the expiry date of the extension as a UTC timestamp
func (builder *ThirdPartyAttributeExtensionBuilder) WithExpiryDate(expiryDate *time.Time) *ThirdPartyAttributeExtensionBuilder {
	builder.extension.expiryDate = expiryDate
	return builder
}

// WithDefinition adds an attribute.AttributeDefinition to the list of definitions
func (builder *ThirdPartyAttributeExtensionBuilder) WithDefinition(definition attribute.AttributeDefinition) *ThirdPartyAttributeExtensionBuilder {
	builder.extension.definitions = append(builder.extension.definitions, definition)
	return builder
}

// WithDefinitions sets the array of attribute.AttributeDefinition on the extension
func (builder *ThirdPartyAttributeExtensionBuilder) WithDefinitions(definitions []attribute.AttributeDefinition) *ThirdPartyAttributeExtensionBuilder {
	builder.extension.definitions = definitions
	return builder
}

// Build creates a ThirdPartyAttributeExtension using the supplied values
func (builder *ThirdPartyAttributeExtensionBuilder) Build() ThirdPartyAttributeExtension {
	return builder.extension
}

// MarshalJSON returns the JSON encoding
func (extension ThirdPartyAttributeExtension) MarshalJSON() ([]byte, error) {
	type thirdPartyAttributeExtension struct {
		ExpiryDate  string                          `json:"expiry_date"`
		Definitions []attribute.AttributeDefinition `json:"definitions"`
	}
	return json.Marshal(&struct {
		Type    string                       `json:"type"`
		Content thirdPartyAttributeExtension `json:"content"`
	}{
		Type: thirdPartyAttributeExtensionTypeConst,
		Content: thirdPartyAttributeExtension{
			ExpiryDate:  extension.expiryDate.UTC().Format("2006-01-02T15:04:05.000Z"),
			Definitions: extension.definitions,
		},
	})
}
