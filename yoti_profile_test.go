package yoti

import (
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/consts"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func createProfileWithSingleAttribute(attr *yotiprotoattr.Attribute) Profile {
	var attributeSlice []*yotiprotoattr.Attribute
	attributeSlice = append(attributeSlice, attr)

	return Profile{
		baseProfile{
			attributeSlice: attributeSlice,
		},
	}
}

func createAppProfileWithSingleAttribute(attr *yotiprotoattr.Attribute) ApplicationProfile {
	var attributeSlice []*yotiprotoattr.Attribute
	attributeSlice = append(attributeSlice, attr)

	return ApplicationProfile{
		baseProfile{
			attributeSlice: attributeSlice,
		},
	}
}

func createProfileWithMultipleAttributes(list ...*yotiprotoattr.Attribute) Profile {
	return Profile{
		baseProfile{
			attributeSlice: list,
		},
	}
}

func TestProfile_AgeVerifications(t *testing.T) {
	ageOver14 := &yotiprotoattr.Attribute{
		Name:        "age_over:14",
		Value:       []byte("true"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	ageUnder18 := &yotiprotoattr.Attribute{
		Name:        "age_under:18",
		Value:       []byte("true"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	ageOver18 := &yotiprotoattr.Attribute{
		Name:        "age_over:18",
		Value:       []byte("false"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithMultipleAttributes(ageOver14, ageUnder18, ageOver18)
	ageVerifications, err := profile.AgeVerifications()

	assert.NilError(t, err)
	assert.Equal(t, len(ageVerifications), 3)

	assert.Equal(t, ageVerifications[0].Age, 14)
	assert.Equal(t, ageVerifications[0].CheckType, "age_over")
	assert.Equal(t, ageVerifications[0].Result, true)

	assert.Equal(t, ageVerifications[1].Age, 18)
	assert.Equal(t, ageVerifications[1].CheckType, "age_under")
	assert.Equal(t, ageVerifications[1].Result, true)

	assert.Equal(t, ageVerifications[2].Age, 18)
	assert.Equal(t, ageVerifications[2].CheckType, "age_over")
	assert.Equal(t, ageVerifications[2].Result, false)
}

func TestProfile_GetAttribute_EmptyString(t *testing.T) {
	emptyString := ""
	attributeValue := []byte(emptyString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	assert.Equal(t, att.Name(), attributeName)
	assert.Equal(t, att.Value().(string), emptyString)
}

func TestProfile_GetApplicationAttribute(t *testing.T) {
	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	appProfile := createProfileWithSingleAttribute(attr)
	attribute := appProfile.GetAttribute(attributeName)
	assert.Equal(t, attribute.Name(), attributeName)
}

func TestProfile_GetApplicationName(t *testing.T) {
	attributeValue := "APPLICATION NAME"
	var attr = &yotiprotoattr.Attribute{
		Name:        "application_name",
		Value:       []byte(attributeValue),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	appProfile := createAppProfileWithSingleAttribute(attr)
	assert.Equal(t, attributeValue, appProfile.ApplicationName().Value())
}

func TestProfile_GetApplicationURL(t *testing.T) {
	attributeValue := "APPLICATION URL"
	var attr = &yotiprotoattr.Attribute{
		Name:        "application_url",
		Value:       []byte(attributeValue),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	appProfile := createAppProfileWithSingleAttribute(attr)
	assert.Equal(t, attributeValue, appProfile.ApplicationURL().Value())
}

func TestProfile_GetApplicationLogo(t *testing.T) {
	attributeValue := "APPLICATION LOGO"
	var attr = &yotiprotoattr.Attribute{
		Name:        "application_logo",
		Value:       []byte(attributeValue),
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	appProfile := createAppProfileWithSingleAttribute(attr)
	assert.Equal(t, 16, len(appProfile.ApplicationLogo().Value().Data))
}

func TestProfile_GetApplicationBGColor(t *testing.T) {
	attributeValue := "BG VALUE"
	var attr = &yotiprotoattr.Attribute{
		Name:        "application_receipt_bgcolor",
		Value:       []byte(attributeValue),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	appProfile := createAppProfileWithSingleAttribute(attr)
	assert.Equal(t, attributeValue, appProfile.ApplicationReceiptBgColor().Value())
}

func TestProfile_GetAttribute_Int(t *testing.T) {
	intValues := [5]int{0, 1, 123, -10, -1}

	for _, integer := range intValues {
		assertExpectedIntegerIsReturned(t, integer)
	}
}

func assertExpectedIntegerIsReturned(t *testing.T, intValue int) {
	intAsString := strconv.Itoa(intValue)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       []byte(intAsString),
		ContentType: yotiprotoattr.ContentType_INT,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	assert.Equal(t, att.Value().(int), intValue)
}

func TestProfile_GetAttribute_InvalidInt_ReturnsNil(t *testing.T) {
	invalidIntValue := "1985-01-01"

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       []byte(invalidIntValue),
		ContentType: yotiprotoattr.ContentType_INT,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)

	log.SetOutput(ioutil.Discard)
	att := result.GetAttribute(attributeName)

	assert.Assert(t, is.Nil(att))
}

func TestProfile_EmptyStringIsAllowed(t *testing.T) {
	attributeValueString := ""
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        consts.AttrGender,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(attr)
	att := profile.Gender()

	assert.Equal(t, att.Value(), attributeValueString)
}

func TestProfile_GetAttribute_Time(t *testing.T) {
	dateStringValue := "1985-01-01"
	expectedDate := time.Date(1985, time.January, 1, 0, 0, 0, 0, time.UTC)

	attributeValueTime := []byte(dateStringValue)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValueTime,
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	assert.Equal(t, expectedDate, att.Value().(*time.Time).UTC())
}

func TestProfile_GetAttribute_Jpeg(t *testing.T) {
	attributeValue := []byte("value")
	expected := attribute.Image{
		Type: attribute.ImageTypeJpeg,
		Data: attributeValue,
	}

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	assert.DeepEqual(t, att.Value().(*attribute.Image), &expected)
}

func TestProfile_GetAttribute_Png(t *testing.T) {
	attributeValue := []byte("value")
	expected := attribute.Image{
		Type: attribute.ImageTypePng,
		Data: attributeValue,
	}

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	assert.DeepEqual(t, att.Value().(*attribute.Image), &expected)
}

func TestProfile_GetAttribute_Bool(t *testing.T) {
	var initialBoolValue = true
	attributeValue := []byte(strconv.FormatBool(initialBoolValue))

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	boolValue, err := strconv.ParseBool(att.Value().(string))

	assert.Assert(t, is.Nil(err))
	assert.Equal(t, initialBoolValue, boolValue)
}

func TestProfile_GetAttribute_JSON(t *testing.T) {
	addressFormat := "2"

	var structuredAddressBytes = []byte(`
		{
			"address_format": "` + addressFormat + `",
			"building": "House No.86-A"
		}`)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       structuredAddressBytes,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	retrievedAttributeMap := att.Value().(map[string]interface{})
	actualAddressFormat := retrievedAttributeMap["address_format"]

	assert.Equal(t, actualAddressFormat, addressFormat)
}

func TestProfile_GetAttribute_Undefined(t *testing.T) {
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	assert.Equal(t, att.Name(), attributeName)
	assert.Equal(t, att.Value().(string), attributeValueString)
}

func TestProfile_GetAttribute_ReturnsNil(t *testing.T) {
	result := Profile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	attribute := result.GetAttribute("attributeName")

	assert.Assert(t, is.Nil(attribute))
}

func TestProfile_StringAttribute(t *testing.T) {
	attributeName := consts.AttrNationality
	attributeValueString := "value"
	attributeValueBytes := []byte(attributeValueString)

	var as = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValueBytes,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(as)

	assert.Equal(t, result.Nationality().Value(), attributeValueString)

	assert.Equal(t, result.Nationality().ContentType(), yotiprotoattr.ContentType_STRING.String())
}

func TestProfile_AttributeProperty_RetrievesAttribute(t *testing.T) {
	attributeName := consts.AttrSelfie
	attributeValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.Equal(t, selfie.Name(), attributeName)
	assert.DeepEqual(t, attributeValue, selfie.Value().Data)
	assert.Equal(t, selfie.ContentType(), yotiprotoattr.ContentType_PNG.String())
}

func TestProfile_DocumentDetails_RetrievesAttribute(t *testing.T) {
	attributeName := consts.AttrDocumentDetails
	attributeValue := []byte("PASSPORT GBR 1234567")

	var proto = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     make([]*yotiprotoattr.Anchor, 0),
	}

	result := createProfileWithSingleAttribute(proto)
	documentDetails, err := result.DocumentDetails()
	assert.NilError(t, err)

	assert.Equal(t, documentDetails.Value().DocumentType, "PASSPORT")
}

func TestProfile_DocumentImages_RetrievesAttribute(t *testing.T) {
	attributeName := consts.AttrDocumentImages
	attributeValue, err := proto.Marshal(&yotiprotoattr.MultiValue{})
	assert.NilError(t, err)

	proto := &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     make([]*yotiprotoattr.Anchor, 0),
	}

	result := createProfileWithSingleAttribute(proto)
	documentImages, err := result.DocumentImages()
	assert.NilError(t, err)

	assert.Equal(t, documentImages.Name(), consts.AttrDocumentImages)
}