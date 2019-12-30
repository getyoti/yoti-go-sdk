package yoti

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

const (
	token             = "NpdmVVGC-28356678-c236-4518-9de4-7a93009ccaf0-c5f92f2a-5539-453e-babc-9b06e1d6b7de"
	encryptedToken    = "b6H19bUCJhwh6WqQX_sEHWX9RP-A_ANr1fkApwA4Dp2nJQFAjrF9e6YCXhNBpAIhfHnN0iXubyXxXZMNwNMSQ5VOxkqiytrvPykfKQWHC6ypSbfy0ex8ihndaAXG5FUF-qcU8QaFPMy6iF3x0cxnY0Ij0kZj0Ng2t6oiNafb7AhT-VGXxbFbtZu1QF744PpWMuH0LVyBsAa5N5GJw2AyBrnOh67fWMFDKTJRziP5qCW2k4h5vJfiYr_EOiWKCB1d_zINmUm94ZffGXxcDAkq-KxhN1ZuNhGlJ2fKcFh7KxV0BqlUWPsIEiwS0r9CJ2o1VLbEs2U_hCEXaqseEV7L29EnNIinEPVbL4WR7vkF6zQCbK_cehlk2Qwda-VIATqupRO5grKZN78R9lBitvgilDaoE7JB_VFcPoljGQ48kX0wje1mviX4oJHhuO8GdFITS5LTbojGVQWT7LUNgAUe0W0j-FLHYYck3v84OhWTqads5_jmnnLkp9bdJSRuJF0e8pNdePnn2lgF-GIcyW_0kyGVqeXZrIoxnObLpF-YeUteRBKTkSGFcy7a_V_DLiJMPmH8UXDLOyv8TVt3ppzqpyUrLN2JVMbL5wZ4oriL2INEQKvw_boDJjZDGeRlu5m1y7vGDNBRDo64-uQM9fRUULPw-YkABNwC0DeShswzT00="
	wrappedReceiptKey = "kyHPjq2+Y48cx+9yS/XzmW09jVUylSdhbP+3Q9Tc9p6bCEnyfa8vj38AIu744RzzE+Dc4qkSF21VfzQKtJVILfOXu5xRc7MYa5k3zWhjiesg/gsrv7J4wDyyBpHIJB8TWXnubYMbSYQJjlsfwyxE9kGe0YI08pRo2Tiht0bfR5Z/YrhAk4UBvjp84D+oyug/1mtGhKphA4vgPhQ9/y2wcInYxju7Q6yzOsXGaRUXR38Tn2YmY9OBgjxiTnhoYJFP1X9YJkHeWMW0vxF1RHxgIVrpf7oRzdY1nq28qzRg5+wC7cjRpS2i/CKUAo0oVG4pbpXsaFhaTewStVC7UFtA77JHb3EnF4HcSWMnK5FM7GGkL9MMXQenh11NZHKPWXpux0nLZ6/vwffXZfsiyTIcFL/NajGN8C/hnNBljoQ+B3fzWbjcq5ueUOPwARZ1y38W83UwMynzkud/iEdHLaZIu4qUCRkfSxJg7Dc+O9/BdiffkOn2GyFmNjVeq754DCUypxzMkjYxokedN84nK13OU4afVyC7t5DDxAK/MqAc69NCBRLqMi5f8BMeOZfMcSWPGC9a2Qu8VgG125TuZT4+wIykUhGyj3Bb2/fdPsxwuKFR+E0uqs0ZKvcv1tkNRRtKYBqTacgGK9Yoehg12cyLrITLdjU1fmIDn4/vrhztN5w="
	attributeName     = "test_attribute_name"
)

type mockHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (mock *mockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if mock.do != nil {
		return mock.do(request)
	}
	return nil, nil
}

