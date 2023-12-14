package digitalidentity

import (
	"reflect"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

func getFormattedAddress(profile *UserProfile, formattedAddress string) *yotiprotoattr.Attribute {
	proto := getProtobufAttribute(*profile, consts.AttrStructuredPostalAddress)

	return &yotiprotoattr.Attribute{
		Name:        consts.AttrAddress,
		Value:       []byte(formattedAddress),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     proto.Anchors,
	}
}

func ensureAddressProfile(p *UserProfile) *attribute.StringAttribute {
	if structuredPostalAddress, err := p.StructuredPostalAddress(); err == nil {
		if (structuredPostalAddress != nil && !reflect.DeepEqual(structuredPostalAddress, attribute.JSONAttribute{})) {
			var formattedAddress string
			formattedAddress, err = retrieveFormattedAddressFromStructuredPostalAddress(structuredPostalAddress.Value())
			if err == nil && formattedAddress != "" {
				return attribute.NewString(getFormattedAddress(p, formattedAddress))
			}
		}
	}

	return nil
}

func retrieveFormattedAddressFromStructuredPostalAddress(structuredPostalAddress interface{}) (address string, err error) {
	parsedStructuredAddressMap := structuredPostalAddress.(map[string]interface{})
	if formattedAddress, ok := parsedStructuredAddressMap["formatted_address"]; ok {
		return formattedAddress.(string), nil
	}
	return
}

func getProtobufAttribute(profile UserProfile, key string) *yotiprotoattr.Attribute {
	for _, v := range profile.attributeSlice {
		if v.Name == key {
			return v
		}
	}

	return nil
}
