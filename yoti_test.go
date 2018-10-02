package yoti

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/big"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/attribute"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

const token = "NpdmVVGC-28356678-c236-4518-9de4-7a93009ccaf0-c5f92f2a-5539-453e-babc-9b06e1d6b7de"
const encryptedToken = "b6H19bUCJhwh6WqQX_sEHWX9RP-A_ANr1fkApwA4Dp2nJQFAjrF9e6YCXhNBpAIhfHnN0iXubyXxXZMNwNMSQ5VOxkqiytrvPykfKQWHC6ypSbfy0ex8ihndaAXG5FUF-qcU8QaFPMy6iF3x0cxnY0Ij0kZj0Ng2t6oiNafb7AhT-VGXxbFbtZu1QF744PpWMuH0LVyBsAa5N5GJw2AyBrnOh67fWMFDKTJRziP5qCW2k4h5vJfiYr_EOiWKCB1d_zINmUm94ZffGXxcDAkq-KxhN1ZuNhGlJ2fKcFh7KxV0BqlUWPsIEiwS0r9CJ2o1VLbEs2U_hCEXaqseEV7L29EnNIinEPVbL4WR7vkF6zQCbK_cehlk2Qwda-VIATqupRO5grKZN78R9lBitvgilDaoE7JB_VFcPoljGQ48kX0wje1mviX4oJHhuO8GdFITS5LTbojGVQWT7LUNgAUe0W0j-FLHYYck3v84OhWTqads5_jmnnLkp9bdJSRuJF0e8pNdePnn2lgF-GIcyW_0kyGVqeXZrIoxnObLpF-YeUteRBKTkSGFcy7a_V_DLiJMPmH8UXDLOyv8TVt3ppzqpyUrLN2JVMbL5wZ4oriL2INEQKvw_boDJjZDGeRlu5m1y7vGDNBRDo64-uQM9fRUULPw-YkABNwC0DeShswzT00="

func TestYotiClient_KeyLoad_Failure(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key-invalid-format.pem")

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		result = &httpResponse{
			Success:    false,
			StatusCode: 500}
		return
	}

	_, _, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)

	if len(errorStrings) == 0 {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(errorStrings[0], "Invalid Key") {
		t.Errorf("expected outcome type starting with %q instead received %q", "Invalid Key", errorStrings[0])
	}
}

func TestYotiClient_HttpFailure_ReturnsFailure(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		result = &httpResponse{
			Success:    false,
			StatusCode: 500}
		return
	}

	_, _, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)
	if len(errorStrings) == 0 {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(errorStrings[0], ErrFailure.Error()) {
		t.Errorf("expected outcome type %q instead received %q", ErrFailure.Error(), errorStrings[0])
	}
}

func TestYotiClient_HttpFailure_ReturnsProfileNotFound(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		result = &httpResponse{
			Success:    false,
			StatusCode: 404}
		return
	}

	_, _, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)
	if len(errorStrings) == 0 {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(errorStrings[0], ErrProfileNotFound.Error()) {
		t.Errorf("expected outcome type %q instead received %q", ErrProfileNotFound.Error(), errorStrings[0])
	}
}

func TestYotiClient_SharingFailure_ReturnsFailure(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		result = &httpResponse{
			Success:    true,
			StatusCode: 200,
			Content:    `{"session_data":"session_data","receipt":{"receipt_id": null,"other_party_profile_content": null,"policy_uri":null,"personal_key":null,"remember_me_id":null, "sharing_outcome":"FAILURE","timestamp":"2016-09-23T13:04:11Z"}}`}
		return
	}

	_, _, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)
	if len(errorStrings) == 0 {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(errorStrings[0], ErrSharingFailure.Error()) {
		t.Errorf("expected outcome type %q instead received %q", ErrSharingFailure.Error(), errorStrings[0])
	}
}

func TestYotiClient_TokenDecodedSuccessfully(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	expectedAbsoluteURL := "/api/v1/profile/" + token

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		var theURL *url.URL
		var theError error
		if theURL, theError = url.Parse(uri); err != nil {
			t.Errorf("Yoti api did not generate a valid uri. instead it generated: %s", theError)
		}

		if theURL.Path != expectedAbsoluteURL {
			t.Errorf("Yoti api did not generate a url path. expected %s, generated: %s", expectedAbsoluteURL, theURL.Path)
		}

		result = &httpResponse{
			Success:    false,
			StatusCode: 500}
		return
	}

	_, _, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)
	if len(errorStrings) == 0 {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(errorStrings[0], ErrFailure.Error()) {
		t.Errorf("expected outcome type %q instead received %q", ErrFailure.Error(), errorStrings[0])
	}
}