func TestAttributeImage_Image_Png(t *testing.T) {
	attributeName := AttrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestAttributeImage_Image_Jpeg(t *testing.T) {
	attributeName := AttrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestAttributeImage_Image_Default(t *testing.T) {
	attributeName := AttrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}
func TestAttributeImage_Base64Selfie_Png(t *testing.T) {
	attributeName := AttrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/png;base64," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}

func TestAttributeImage_Base64URL_Jpeg(t *testing.T) {
	attributeName := AttrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/jpeg;base64," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}

func TestAnchorParser_Passport(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_passport.txt")

	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A"
	}`)

	a := &yotiprotoattr.Attribute{
		Name:        AttrConstStructuredPostalAddress,
		Value:       structuredAddressBytes,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(a)

	var actualStructuredPostalAddress *attribute.JSONAttribute

	actualStructuredPostalAddress, err := result.StructuredPostalAddress()

	assert.Assert(t, is.Nil(err))

	actualAnchor := actualStructuredPostalAddress.Anchors()[0]

	assert.Equal(t, actualAnchor, actualStructuredPostalAddress.Sources()[0], "Anchors and Sources should be the same when there is only one Source")
	assert.Equal(t, actualAnchor.Type(), anchor.AnchorTypeSource)

	expectedDate := time.Date(2018, time.April, 12, 13, 14, 32, 835537e3, time.UTC)
	actualDate := actualAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := "OCR"
	assert.Equal(t, actualAnchor.SubType(), expectedSubType)

	expectedValue := "PASSPORT"
	assert.Equal(t, actualAnchor.Value()[0], expectedValue)

	actualSerialNo := actualAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "277870515583559162487099305254898397834", actualSerialNo)
}

func TestAnchorParser_DrivingLicense(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_driving_license.txt")

	attribute := &yotiprotoattr.Attribute{
		Name:        AttrConstGender,
		Value:       []byte("value"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attribute)

	genderAttribute := result.Gender()
	resultAnchor := genderAttribute.Anchors()[0]

	assert.Equal(t, resultAnchor, genderAttribute.Sources()[0], "Anchors and Sources should be the same when there is only one Source")
	assert.Equal(t, resultAnchor.Type(), anchor.AnchorTypeSource)

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 3, 923537e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := ""
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)

	expectedValue := "DRIVING_LICENCE"
	assert.Equal(t, resultAnchor.Value()[0], expectedValue)

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "46131813624213904216516051554755262812", actualSerialNo)
}

func TestAnchorParser_UnknownAnchor(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_unknown.txt")

	attr := &yotiprotoattr.Attribute{
		Name:        AttrConstDateOfBirth,
		Value:       []byte("1999-01-01"),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attr)

	DoB, err := result.DateOfBirth()

	assert.Assert(t, is.Nil(err))
	resultAnchor := DoB.Anchors()[0]

	expectedDate := time.Date(2019, time.March, 5, 10, 45, 11, 840037e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := "TEST UNKNOWN SUB TYPE"
	expectedType := anchor.AnchorTypeUnknown
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)
	assert.Equal(t, resultAnchor.Type(), expectedType)
	assert.Equal(t, len(resultAnchor.Value()), 0)
}

func TestAnchorParser_YotiAdmin(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "fixtures/test_anchor_yoti_admin.txt")

	attr := &yotiprotoattr.Attribute{
		Name:        AttrConstDateOfBirth,
		Value:       []byte("1999-01-01"),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attr)

	DoB, err := result.DateOfBirth()

	assert.Assert(t, is.Nil(err))

	resultAnchor := DoB.Anchors()[0]

	assert.Equal(t, resultAnchor, DoB.Verifiers()[0])

	assert.Equal(t, resultAnchor.Type(), anchor.AnchorTypeVerifier)

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 4, 95238e3, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	assert.Equal(t, actualDate, expectedDate)

	expectedSubType := ""
	assert.Equal(t, resultAnchor.SubType(), expectedSubType)

	expectedValue := "YOTI_ADMIN"
	assert.Equal(t, resultAnchor.Value()[0], expectedValue)

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "256616937783084706710155170893983549581", actualSerialNo)
}

func TestAnchors_None(t *testing.T) {
	anchorSlice := []*anchor.Anchor{}

	sources := anchor.GetSources(anchorSlice)
	assert.Equal(t, len(sources), 0, "GetSources should not return anything with empty anchors")

	verifiers := anchor.GetVerifiers(anchorSlice)
	assert.Equal(t, len(verifiers), 0, "GetVerifiers should not return anything with empty anchors")
}

func TestDateOfBirthAttribute(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_date_of_birth.txt")

	dateOfBirthAttribute, err := attribute.NewTime(protoAttribute)

	assert.Assert(t, is.Nil(err))

	expectedDateOfBirth := time.Date(1970, time.December, 01, 0, 0, 0, 0, time.UTC)
	actualDateOfBirth := dateOfBirthAttribute.Value()

	assert.Assert(t, actualDateOfBirth.Equal(expectedDateOfBirth))
}

func TestNewImageSlice(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt")

	documentImagesAttribute, err := attribute.NewImageSlice(protoAttribute)

	assert.Assert(t, is.Nil(err))

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttribute.Value(), documentImagesAttribute.Anchors()[0])
}

func TestImageSliceNotCreatedWithNonMultiValueType(t *testing.T) {
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	_, err := attribute.NewImageSlice(attr)

	assert.Assert(t, err != nil, "Expected error when creating image slice from attribute which isn't of multi-value type")
}

func TestMultiValueNotCreatedWithNonMultiValueType(t *testing.T) {
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	_, err := attribute.NewMultiValue(attr)

	assert.Assert(t, err != nil, "Expected error when creating multi value from attribute which isn't of multi-value type")
}

func TestNewMultiValue(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt")

	multiValueAttribute, err := attribute.NewMultiValue(protoAttribute)

	assert.Assert(t, is.Nil(err))

	var documentImagesAttributeItems []*attribute.Image = attribute.CreateImageSlice(multiValueAttribute.Value())

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttributeItems, multiValueAttribute.Anchors()[0])
}

func TestInvalidMultiValueNotReturned(t *testing.T) {
	var invalidMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_DATE,
		Data:        []byte("invalid"),
	}

	var stringMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_STRING,
		Data:        []byte("string"),
	}

	var multiValueItemSlice = []*yotiprotoattr.MultiValue_Value{invalidMultiValueItem, stringMultiValueItem}

	var multiValueStruct = &yotiprotoattr.MultiValue{
		Values: multiValueItemSlice,
	}

	var marshalledMultiValueData = marshallMultiValue(t, multiValueStruct)
	attributeName := "nestedMultiValue"

	var protoAttribute = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       marshalledMultiValueData,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(protoAttribute)

	assert.Assert(t, is.Nil(profile.GetAttribute(attributeName)))
}

func TestNewGeneric_ShouldParseUnknownTypeAsString(t *testing.T) {
	value := []byte("value")
	parsed := attribute.NewGeneric(&yotiprotoattr.Attribute{
		ContentType: yotiprotoattr.ContentType_UNDEFINED,
		Value:       value,
	})

	stringValue, ok := parsed.Value().(string)
	assert.Check(t, ok)

	assert.Equal(t, stringValue, string(value))
}

func TestNestedMultiValue(t *testing.T) {
	var innerMultiValueProtoValue []byte = createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt").Value

	var stringMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_STRING,
		Data:        []byte("string"),
	}

	var multiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Data:        innerMultiValueProtoValue,
	}

	var multiValueItemSlice = []*yotiprotoattr.MultiValue_Value{stringMultiValueItem, multiValueItem}

	multiValueAttribute, err := createMultiValueAttribute(t, multiValueItemSlice)

	assert.Assert(t, is.Nil(err))

	for key, value := range multiValueAttribute.Value() {
		switch key {
		case 0:
			value0 := value.GetValue()

			assert.Equal(t, value0.(string), "string")
		case 1:
			value1 := value.GetValue()

			innerItems, ok := value1.([]*attribute.Item)
			assert.Assert(t, ok)

			for innerKey, item := range innerItems {
				switch innerKey {
				case 0:
					assertIsExpectedImage(t, item.GetValue().(*attribute.Image), "jpeg", "vWgD//2Q==")

				case 1:
					assertIsExpectedImage(t, item.GetValue().(*attribute.Image), "jpeg", "38TVEH/9k=")
				}
			}
		}
	}
}

func TestMultiValueGenericGetter(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt")
	profile := createProfileWithSingleAttribute(protoAttribute)

	multiValueAttribute := profile.GetAttribute(AttrConstDocumentImages)

	// We need to cast, since GetAttribute always returns generic attributes
	multiValueAttributeValue := multiValueAttribute.Value().([]*attribute.Item)
	imageSlice := attribute.CreateImageSlice(multiValueAttributeValue)

	assertIsExpectedDocumentImagesAttribute(t, imageSlice, multiValueAttribute.Anchors()[0])
}

func TestNewThirdPartyAttribute(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_third_party.txt")

	stringAttribute := attribute.NewString(protoAttribute)

	assert.Equal(t, stringAttribute.Value(), "test-third-party-attribute-0")
	assert.Equal(t, stringAttribute.Name(), "com.thirdparty.id")

	assert.Equal(t, stringAttribute.Sources()[0].Value()[0], "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Sources()[0].SubType(), "orgName")

	assert.Equal(t, stringAttribute.Verifiers()[0].Value()[0], "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Verifiers()[0].SubType(), "orgName")
}

func assertIsExpectedDocumentImagesAttribute(t *testing.T, actualDocumentImages []*attribute.Image, anchor *anchor.Anchor) {

	assert.Equal(t, len(actualDocumentImages), 2, "This Document Images attribute should have two images")

	assertIsExpectedImage(t, actualDocumentImages[0], "jpeg", "vWgD//2Q==")
	assertIsExpectedImage(t, actualDocumentImages[1], "jpeg", "38TVEH/9k=")

	expectedValue := "NATIONAL_ID"
	assert.Equal(t, anchor.Value()[0], expectedValue)

	expectedSubType := "STATE_ID"
	assert.Equal(t, anchor.SubType(), expectedSubType)
}

func assertIsExpectedImage(t *testing.T, image *attribute.Image, imageType string, expectedBase64URLLast10 string) {
	assert.Equal(t, image.Type, imageType)

	actualBase64URL := image.Base64URL()

	ActualBase64URLLast10Chars := actualBase64URL[len(actualBase64URL)-10:]

	assert.Equal(t, ActualBase64URLLast10Chars, expectedBase64URLLast10)
}

func marshallMultiValue(t *testing.T, multiValue *yotiprotoattr.MultiValue) []byte {
	marshalled, err := proto.Marshal(multiValue)

	assert.Assert(t, is.Nil(err))

	return marshalled
}

func assertServerCertSerialNo(t *testing.T, expectedSerialNo string, actualSerialNo *big.Int) {
	expectedSerialNoBigInt := new(big.Int)
	expectedSerialNoBigInt, ok := expectedSerialNoBigInt.SetString(expectedSerialNo, 10)
	assert.Assert(t, ok, "Unexpected error when setting string as big int")

	assert.Equal(t, expectedSerialNoBigInt.Cmp(actualSerialNo), 0) //0 == equivalent
}

func createMultiValueAttribute(t *testing.T, multiValueItemSlice []*yotiprotoattr.MultiValue_Value) (*attribute.MultiValueAttribute, error) {
	var multiValueStruct = &yotiprotoattr.MultiValue{
		Values: multiValueItemSlice,
	}

	var marshalledMultiValueData = marshallMultiValue(t, multiValueStruct)
	attributeName := "nestedMultiValue"

	var protoAttribute = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       marshalledMultiValueData,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	return attribute.NewMultiValue(protoAttribute)
}

func createAttributeFromTestFile(t *testing.T, filename string) *yotiprotoattr.Attribute {
	attributeBytes := decodeTestFile(t, filename)

	attributeStruct := &yotiprotoattr.Attribute{}

	err2 := proto.Unmarshal(attributeBytes, attributeStruct)

	assert.Assert(t, is.Nil(err2))

	return attributeStruct
}

func createAnchorSliceFromTestFile(t *testing.T, filename string) []*yotiprotoattr.Anchor {
	anchorBytes := decodeTestFile(t, filename)

	protoAnchor := &yotiprotoattr.Anchor{}
	err2 := proto.Unmarshal(anchorBytes, protoAnchor)
	assert.Assert(t, is.Nil(err2))

	protoAnchors := append([]*yotiprotoattr.Anchor{}, protoAnchor)

	return protoAnchors
}

func decodeTestFile(t *testing.T, filename string) (result []byte) {
	base64Bytes := readTestFile(t, filename)
	base64String := string(base64Bytes)
	filebytes, err := base64.StdEncoding.DecodeString(base64String)

	assert.Assert(t, is.Nil(err))

	return filebytes
}

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

func readTestFile(t *testing.T, filename string) (result []byte) {
	b, err := ioutil.ReadFile(filename)
	assert.Assert(t, is.Nil(err))

	return b
}
