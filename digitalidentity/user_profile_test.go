package digitalidentity

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/file"
	"github.com/getyoti/yoti-go-sdk/v3/media"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

const (
	attributeName        = "test_attribute_name"
	attributeValueString = "value"

	documentImagesAttributeID = "document-images-attribute-id-123"
	selfieAttributeID         = "selfie-attribute-id-123"
	fullNameAttributeID       = "full-name-id-123"
)

var attributeValue = []byte(attributeValueString)

func getUserProfile() UserProfile {
	userProfile := createProfileWithMultipleAttributes(
		createDocumentImagesAttribute(documentImagesAttributeID),
		createSelfieAttribute(yotiprotoattr.ContentType_JPEG, selfieAttributeID),
		createStringAttribute("full_name", []byte("John Smith"), []*yotiprotoattr.Anchor{}, fullNameAttributeID))

	return userProfile
}

func ExampleUserProfile_GetAttributeByID() {
	userProfile := getUserProfile()
	fullNameAttribute := userProfile.GetAttributeByID("full-name-id-123")
	value := fullNameAttribute.Value().(string)

	fmt.Println(value)
	// Output: John Smith
}

func ExampleUserProfile_GetDocumentImagesAttributeByID() {
	userProfile := getUserProfile()
	documentImagesAttribute, err := userProfile.GetDocumentImagesAttributeByID("document-images-attribute-id-123")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(*documentImagesAttribute.ID())
	// Output: document-images-attribute-id-123
}

func ExampleUserProfile_GetSelfieAttributeByID() {
	userProfile := getUserProfile()
	selfieAttribute, err := userProfile.GetSelfieAttributeByID("selfie-attribute-id-123")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(*selfieAttribute.ID())
	// Output: selfie-attribute-id-123
}