func TestYotiClient_ParseProfile_Success(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	wrappedReceiptKey := "kyHPjq2+Y48cx+9yS/XzmW09jVUylSdhbP+3Q9Tc9p6bCEnyfa8vj38AIu744RzzE+Dc4qkSF21VfzQKtJVILfOXu5xRc7MYa5k3zWhjiesg/gsrv7J4wDyyBpHIJB8TWXnubYMbSYQJjlsfwyxE9kGe0YI08pRo2Tiht0bfR5Z/YrhAk4UBvjp84D+oyug/1mtGhKphA4vgPhQ9/y2wcInYxju7Q6yzOsXGaRUXR38Tn2YmY9OBgjxiTnhoYJFP1X9YJkHeWMW0vxF1RHxgIVrpf7oRzdY1nq28qzRg5+wC7cjRpS2i/CKUAo0oVG4pbpXsaFhaTewStVC7UFtA77JHb3EnF4HcSWMnK5FM7GGkL9MMXQenh11NZHKPWXpux0nLZ6/vwffXZfsiyTIcFL/NajGN8C/hnNBljoQ+B3fzWbjcq5ueUOPwARZ1y38W83UwMynzkud/iEdHLaZIu4qUCRkfSxJg7Dc+O9/BdiffkOn2GyFmNjVeq754DCUypxzMkjYxokedN84nK13OU4afVyC7t5DDxAK/MqAc69NCBRLqMi5f8BMeOZfMcSWPGC9a2Qu8VgG125TuZT4+wIykUhGyj3Bb2/fdPsxwuKFR+E0uqs0ZKvcv1tkNRRtKYBqTacgGK9Yoehg12cyLrITLdjU1fmIDn4/vrhztN5w="
	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	rememberMeID := "remember_me_id0123456789"

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		result = &httpResponse{
			Success:    true,
			StatusCode: 200,
			Content:    `{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`}
		return
	}

	userProfile, activityDetails, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)

	if errorStrings != nil {
		t.Error(errorStrings)
	}

	if userProfile.ID != rememberMeID {
		t.Errorf("expected id %q instead received %q", rememberMeID, userProfile.ID)
	}

	if userProfile.Selfie == nil {
		t.Error(`expected selfie attribute, but it was not present in the returned userProfile`)
	} else if string(userProfile.Selfie.Data) != "selfie0123456789" {
		t.Errorf("expected selfie attribute %q, instead received %q", "selfie0123456789", string(userProfile.Selfie.Data))
	}

	if userProfile.MobileNumber != "phone_number0123456789" {
		t.Errorf("expected mobileNumber value %q, instead received %q", "phone_number0123456789", userProfile.MobileNumber)
	}

	dobUserProfile := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)
	if userProfile.DateOfBirth == nil {
		t.Error(`expected date of birth but it was not present in the returned userProfile`)
	} else if !userProfile.DateOfBirth.Equal(dobUserProfile) {
		t.Errorf("expected date of birth %q, instead received %q", userProfile.DateOfBirth.Format(time.UnixDate), dobUserProfile.Format(time.UnixDate))
	}

	profile := activityDetails.UserProfile

	if activityDetails.RememberMeID != rememberMeID {
		t.Errorf("expected id %q, instead received %q", rememberMeID, activityDetails.RememberMeID)
	}

	expectedSelfieValue := "selfie0123456789"
	if profile.Selfie() == nil {
		t.Error(`expected selfie attribute, but it was not present in the returned profile`)
	} else if !cmp.Equal(profile.Selfie().Value, []byte(expectedSelfieValue)) {
		t.Errorf("expected selfie %q, instead received %q", expectedSelfieValue, string(profile.Selfie().Value))
	}

	if !cmp.Equal(profile.MobileNumber().Value, "phone_number0123456789") {
		t.Errorf("expected mobileNumber %q, instead received %q", "phone_number0123456789", profile.MobileNumber().Value)
	}

	expectedDoB := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)
	actualDoB := profile.DateOfBirth()

	if actualDoB == nil {
		t.Error(`expected date of birth, but it was not present in the returned profile`)
	} else if !actualDoB.Value.Equal(expectedDoB) {
		t.Errorf("expected date of birth: %q, instead received: %q", expectedDoB.Format(time.UnixDate), actualDoB.Value.Format(time.UnixDate))
	}
}

