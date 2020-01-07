package extension

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"gotest.tools/assert"
)

func createDefinitionByName(name string) attribute.AttributeDefinition {
	return attribute.NewAttributeDefinition(name)
}

func ExampleThirdPartyAttributeExtension() {
	attributeDefinition := attribute.NewAttributeDefinition("some_value")

	datetime, err := time.Parse("2006-01-02T15:04:05.999Z", "2019-10-30T12:10:09.458Z")
	if err != nil {
		log.Printf("Error parsing date, %v", err)
	}

	extension := (&ThirdPartyAttributeExtensionBuilder{}).
		WithExpiryDate(&datetime).
		WithDefinition(attributeDefinition).
		Build()

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"THIRD_PARTY_ATTRIBUTE","content":{"expiry_date":"2019-10-30T12:10:09.458Z","definitions":[{"name":"some_value"}]}}
}

func TestWithDefinitionShouldAddToList(t *testing.T) {
	datetime, err := time.Parse("2006-01-02T15:04:05.999Z", "2019-10-30T12:10:09.458Z")
	if err != nil {
		log.Printf("Error parsing date, %v", err)
	}

	definitionList := []attribute.AttributeDefinition{
		createDefinitionByName("some_attribute"),
		createDefinitionByName("some_other_attribute"),
	}

	someOtherDefinition := createDefinitionByName("wanted_definition")

	extension := (&ThirdPartyAttributeExtensionBuilder{}).
		WithExpiryDate(&datetime).
		WithDefinitions(definitionList).
		WithDefinition(someOtherDefinition).
		Build()

	assert.Equal(t, len(extension.definitions), 3)
	assert.Equal(t, extension.definitions[0].Name(), "some_attribute")
	assert.Equal(t, extension.definitions[1].Name(), "some_other_attribute")
	assert.Equal(t, extension.definitions[2].Name(), "wanted_definition")
}

func TestWithDefinitionsShouldOverwriteList(t *testing.T) {
	datetime, err := time.Parse("2006-01-02T15:04:05.999Z", "2019-10-30T12:10:09.458Z")
	if err != nil {
		log.Printf("Error parsing date, %v", err)
	}

	definitionList := []attribute.AttributeDefinition{
		createDefinitionByName("some_attribute"),
		createDefinitionByName("some_other_attribute"),
	}

	someOtherDefinition := createDefinitionByName("wanted_definition")

	extension := (&ThirdPartyAttributeExtensionBuilder{}).
		WithExpiryDate(&datetime).
		WithDefinition(someOtherDefinition).
		WithDefinitions(definitionList).
		Build()

	assert.Equal(t, len(extension.definitions), 2)
	assert.Equal(t, extension.definitions[0].Name(), "some_attribute")
	assert.Equal(t, extension.definitions[1].Name(), "some_other_attribute")
}