func createProfileWithSingleAttribute(attr *yotiprotoattr.Attribute) UserProfile {
	var attributeSlice []*yotiprotoattr.Attribute
	attributeSlice = append(attributeSlice, attr)

	return UserProfile{
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

func createProfileWithMultipleAttributes(list ...*yotiprotoattr.Attribute) UserProfile {
	return UserProfile{
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
	applicationAttribute := appProfile.GetAttribute(attributeName)
	assert.Equal(t, applicationAttribute.Name(), attributeName)
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
	assert.Equal(t, 16, len(appProfile.ApplicationLogo().Value().Data()))
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

	att := result.GetAttribute(attributeName)

	assert.Assert(t, is.Nil(att))
}

func TestProfile_EmptyStringIsAllowed(t *testing.T) {
	emptyString := ""
	attrValue := []byte(emptyString)

	var attr = &yotiprotoattr.Attribute{
		Name:        consts.AttrGender,
		Value:       attrValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(attr)
	att := profile.Gender()

	assert.Equal(t, att.Value(), emptyString)
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
	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(attr)
	att := profile.GetAttribute(attributeName)

	expected := media.JPEGImage(attributeValue)
	result := att.Value().(media.JPEGImage)

	assert.DeepEqual(t, expected, result)
	assert.Equal(t, expected.Base64URL(), result.Base64URL())
}

func TestProfile_GetAttribute_Png(t *testing.T) {
	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(attr)
	att := profile.GetAttribute(attributeName)

	expected := media.PNGImage(attributeValue)
	result := att.Value().(media.PNGImage)

	assert.DeepEqual(t, expected, result)
	assert.Equal(t, expected.Base64URL(), result.Base64URL())
}

func TestProfile_GetAttribute_Bool(t *testing.T) {
	var initialBoolValue = true
	attrValue := []byte(strconv.FormatBool(initialBoolValue))

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attrValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	boolValue, err := strconv.ParseBool(att.Value().(string))

	assert.NilError(t, err)
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
	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	result := userProfile.GetAttribute("attributeName")

	assert.Assert(t, is.Nil(result))
}

func TestProfile_GetAttributeByID(t *testing.T) {
	attributeID := "att-id-123"

	var attr1 = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
		EphemeralId: attributeID,
	}
	var attr2 = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
		EphemeralId: "non-matching-attribute-ID",
	}

	profile := createProfileWithMultipleAttributes(attr1, attr2)

	result := profile.GetAttributeByID(attributeID)
	assert.DeepEqual(t, result.ID(), &attributeID)
}

func TestProfile_GetAttributeByID_ReturnsNil(t *testing.T) {
	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	result := userProfile.GetAttributeByID("attributeName")

	assert.Assert(t, is.Nil(result))
}

func TestProfile_GetDocumentImagesAttributeByID_ReturnsNil(t *testing.T) {
	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	result, err := userProfile.GetDocumentImagesAttributeByID("attributeName")
	assert.NilError(t, err)
	assert.Assert(t, is.Nil(result))
}

func TestProfile_GetSelfieAttributeByID_ReturnsNil(t *testing.T) {
	userProfile := UserProfile{
		baseProfile{
			attributeSlice: []*yotiprotoattr.Attribute{},
		},
	}

	result, err := userProfile.GetSelfieAttributeByID("attributeName")
	assert.NilError(t, err)
	assert.Assert(t, is.Nil(result))
}

func TestProfile_StringAttribute(t *testing.T) {
	nationalityName := consts.AttrNationality

	var as = &yotiprotoattr.Attribute{
		Name:        nationalityName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(as)

	assert.Equal(t, result.Nationality().Value(), attributeValueString)

	assert.Equal(t, result.Nationality().ContentType(), yotiprotoattr.ContentType_STRING.String())
}

func TestProfile_AttributeProperty_RetrievesAttribute(t *testing.T) {
	attributeImage := createSelfieAttribute(yotiprotoattr.ContentType_PNG, "id")

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.Equal(t, selfie.Name(), consts.AttrSelfie)
	assert.DeepEqual(t, attributeValue, selfie.Value().Data())
	assert.Equal(t, selfie.ContentType(), yotiprotoattr.ContentType_PNG.String())
}

func TestProfile_DocumentDetails_RetrievesAttribute(t *testing.T) {
	documentDetailsName := consts.AttrDocumentDetails
	attributeValue := []byte("PASSPORT GBR 1234567")

	var protoAttribute = &yotiprotoattr.Attribute{
		Name:        documentDetailsName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     make([]*yotiprotoattr.Anchor, 0),
	}

	result := createProfileWithSingleAttribute(protoAttribute)
	documentDetails, err := result.DocumentDetails()
	assert.NilError(t, err)

	assert.Equal(t, documentDetails.Value().DocumentType, "PASSPORT")
}

func TestProfile_DocumentImages_RetrievesAttribute(t *testing.T) {
	protoAttribute := createDocumentImagesAttribute("attr-id")

	result := createProfileWithSingleAttribute(protoAttribute)
	documentImages, err := result.DocumentImages()
	assert.NilError(t, err)

	assert.Equal(t, documentImages.Name(), consts.AttrDocumentImages)
}

func TestProfile_AttributesReturnsNilWhenNotPresent(t *testing.T) {
	documentImagesName := consts.AttrDocumentImages
	multiValue, err := proto.Marshal(&yotiprotoattr.MultiValue{})
	assert.NilError(t, err)

	protoAttribute := &yotiprotoattr.Attribute{
		Name:        documentImagesName,
		Value:       multiValue,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     make([]*yotiprotoattr.Anchor, 0),
	}

	result := createProfileWithSingleAttribute(protoAttribute)

	DoB, err := result.DateOfBirth()
	assert.Check(t, DoB == nil)
	assert.Check(t, err == nil)
	assert.Check(t, result.Address() == nil)
}

func TestMissingPostalAddress_UsesFormattedAddress(t *testing.T) {
	var formattedAddressText = `House No.86-A\nRajgura Nagar\nLudhina\nPunjab\n141012\nIndia`

	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A",
		"formatted_address": "` + formattedAddressText + `"
	}
	`)

	var jsonAttribute = &yotiprotoattr.Attribute{
		Name:        consts.AttrStructuredPostalAddress,
		Value:       structuredAddressBytes,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(jsonAttribute)

	ensureAddressProfile(&profile)

	escapedFormattedAddressText := strings.Replace(formattedAddressText, `\n`, "\n", -1)

	profileAddress := profile.Address().Value()
	assert.Equal(t, profileAddress, escapedFormattedAddressText, "Address does not equal the expected formatted address.")

	structuredPostalAddress, err := profile.StructuredPostalAddress()
	assert.NilError(t, err)
	assert.Equal(t, structuredPostalAddress.ContentType(), "JSON")
}

func TestAttributeImage_Image_Png(t *testing.T) {
	attributeImage := createSelfieAttribute(yotiprotoattr.ContentType_PNG, "id")

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data(), attributeValue)
}

func TestAttributeImage_Image_Jpeg(t *testing.T) {
	attributeImage := createSelfieAttribute(yotiprotoattr.ContentType_JPEG, "id")

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data(), attributeValue)
}

func TestAttributeImage_Image_Default(t *testing.T) {
	attributeImage := createSelfieAttribute(yotiprotoattr.ContentType_PNG, "id")

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data(), attributeValue)
}
func TestAttributeImage_Base64Selfie_Png(t *testing.T) {
	attributeImage := createSelfieAttribute(yotiprotoattr.ContentType_PNG, "id")

	result := createProfileWithSingleAttribute(attributeImage)
	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(attributeValue)
	expectedBase64Selfie := "data:image/png;base64," + base64ImageExpectedValue
	base64Selfie := result.Selfie().Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}

func TestAttributeImage_Base64URL_Jpeg(t *testing.T) {
	attributeImage := createSelfieAttribute(yotiprotoattr.ContentType_JPEG, "id")

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(attributeValue)

	expectedBase64Selfie := "data:image/jpeg;base64," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}

func TestProfile_IdentityProfileReport_RetrievesAttribute(t *testing.T) {
	identityProfileReportJSON, err := file.ReadFile("../test/fixtures/RTWIdentityProfileReport.json")
	assert.NilError(t, err)

	var attr = &yotiprotoattr.Attribute{
		Name:        consts.AttrIdentityProfileReport,
		Value:       identityProfileReportJSON,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att, err := result.IdentityProfileReport()
	assert.NilError(t, err)

	retrievedIdentityProfile := att.Value()
	gotProof := retrievedIdentityProfile["proof"]

	assert.Equal(t, gotProof, "<signature provided here>")
}

func TestProfileAllowsMultipleAttributesWithSameName(t *testing.T) {
	firstAttribute := createStringAttribute("full_name", []byte("some_value"), []*yotiprotoattr.Anchor{}, "id")
	secondAttribute := createStringAttribute("full_name", []byte("some_other_value"), []*yotiprotoattr.Anchor{}, "id")

	var attributeSlice []*yotiprotoattr.Attribute
	attributeSlice = append(attributeSlice, firstAttribute, secondAttribute)

	var profile = UserProfile{
		baseProfile{
			attributeSlice: attributeSlice,
		},
	}

	var fullNames = profile.GetAttributes("full_name")

	assert.Assert(t, is.Equal(len(fullNames), 2))
	assert.Assert(t, is.Equal(fullNames[0].Value().(string), "some_value"))
	assert.Assert(t, is.Equal(fullNames[1].Value().(string), "some_other_value"))
}

func createStringAttribute(name string, value []byte, anchors []*yotiprotoattr.Anchor, attributeID string) *yotiprotoattr.Attribute {
	return &yotiprotoattr.Attribute{
		Name:        name,
		Value:       value,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     anchors,
		EphemeralId: attributeID,
	}
}

func createSelfieAttribute(contentType yotiprotoattr.ContentType, attributeID string) *yotiprotoattr.Attribute {
	var attributeImage = &yotiprotoattr.Attribute{
		Name:        consts.AttrSelfie,
		Value:       attributeValue,
		ContentType: contentType,
		Anchors:     []*yotiprotoattr.Anchor{},
		EphemeralId: attributeID,
	}
	return attributeImage
}

func createDocumentImagesAttribute(attributeID string) *yotiprotoattr.Attribute {
	multiValue, err := proto.Marshal(&yotiprotoattr.MultiValue{})
	if err != nil {
		panic(err)
	}

	protoAttribute := &yotiprotoattr.Attribute{
		Name:        consts.AttrDocumentImages,
		Value:       multiValue,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     make([]*yotiprotoattr.Anchor, 0),
		EphemeralId: attributeID,
	}
	return protoAttribute
}