func TestYotiClient_ParseWithoutProfile_Success(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	wrappedReceiptKey := "kyHPjq2+Y48cx+9yS/XzmW09jVUylSdhbP+3Q9Tc9p6bCEnyfa8vj38AIu744RzzE+Dc4qkSF21VfzQKtJVILfOXu5xRc7MYa5k3zWhjiesg/gsrv7J4wDyyBpHIJB8TWXnubYMbSYQJjlsfwyxE9kGe0YI08pRo2Tiht0bfR5Z/YrhAk4UBvjp84D+oyug/1mtGhKphA4vgPhQ9/y2wcInYxju7Q6yzOsXGaRUXR38Tn2YmY9OBgjxiTnhoYJFP1X9YJkHeWMW0vxF1RHxgIVrpf7oRzdY1nq28qzRg5+wC7cjRpS2i/CKUAo0oVG4pbpXsaFhaTewStVC7UFtA77JHb3EnF4HcSWMnK5FM7GGkL9MMXQenh11NZHKPWXpux0nLZ6/vwffXZfsiyTIcFL/NajGN8C/hnNBljoQ+B3fzWbjcq5ueUOPwARZ1y38W83UwMynzkud/iEdHLaZIu4qUCRkfSxJg7Dc+O9/BdiffkOn2GyFmNjVeq754DCUypxzMkjYxokedN84nK13OU4afVyC7t5DDxAK/MqAc69NCBRLqMi5f8BMeOZfMcSWPGC9a2Qu8VgG125TuZT4+wIykUhGyj3Bb2/fdPsxwuKFR+E0uqs0ZKvcv1tkNRRtKYBqTacgGK9Yoehg12cyLrITLdjU1fmIDn4/vrhztN5w="
	rememberMeID := "remember_me_id0123456789"

	var otherPartyProfileContents = []string{
		`"other_party_profile_content": null,`,
		`"other_party_profile_content": "",`,
		``}

	for _, otherPartyProfileContent := range otherPartyProfileContents {

		var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
			result = &httpResponse{
				Success:    true,
				StatusCode: 200,
				Content:    `{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `",` + otherPartyProfileContent + `"remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`}
			return
		}

		userProfile, activityDetails, err := getActivityDetails(requester, encryptedToken, sdkID, key)

		if err != nil {
			t.Error(err)
		}

		if userProfile.ID != rememberMeID {
			t.Errorf("expected id %q instead received %q", rememberMeID, userProfile.ID)
		}

		if activityDetails.RememberMeID != rememberMeID {
			t.Errorf("expected id %q instead received %q", rememberMeID, activityDetails.RememberMeID)
		}
	}
}

func TestYotiClient_UnsupportedHttpMethod_ReturnsError(t *testing.T) {
	uri := "http://www.url.com"
	headers := CreateHeaders()
	httpRequestMethod := "UNSUPPORTEDMETHOD"
	contentBytes := make([]byte, 0)

	_, err := doRequest(uri, headers, httpRequestMethod, contentBytes)

	if err == nil {
		t.Error("Expected failure")
	}
}

func TestYotiClient_SupportedHttpMethod(t *testing.T) {
	uri := "http://www.url.com"
	headers := CreateHeaders()
	httpRequestMethod := HTTPMethodGet
	contentBytes := make([]byte, 0)

	_, err := doRequest(uri, headers, httpRequestMethod, contentBytes)

	if err != nil {
		t.Error(err)
	}
}

func TestYotiClient_PerformAmlCheck_Success(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {

		result = &httpResponse{
			Success:    true,
			StatusCode: 200,
			Content:    `{"on_fraud_list":true,"on_pep_list":true,"on_watch_list":true}`}
		return
	}

	result, err := performAmlCheck(
		createStandardAmlProfile(),
		requester,
		sdkID,
		key)

	if err != nil {
		t.Error(err)
	}

	if !result.OnFraudList {
		t.Errorf("'OnFraudList' value is expected to be true")
	}
	if !result.OnPEPList {
		t.Errorf("'OnPEPList' value is expected to be true")
	}
	if !result.OnWatchList {
		t.Errorf("'OnWatchList' value is expected to be true")
	}
}

func TestYotiClient_PerformAmlCheck_Unsuccessful(t *testing.T) {
	sdkID := "fake-sdk-id"
	key, _ := ioutil.ReadFile("test-key.pem")

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {

		result = &httpResponse{
			Success:    false,
			StatusCode: 503,
			Content:    `SERVICE UNAVAILABLE - Unable to reach the Integrity Service`}
		return
	}

	_, err := performAmlCheck(
		createStandardAmlProfile(),
		requester,
		sdkID,
		key)

	if err == nil {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(err.Error(), "AML Check was unsuccessful") {
		t.Errorf("expected outcome type starting with %q instead received %q", "AML Check was unsuccessful", err.Error())
	}
}

func TestYotiClient_ParseIsAgeVerifiedValue_True(t *testing.T) {
	trueValue := []byte("true")

	isAgeVerified, err := parseIsAgeVerifiedValue(trueValue)

	if err != nil {
		t.Errorf("Failed to parse IsAgeVerified value, error was %q", err.Error())
	}

	if !*isAgeVerified {
		t.Error("Expected true")
	}
}

func TestYotiClient_ParseIsAgeVerifiedValue_False(t *testing.T) {
	falseValue := []byte("false")

	isAgeVerified, err := parseIsAgeVerifiedValue(falseValue)

	if err != nil {
		t.Errorf("Failed to parse IsAgeVerified value, error was %q", err.Error())
	}

	if *isAgeVerified {
		t.Error("Expected false")
	}
}
func TestYotiClient_ParseIsAgeVerifiedValue_InvalidValueThrowsError(t *testing.T) {
	invalidValue := []byte("invalidBool")

	_, err := parseIsAgeVerifiedValue(invalidValue)

	if err == nil {
		t.Error("Expected error")
	}
}
func TestYotiClient_UnmarshallJSONValue_InvalidValueThrowsError(t *testing.T) {
	invalidStructuredAddress := []byte("invalidBool")

	_, err := attribute.UnmarshallJSON(invalidStructuredAddress)

	if err == nil {
		t.Error("Expected error")
	}
}

func TestYotiClient_UnmarshallJSONValue_ValidValue(t *testing.T) {
	const countryIso = "IND"
	const nestedValue = "NestedValue"

	var structuredAddress = []byte(`[
	{
		"address_format": 2,
		"building": "House No.86-A",		
		"state": "Punjab",
		"postal_code": "141012",
		"country_iso": "` + countryIso + `",
		"country": "India",
		"formatted_address": "House No.86-A\nRajgura Nagar\nLudhina\nPunjab\n141012\nIndia",
		"1":
		{
			"1-1":
			{
			  "1-1-1": "` + nestedValue + `"
			}
		}
	}
	]`)

	parsedStructuredAddress, err := attribute.UnmarshallJSON(structuredAddress)

	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	parsedStructuredAddressInterfaceArray := parsedStructuredAddress.([]interface{})

	parsedStructuredAddressMap := parsedStructuredAddressInterfaceArray[0].(map[string]interface{})
	actualCountryIso := parsedStructuredAddressMap["country_iso"]

	if countryIso != actualCountryIso {
		t.Errorf("expected country_iso: %q, actual value was: %q", countryIso, actualCountryIso)
	}
}

func TestYotiClient_MissingPostalAddress_UsesFormattedAddress(t *testing.T) {
	var formattedAddressText = `House No.86-A\nRajgura Nagar\nLudhina\nPunjab\n141012\nIndia`

	var structuredAddressBytes = []byte(`[
	{
		"address_format": 2,
		"building": "House No.86-A",
		"formatted_address": "` + formattedAddressText + `"
	}
	]`)

	structuredAddress, err := attribute.UnmarshallJSON(structuredAddressBytes)
	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var userProfile = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 ""}

	var jsonAttribute = &attribute.Attribute{
		Name:    attrConstStructuredPostalAddress,
		Value:   structuredAddressBytes,
		Type:    attribute.AttrTypeJSON,
		Anchors: []*anchor.Anchor{},
	}

	profile := createProfileWithSingleAttribute(jsonAttribute)

	profileAddress, profileErr := ensureAddressProfile(profile)
	if profileErr != nil {
		t.Errorf("Failed to add formatted address to address on Profile, error was %q", err.Error())
	}

	userProfileAddress, userProfileErr := ensureAddressUserProfile(userProfile)
	if userProfileErr != nil {
		t.Errorf("Failed to add formatted address to address on UserProfile, error was %q", err.Error())
	}

	escapedFormattedAddressText := strings.Replace(formattedAddressText, `\n`, "\n", -1)

	if profileAddress != escapedFormattedAddressText {
		t.Errorf("Address does not equal the expected formatted address. address: %q, formatted address: %q", profileAddress, formattedAddressText)
	}

	if userProfileAddress != escapedFormattedAddressText {
		t.Errorf("Address does not equal the expected formatted address. address: %q, formatted address: %q", userProfileAddress, formattedAddressText)
	}

	if !cmp.Equal(profile.StructuredPostalAddress().Anchors, []*anchor.Anchor{}) {
		t.Errorf("Retrieved attribute does not have the correct anchors. Expected %v, actual: %v", []anchor.Anchor{}, profile.StructuredPostalAddress().Anchors)
	}

	if !cmp.Equal(profile.StructuredPostalAddress().Type, attribute.AttrTypeJSON) {
		t.Errorf("Retrieved attribute does not have the correct type. Expected %q, actual: %q", attribute.AttrTypeJSON, profile.StructuredPostalAddress().Type)
	}
}

func TestYotiClient_PresentPostalAddress_DoesntUseFormattedAddress(t *testing.T) {
	var addressText = `PostalAddress`

	var structuredAddressBytes = []byte(`[
	{
		"address_format": 2,
		"building": "House No.86-A",
		"formatted_address": "FormattedAddress"
	}
	]`)
	structuredAddress, err := attribute.UnmarshallJSON(structuredAddressBytes)

	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var result = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 addressText}

	newFormattedAddress, err := ensureAddressUserProfile(result)

	if err != nil {
		t.Errorf("Failure when getting formatted address, error was %q", err.Error())
	}

	if newFormattedAddress != "" {
		t.Errorf("Address should be unchanged when it is present, but it is : %q", newFormattedAddress)
	}
}

func TestYotiClient_MissingFormattedAddress_AddressUnchanged(t *testing.T) {
	var structuredAddressBytes = []byte(`[
	{
		"address_format": 2,
		"building": "House No.86-A"
	}
	]`)

	structuredAddress, err := attribute.UnmarshallJSON(structuredAddressBytes)

	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var result = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 ""}

	address, err := ensureAddressUserProfile(result)

	if err != nil {
		t.Errorf("Failed to add formatted address to address, error was %q", err.Error())
	}

	if address != "" {
		t.Errorf("Formatted address missing, but address was still changed to: %q", address)
	}
}

func TestProfile_GetAttribute_RetrievesAttribute(t *testing.T) {
	attributeName := "test_attribute_name"
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &attribute.Attribute{
		Name:    attributeName,
		Value:   attributeValue,
		Type:    attribute.AttrTypeString,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	if att.Name != attributeName {
		t.Errorf("Retrieved attribute does not have the correct name. Expected %q, actual: %q", attributeName, att.Name)
	}

	if !cmp.Equal(att.Value, attributeValueString) {
		t.Errorf("Retrieved attribute does not have the correct value. Expected %q, actual: %q", attributeValue, att.Value)
	}
}

func TestProfile_StringAttribute(t *testing.T) {
	attributeName := attrConstNationality
	attributeValueString := "value"
	attributeValueBytes := []byte(attributeValueString)

	var as = &attribute.Attribute{
		Name:    attributeName,
		Value:   attributeValueBytes,
		Type:    attribute.AttrTypeString,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(as)

	if result.Nationality().Value != attributeValueString {
		t.Errorf("Retrieved attribute does not have the correct value. Expected %q, actual: %q", attributeValueString, result.Nationality().Value)
	}

	if !cmp.Equal(result.Nationality().Anchors, []*anchor.Anchor{}) {
		t.Errorf("Retrieved attribute does not have the correct anchors. Expected %v, actual: %v", []anchor.Anchor{}, result.Nationality().Anchors)
	}

	if !cmp.Equal(result.Nationality().Type, attribute.AttrTypeString) {
		t.Errorf("Retrieved attribute does not have the correct type. Expected %q, actual: %q", attribute.AttrTypeString, result.Nationality().Type)
	}
}

func TestProfile_GenericAttribute(t *testing.T) {
	attributeName := "genericAttr"
	attributeValueString := "value"
	attributeValueBytes := []byte(attributeValueString)

	var ag = &attribute.Attribute{
		Name:    attributeName,
		Value:   attributeValueBytes,
		Type:    attribute.AttrTypeInterface,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(ag)

	genericAttr := result.GetAttribute(attributeName)

	if !cmp.Equal(genericAttr.Value, attributeValueString) {
		t.Errorf("Retrieved attribute does not have the correct value. Expected %q, actual: %q", attributeValueBytes, genericAttr.Value)
	}

	if !cmp.Equal(genericAttr.Anchors, []*anchor.Anchor{}) {
		t.Errorf("Retrieved attribute does not have the correct anchors. Expected %v, actual: %v", []anchor.Anchor{}, genericAttr.Anchors)
	}

	if !cmp.Equal(genericAttr.Name, attributeName) {
		t.Errorf("Retrieved attribute does not have the correct name. Expected %q, actual: %q", attributeName, genericAttr.Name)
	}

	if !cmp.Equal(genericAttr.Type, attribute.AttrTypeInterface) {
		t.Errorf("Retrieved attribute does not have the correct type. Expected %q, actual: %q", attribute.AttrTypeInterface, genericAttr.Type)
	}
}

func TestProfile_AttributeProperty_RetrievesAttribute(t *testing.T) {
	attributeName := attrConstSelfie
	attributeValue := []byte("value")

	var attributeImage = &attribute.Attribute{
		Name:    attributeName,
		Value:   attributeValue,
		Type:    attribute.AttrTypePNG,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if selfie.Name != attributeName {
		t.Errorf("Retrieved attribute does not have the correct name. Expected %q, actual: %q", attributeName, selfie.Name)
	}

	if !reflect.DeepEqual(selfie.Value, attributeValue) {
		t.Errorf("Retrieved attribute does not have the correct value. Expected %q, actual: %q", attributeValue, selfie.Value)
	}

	if !cmp.Equal(selfie.Type, attribute.AttrTypePNG) {
		t.Errorf("Retrieved attribute does not have the correct type. Expected %q, actual: %q", attribute.AttrTypePNG, selfie.Type)
	}

	if !cmp.Equal(selfie.Anchors, []*anchor.Anchor{}) {
		t.Errorf("Retrieved attribute does not have the correct anchors. Expected %v, actual: %v", []anchor.Anchor{}, selfie.Anchors)
	}
}

func TestAttributeImage_Image_Png(t *testing.T) {
	attributeName := attrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &attribute.Attribute{
		Name:    attributeName,
		Value:   byteValue,
		Type:    attribute.AttrTypePNG,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if !cmp.Equal(selfie.Value, byteValue) {
		t.Errorf("Retrieved attribute does not have the correct Image. Expected %v, actual: %v", byteValue, selfie.Value)
	}
}

func TestAttributeImage_Image_Jpeg(t *testing.T) {
	attributeName := attrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &attribute.Attribute{
		Name:    attributeName,
		Value:   byteValue,
		Type:    attribute.AttrTypeJPEG,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if !cmp.Equal(selfie.Value, byteValue) {
		t.Errorf("Retrieved attribute does not have the correct byte value. Expected %v, actual: %v", byteValue, selfie.Value)
	}
}

func TestAttributeImage_Image_Default(t *testing.T) {
	attributeName := attrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &attribute.Attribute{
		Name:    attributeName,
		Value:   byteValue,
		Type:    attribute.AttrTypePNG,
		Anchors: []*anchor.Anchor{},
	}
	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if !cmp.Equal(selfie.Value, byteValue) {
		t.Errorf("Retrieved attribute does not have the correct byte value. Expected %v, actual: %v", byteValue, selfie.Value)
	}
}
func TestAttributeImage_Base64Selfie_Png(t *testing.T) {
	attributeName := attrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &attribute.Attribute{
		Name:    attributeName,
		Value:   imageBytes,
		Type:    attribute.AttrTypePNG,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/png;base64;," + base64ImageExpectedValue

	base64Selfie, err := result.Selfie().Base64URL()

	if err != nil {
		t.Error(err)
	}

	if base64Selfie != expectedBase64Selfie {
		t.Errorf("Base64Selfie does not have the correct value. Expected %q, actual: %q", expectedBase64Selfie, base64Selfie)
	}
}

func TestAttributeImage_Base64URL_Jpeg(t *testing.T) {
	attributeName := attrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &attribute.Attribute{
		Name:    attributeName,
		Value:   imageBytes,
		Type:    attribute.AttrTypeJPEG,
		Anchors: []*anchor.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/jpeg;base64;," + base64ImageExpectedValue

	base64Selfie, err := result.Selfie().Base64URL()

	if err != nil {
		t.Error(err)
	}

	if base64Selfie != expectedBase64Selfie {
		t.Errorf("Base64Selfie does not have the correct value. Expected %q, actual: %q", expectedBase64Selfie, base64Selfie)
	}
}

func TestProfile_GetAttribute_ReturnsNil(t *testing.T) {
	result := Profile{
		AttributeSlice: []*attribute.Attribute{},
	}

	attribute := result.GetAttribute("attributeName")

	if attribute != nil {
		t.Error("Attribute should not be retrieved if it is not present")
	}
}

func TestAnchorParser_Passport(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	anchorSlice := CreateAnchorSliceFromTestFile(t, "testanchorpassport.txt")

	var structuredAddressBytes = []byte(`[
		{
			"address_format": 2,
			"building": "House No.86-A"
		}
		]`)

	a := &attribute.Attribute{
		Name:    attrConstStructuredPostalAddress,
		Value:   structuredAddressBytes,
		Type:    attribute.AttrTypeJSON,
		Anchors: anchorSlice,
	}

	result := createProfileWithSingleAttribute(a)

	actualStructuredPostalAddress := result.StructuredPostalAddress()

	if actualStructuredPostalAddress.Err != nil {
		t.Error(actualStructuredPostalAddress.Err)
	}

	actualAnchor := actualStructuredPostalAddress.Anchors[0]

	if actualAnchor.Type != anchor.AnchorTypeSource {
		t.Errorf("Parsed anchor type is incorrect. Expected: %q, actual: %q", anchor.AnchorTypeSource, actualAnchor.Type)
	}

	expectedDate := time.Date(2018, time.April, 12, 13, 14, 32, 0, time.UTC)
	actualDate := actualAnchor.SignedTimestamp().Timestamp.UTC()
	if actualDate != expectedDate {
		t.Errorf("Parsed anchor SignedTimestamp is incorrect. Expected: %q, actual: %q", expectedDate, actualDate)
	}

	expectedSubType := "OCR"
	if actualAnchor.SubType() != expectedSubType {
		t.Errorf("Parsed anchor SubType is incorrect. Expected: %q, actual: %q", expectedSubType, actualAnchor.SubType())
	}

	expectedValue := "PASSPORT"
	if actualAnchor.Value()[0] != expectedValue {
		t.Errorf("Parsed anchor Value is incorrect. Expected: %q, actual: %q", expectedValue, actualAnchor.Value()[0])
	}

	actualSerialNo := actualAnchor.OriginServerCerts()[0].SerialNumber
	AssertServerCertSerialNo(t, "277870515583559162487099305254898397834", actualSerialNo)
}

func TestAnchorParser_DrivingLicense(t *testing.T) {
	anchorSlice := CreateAnchorSliceFromTestFile(t, "testanchordrivinglicense.txt")

	attribute := &attribute.Attribute{
		Name:    attrConstGender,
		Value:   []byte("value"),
		Type:    attribute.AttrTypeString,
		Anchors: anchorSlice,
	}

	result := createProfileWithSingleAttribute(attribute)

	resultAnchor := result.Gender().Anchors[0]

	if resultAnchor.Type != anchor.AnchorTypeSource {
		t.Errorf("Parsed anchor type is incorrect. Expected: %q, actual: %q", anchor.AnchorTypeSource, resultAnchor.Type)
	}

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 3, 0, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp.UTC()
	if actualDate != expectedDate {
		t.Errorf("Parsed anchor SignedTimestamp is incorrect. Expected: %q, actual: %q", expectedDate, actualDate)
	}

	expectedSubType := ""
	if resultAnchor.SubType() != expectedSubType {
		t.Errorf("Parsed anchor SubType is incorrect. Expected: %q, actual: %q", expectedSubType, resultAnchor.SubType())
	}

	expectedValue := "DRIVING_LICENCE"
	if resultAnchor.Value()[0] != expectedValue {
		t.Errorf("Parsed anchor Value is incorrect. Expected: %q, actual: %q", expectedValue, resultAnchor.Value()[0])
	}

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	AssertServerCertSerialNo(t, "46131813624213904216516051554755262812", actualSerialNo)
}
func TestAnchorParser_YotiAdmin(t *testing.T) {
	anchorSlice := CreateAnchorSliceFromTestFile(t, "testanchoryotiadmin.txt")

	attr := &attribute.Attribute{
		Name:    attrConstDateOfBirth,
		Value:   []byte("1999-01-01"),
		Type:    attribute.AttrTypeTime,
		Anchors: anchorSlice,
	}

	result := createProfileWithSingleAttribute(attr)

	DoB := result.DateOfBirth()

	if DoB.Err != nil {
		t.Error(DoB.Err)
	}

	resultAnchor := DoB.Anchors[0]

	if resultAnchor.Type != anchor.AnchorTypeVerifier {
		t.Errorf("Parsed anchor type is incorrect. Expected: %q, actual: %q", anchor.AnchorTypeVerifier, resultAnchor.Type)
	}

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 4, 0, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp.UTC()
	if actualDate != expectedDate {
		t.Errorf("Parsed anchor SignedTimestamp is incorrect. Expected: %q, actual: %q", expectedDate, actualDate)
	}

	expectedSubType := ""
	if resultAnchor.SubType() != expectedSubType {
		t.Errorf("Parsed anchor SubType is incorrect. Expected: %q, actual: %q", expectedSubType, resultAnchor.SubType())
	}

	expectedValue := "YOTI_ADMIN"
	if resultAnchor.Value()[0] != expectedValue {
		t.Errorf("Parsed anchor Value is incorrect. Expected: %q, actual: %q", expectedValue, resultAnchor.Value()[0])
	}

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	AssertServerCertSerialNo(t, "256616937783084706710155170893983549581", actualSerialNo)
}

func createProfileWithSingleAttribute(attr *attribute.Attribute) Profile {
	var attributeSlice []*attribute.Attribute
	attributeSlice = append(attributeSlice, attr)

	return Profile{
		AttributeSlice: attributeSlice,
	}
}

func AssertServerCertSerialNo(t *testing.T, expectedSerialNo string, actualSerialNo *big.Int) {
	expectedSerialNoBigInt := new(big.Int)
	expectedSerialNoBigInt, ok := expectedSerialNoBigInt.SetString(expectedSerialNo, 10)
	if !ok {
		t.Error("Unexpected error when setting string as big int")
	}

	if expectedSerialNoBigInt.Cmp(actualSerialNo) != 0 { //0 == equivalent
		t.Errorf("Parsed anchor OriginServerCerts is incorrect. Expected: %q, actual: %q", expectedSerialNo, actualSerialNo)
	}
}

func CreateAnchorSliceFromTestFile(t *testing.T, filename string) []*anchor.Anchor {
	anchorBytes, err := DecodeTestFile(t, filename)

	if err != nil {
		t.Errorf("error decoding test file: %q", err)
	}

	protoAnchor := &yotiprotoattr_v3.Anchor{}
	if err := proto.Unmarshal(anchorBytes, protoAnchor); err != nil {
		t.Errorf("Error converting test anchor bytes into a Protobuf anchor. Error: %q", err)
	}

	protoAnchors := append([]*yotiprotoattr_v3.Anchor{}, protoAnchor)

	return anchor.ParseAnchors(protoAnchors)
}

func DecodeTestFile(t *testing.T, filename string) (result []byte, err error) {
	base64Bytes := readTestFile(t, filename)
	base64String := string(base64Bytes)
	anchorBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return anchorBytes, nil
}

func readTestFile(t *testing.T, filename string) (result []byte) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}

	return b
}

func CreateHeaders() (result map[string]string) {

	headers := make(map[string]string)

	headers["Header1"] = "test"

	return headers
}

func createStandardAmlProfile() (result AmlProfile) {
	var amlAddress = AmlAddress{
		Country: "GBR"}

	var amlProfile = AmlProfile{
		GivenNames: "Edward Richard George",
		FamilyName: "Heath",
		Address:    amlAddress}

	return amlProfile
}
