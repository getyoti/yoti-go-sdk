package yoti

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/big"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/attribute"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

const (
	token             = "NpdmVVGC-28356678-c236-4518-9de4-7a93009ccaf0-c5f92f2a-5539-453e-babc-9b06e1d6b7de"
	encryptedToken    = "b6H19bUCJhwh6WqQX_sEHWX9RP-A_ANr1fkApwA4Dp2nJQFAjrF9e6YCXhNBpAIhfHnN0iXubyXxXZMNwNMSQ5VOxkqiytrvPykfKQWHC6ypSbfy0ex8ihndaAXG5FUF-qcU8QaFPMy6iF3x0cxnY0Ij0kZj0Ng2t6oiNafb7AhT-VGXxbFbtZu1QF744PpWMuH0LVyBsAa5N5GJw2AyBrnOh67fWMFDKTJRziP5qCW2k4h5vJfiYr_EOiWKCB1d_zINmUm94ZffGXxcDAkq-KxhN1ZuNhGlJ2fKcFh7KxV0BqlUWPsIEiwS0r9CJ2o1VLbEs2U_hCEXaqseEV7L29EnNIinEPVbL4WR7vkF6zQCbK_cehlk2Qwda-VIATqupRO5grKZN78R9lBitvgilDaoE7JB_VFcPoljGQ48kX0wje1mviX4oJHhuO8GdFITS5LTbojGVQWT7LUNgAUe0W0j-FLHYYck3v84OhWTqads5_jmnnLkp9bdJSRuJF0e8pNdePnn2lgF-GIcyW_0kyGVqeXZrIoxnObLpF-YeUteRBKTkSGFcy7a_V_DLiJMPmH8UXDLOyv8TVt3ppzqpyUrLN2JVMbL5wZ4oriL2INEQKvw_boDJjZDGeRlu5m1y7vGDNBRDo64-uQM9fRUULPw-YkABNwC0DeShswzT00="
	sdkID             = "fake-sdk-id"
	wrappedReceiptKey = "kyHPjq2+Y48cx+9yS/XzmW09jVUylSdhbP+3Q9Tc9p6bCEnyfa8vj38AIu744RzzE+Dc4qkSF21VfzQKtJVILfOXu5xRc7MYa5k3zWhjiesg/gsrv7J4wDyyBpHIJB8TWXnubYMbSYQJjlsfwyxE9kGe0YI08pRo2Tiht0bfR5Z/YrhAk4UBvjp84D+oyug/1mtGhKphA4vgPhQ9/y2wcInYxju7Q6yzOsXGaRUXR38Tn2YmY9OBgjxiTnhoYJFP1X9YJkHeWMW0vxF1RHxgIVrpf7oRzdY1nq28qzRg5+wC7cjRpS2i/CKUAo0oVG4pbpXsaFhaTewStVC7UFtA77JHb3EnF4HcSWMnK5FM7GGkL9MMXQenh11NZHKPWXpux0nLZ6/vwffXZfsiyTIcFL/NajGN8C/hnNBljoQ+B3fzWbjcq5ueUOPwARZ1y38W83UwMynzkud/iEdHLaZIu4qUCRkfSxJg7Dc+O9/BdiffkOn2GyFmNjVeq754DCUypxzMkjYxokedN84nK13OU4afVyC7t5DDxAK/MqAc69NCBRLqMi5f8BMeOZfMcSWPGC9a2Qu8VgG125TuZT4+wIykUhGyj3Bb2/fdPsxwuKFR+E0uqs0ZKvcv1tkNRRtKYBqTacgGK9Yoehg12cyLrITLdjU1fmIDn4/vrhztN5w="
)

func TestYotiClient_KeyLoad_Failure(t *testing.T) {
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
	key, _ := ioutil.ReadFile("test-key.pem")

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

	if activityDetails.RememberMeID() != rememberMeID {
		t.Errorf("expected id %q, instead received %q", rememberMeID, activityDetails.RememberMeID())
	}

	expectedSelfieValue := "selfie0123456789"
	if profile.Selfie() == nil {
		t.Error(`expected selfie attribute, but it was not present in the returned profile`)
	} else if !cmp.Equal(profile.Selfie().Value().Data, []byte(expectedSelfieValue)) {
		t.Errorf("expected selfie %q, instead received %q", expectedSelfieValue, string(profile.Selfie().Value().Data))
	}

	if !cmp.Equal(profile.MobileNumber().Value(), "phone_number0123456789") {
		t.Errorf("expected mobileNumber %q, instead received %q", "phone_number0123456789", profile.MobileNumber().Value())
	}

	expectedDoB := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)

	actualDoB, err := profile.DateOfBirth()
	if err != nil {
		t.Error(err)
	}

	if actualDoB == nil {
		t.Error(`expected date of birth, but it was not present in the returned profile`)
	} else if !actualDoB.Value().Equal(expectedDoB) {
		t.Errorf("expected date of birth: %q, instead received: %q", expectedDoB.Format(time.UnixDate), actualDoB.Value().Format(time.UnixDate))
	}
}

func TestYotiClient_ParentRememberMeID(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")
	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	parentRememberMeID := "parent_remember_me_id0123456789"

	var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
		result = &httpResponse{
			Success:    true,
			StatusCode: 200,
			Content: `{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey +
				`","other_party_profile_content": "` + otherPartyProfileContent +
				`","parent_remember_me_id":"` + parentRememberMeID + `", "sharing_outcome":"SUCCESS"}}`}
		return
	}

	_, activityDetails, errorStrings := getActivityDetails(requester, encryptedToken, sdkID, key)

	if errorStrings != nil {
		t.Error(errorStrings)
	}

	if activityDetails.ParentRememberMeID() != parentRememberMeID {
		t.Errorf("expected id %q, instead received %q", parentRememberMeID, activityDetails.ParentRememberMeID())
	}
}
func TestYotiClient_ParseWithoutProfile_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")
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
				Content: `{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `",` +
					otherPartyProfileContent + `"remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`}
			return
		}

		userProfile, activityDetails, err := getActivityDetails(requester, encryptedToken, sdkID, key)

		if err != nil {
			t.Error(err)
		}

		if userProfile.ID != rememberMeID {
			t.Errorf("expected id %q instead received %q", rememberMeID, userProfile.ID)
		}

		if activityDetails.RememberMeID() != rememberMeID {
			t.Errorf("expected id %q instead received %q", rememberMeID, activityDetails.RememberMeID())
		}
	}
}

func TestYotiClient_ParseWithoutRememberMeID_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	var otherPartyProfileContents = []string{
		`"other_party_profile_content": null,`,
		`"other_party_profile_content": "",`}

	for _, otherPartyProfileContent := range otherPartyProfileContents {

		var requester = func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
			result = &httpResponse{
				Success:    true,
				StatusCode: 200,
				Content: `{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `",` +
					otherPartyProfileContent + `"sharing_outcome":"SUCCESS"}}`}
			return
		}

		_, _, err := getActivityDetails(requester, encryptedToken, sdkID, key)

		if err != nil {
			t.Error(err)
		}
	}
}

func TestYotiClient_UnsupportedHttpMethod_ReturnsError(t *testing.T) {
	uri := "http://www.url.com"
	headers := createTestHeaders()
	httpRequestMethod := "UNSUPPORTEDMETHOD"
	contentBytes := make([]byte, 0)

	_, err := doRequest(uri, headers, httpRequestMethod, contentBytes)

	if err == nil {
		t.Error("Expected failure")
	}
}

func TestYotiClient_SupportedHttpMethod(t *testing.T) {
	uri := "http://www.url.com"
	headers := createTestHeaders()
	httpRequestMethod := HTTPMethodGet
	contentBytes := make([]byte, 0)

	_, err := doRequest(uri, headers, httpRequestMethod, contentBytes)

	if err != nil {
		t.Error(err)
	}
}

func TestYotiClient_PerformAmlCheck_Success(t *testing.T) {
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

	var expectedErrString = "AML Check was unsuccessful"
	if err == nil {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(err.Error(), expectedErrString) {
		t.Errorf(
			"expected outcome type starting with %q instead received %q",
			expectedErrString,
			err.Error())
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
	const (
		countryIso  = "IND"
		nestedValue = "NestedValue"
	)

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

	parsedStructuredAddressInterfaceSlice := parsedStructuredAddress.([]interface{})

	parsedStructuredAddressMap := parsedStructuredAddressInterfaceSlice[0].(map[string]interface{})
	actualCountryIso := parsedStructuredAddressMap["country_iso"]

	if countryIso != actualCountryIso {
		t.Errorf("expected country_iso: %q, actual value was: %q", countryIso, actualCountryIso)
	}
}

func TestYotiClient_MissingPostalAddress_UsesFormattedAddress(t *testing.T) {
	var formattedAddressText = `House No.86-A\nRajgura Nagar\nLudhina\nPunjab\n141012\nIndia`

	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A",
		"formatted_address": "` + formattedAddressText + `"
	}
	`)

	structuredAddress, err := attribute.UnmarshallJSON(structuredAddressBytes)
	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var userProfile = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 ""}

	var jsonAttribute = &yotiprotoattr.Attribute{
		Name:        attrConstStructuredPostalAddress,
		Value:       structuredAddressBytes,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     []*yotiprotoattr.Anchor{},
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
		t.Errorf(
			"Address does not equal the expected formatted address. address: %q, formatted address: %q",
			profileAddress,
			formattedAddressText)
	}

	if userProfileAddress != escapedFormattedAddressText {
		t.Errorf(
			"Address does not equal the expected formatted address. address: %q, formatted address: %q",
			userProfileAddress,
			formattedAddressText)
	}

	var structuredPostalAddress *attribute.JSONAttribute

	structuredPostalAddress, err = profile.StructuredPostalAddress()
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(structuredPostalAddress.ContentType, yotiprotoattr.ContentType_JSON) {
		t.Errorf(
			"Retrieved attribute does not have the correct type. Expected %q, actual: %q",
			yotiprotoattr.ContentType_JSON,
			structuredPostalAddress.ContentType)
	}
}

func TestYotiClient_PresentPostalAddress_DoesntUseFormattedAddress(t *testing.T) {
	var addressText = `PostalAddress`

	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A",
		"formatted_address": "FormattedAddress"
	}`)
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
	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A"
	}`)

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

func TestProfile_GetAttribute_String(t *testing.T) {
	attributeName := "test_attribute_name"
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

	if att.Name != attributeName {
		t.Errorf(
			"Retrieved attribute does not have the correct name. Expected %q, actual: %q",
			attributeName,
			att.Name)
	}

	if !cmp.Equal(att.Value().(string), attributeValueString) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			attributeValue,
			att.Value())
	}
}

func TestProfile_GetAttribute_Time(t *testing.T) {
	attributeName := "test_attribute_name"

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

	if !cmp.Equal(expectedDate, att.Value().(*time.Time).UTC()) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			expectedDate,
			att.Value().(*time.Time))
	}
}

func TestProfile_GetAttribute_Jpeg(t *testing.T) {
	attributeName := "test_attribute_name"
	attributeValue := []byte("value")

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	if !cmp.Equal(att.Value().([]byte), attributeValue) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			attributeValue,
			att.Value())
	}
}

func TestProfile_GetAttribute_Png(t *testing.T) {
	attributeName := "test_attribute_name"
	attributeValue := []byte("value")

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attr)
	att := result.GetAttribute(attributeName)

	if !cmp.Equal(att.Value().([]byte), attributeValue) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			attributeValue,
			att.Value())
	}
}

func TestProfile_GetAttribute_Bool(t *testing.T) {
	attributeName := "test_attribute_name"
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
	if err != nil {
		t.Errorf("Unable to parse string to bool. Error: %s", err)
	}

	if !cmp.Equal(initialBoolValue, boolValue) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %v, actual: %v",
			initialBoolValue,
			boolValue)
	}
}

func TestProfile_GetAttribute_JSON(t *testing.T) {
	attributeName := "test_attribute_name"
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

	if !cmp.Equal(actualAddressFormat, addressFormat) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			addressFormat,
			actualAddressFormat)
	}
}

func TestProfile_GetAttribute_Undefined(t *testing.T) {
	attributeName := "test_attribute_name"
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

	if att.Name != attributeName {
		t.Errorf(
			"Retrieved attribute does not have the correct name. Expected %q, actual: %q",
			attributeName,
			att.Name)
	}

	if !cmp.Equal(att.Value().(string), attributeValueString) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			attributeValue,
			att.Value())
	}
}
func TestProfile_GetAttribute_ReturnsNil(t *testing.T) {
	result := Profile{
		attributeSlice: []*yotiprotoattr.Attribute{},
	}

	attribute := result.GetAttribute("attributeName")

	if attribute != nil {
		t.Error("Attribute should not be retrieved if it is not present")
	}
}

func TestProfile_StringAttribute(t *testing.T) {
	attributeName := attrConstNationality
	attributeValueString := "value"
	attributeValueBytes := []byte(attributeValueString)

	var as = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValueBytes,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(as)

	if result.Nationality().Value() != attributeValueString {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			attributeValueString,
			result.Nationality().Value())
	}

	if !cmp.Equal(result.Nationality().ContentType, yotiprotoattr.ContentType_STRING) {
		t.Errorf(
			"Retrieved attribute does not have the correct type. Expected %q, actual: %q",
			yotiprotoattr.ContentType_STRING,
			result.Nationality().ContentType)
	}
}

func TestProfile_AttributeProperty_RetrievesAttribute(t *testing.T) {
	attributeName := attrConstSelfie
	attributeValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if selfie.Name != attributeName {
		t.Errorf(
			"Retrieved attribute does not have the correct name. Expected %q, actual: %q",
			attributeName,
			selfie.Name)
	}

	if !reflect.DeepEqual(attributeValue, selfie.Value().Data) {
		t.Errorf(
			"Retrieved attribute does not have the correct value. Expected %q, actual: %q",
			attributeValue,
			selfie.Value().Data)
	}

	if !cmp.Equal(selfie.ContentType, yotiprotoattr.ContentType_PNG) {
		t.Errorf(
			"Retrieved attribute does not have the correct type. Expected %q, actual: %q",
			yotiprotoattr.ContentType_PNG,
			selfie.ContentType)
	}
}

func TestAttributeImage_Image_Png(t *testing.T) {
	attributeName := attrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if !cmp.Equal(selfie.Value().Data, byteValue) {
		t.Errorf(
			"Retrieved attribute does not have the correct Image. Expected %v, actual: %v",
			byteValue,
			selfie.Value().Data)
	}
}

func TestAttributeImage_Image_Jpeg(t *testing.T) {
	attributeName := attrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if !cmp.Equal(selfie.Value().Data, byteValue) {
		t.Errorf(
			"Retrieved attribute does not have the correct byte value. Expected %v, actual: %v",
			byteValue,
			selfie.Value().Data)
	}
}

func TestAttributeImage_Image_Default(t *testing.T) {
	attributeName := attrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	if !cmp.Equal(selfie.Value().Data, byteValue) {
		t.Errorf(
			"Retrieved attribute does not have the correct byte value. Expected %v, actual: %v",
			byteValue,
			selfie.Value().Data)
	}
}
func TestAttributeImage_Base64Selfie_Png(t *testing.T) {
	attributeName := attrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/png;base64;," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	if base64Selfie != expectedBase64Selfie {
		t.Errorf(
			"Base64Selfie does not have the correct value. Expected %q, actual: %q",
			expectedBase64Selfie,
			base64Selfie)
	}
}

func TestAttributeImage_Base64URL_Jpeg(t *testing.T) {
	attributeName := attrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/jpeg;base64;," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	if base64Selfie != expectedBase64Selfie {
		t.Errorf(
			"Base64Selfie does not have the correct value. Expected %q, actual: %q",
			expectedBase64Selfie,
			base64Selfie)
	}
}

func TestAnchorParser_Passport(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	anchorSlice := createAnchorSliceFromTestFile(t, "testanchorpassport.txt")

	var structuredAddressBytes = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A"
	}`)

	a := &yotiprotoattr.Attribute{
		Name:        attrConstStructuredPostalAddress,
		Value:       structuredAddressBytes,
		ContentType: yotiprotoattr.ContentType_JSON,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(a)

	var actualStructuredPostalAddress *attribute.JSONAttribute

	actualStructuredPostalAddress, err := result.StructuredPostalAddress()

	if err != nil {
		t.Error(err)
	}

	actualAnchor := actualStructuredPostalAddress.Anchors()[0]

	if actualAnchor != actualStructuredPostalAddress.Sources()[0] {
		t.Error("Anchors and Sources should be the same when there is only one Source")
	}

	if actualAnchor.Type() != anchor.AnchorTypeSource {
		t.Errorf(
			"Parsed anchor type is incorrect. Expected: %q, actual: %q",
			anchor.AnchorTypeSource,
			actualAnchor.Type())
	}

	expectedDate := time.Date(2018, time.April, 12, 13, 14, 32, 0, time.UTC)
	actualDate := actualAnchor.SignedTimestamp().Timestamp().UTC()
	if actualDate != expectedDate {
		t.Errorf(
			"Parsed anchor SignedTimestamp is incorrect. Expected: %q, actual: %q",
			expectedDate,
			actualDate)
	}

	expectedSubType := "OCR"
	if actualAnchor.SubType() != expectedSubType {
		t.Errorf(
			"Parsed anchor SubType is incorrect. Expected: %q, actual: %q",
			expectedSubType,
			actualAnchor.SubType())
	}

	expectedValue := "PASSPORT"
	if actualAnchor.Value()[0] != expectedValue {
		t.Errorf(
			"Parsed anchor Value is incorrect. Expected: %q, actual: %q",
			expectedValue,
			actualAnchor.Value()[0])
	}

	actualSerialNo := actualAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "277870515583559162487099305254898397834", actualSerialNo)
}

func TestAnchorParser_DrivingLicense(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "testanchordrivinglicense.txt")

	attribute := &yotiprotoattr.Attribute{
		Name:        attrConstGender,
		Value:       []byte("value"),
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attribute)

	genderAttribute := result.Gender()
	resultAnchor := genderAttribute.Anchors()[0]

	if resultAnchor != genderAttribute.Sources()[0] {
		t.Error("Anchors and Sources should be the same when there is only one Source")
	}

	if resultAnchor.Type() != anchor.AnchorTypeSource {
		t.Errorf(
			"Parsed anchor type is incorrect. Expected: %q, actual: %q",
			anchor.AnchorTypeSource,
			resultAnchor.Type())
	}

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 3, 0, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	if actualDate != expectedDate {
		t.Errorf(
			"Parsed anchor SignedTimestamp is incorrect. Expected: %q, actual: %q",
			expectedDate,
			actualDate)
	}

	expectedSubType := ""
	if resultAnchor.SubType() != expectedSubType {
		t.Errorf(
			"Parsed anchor SubType is incorrect. Expected: %q, actual: %q",
			expectedSubType,
			resultAnchor.SubType())
	}

	expectedValue := "DRIVING_LICENCE"
	if resultAnchor.Value()[0] != expectedValue {
		t.Errorf(
			"Parsed anchor Value is incorrect. Expected: %q, actual: %q",
			expectedValue,
			resultAnchor.Value()[0])
	}

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "46131813624213904216516051554755262812", actualSerialNo)
}

func TestAnchorParser_YotiAdmin(t *testing.T) {
	anchorSlice := createAnchorSliceFromTestFile(t, "testanchoryotiadmin.txt")

	attr := &yotiprotoattr.Attribute{
		Name:        attrConstDateOfBirth,
		Value:       []byte("1999-01-01"),
		ContentType: yotiprotoattr.ContentType_DATE,
		Anchors:     anchorSlice,
	}

	result := createProfileWithSingleAttribute(attr)

	DoB, err := result.DateOfBirth()

	if err != nil {
		t.Error(err)
	}

	resultAnchor := DoB.Anchors()[0]

	if resultAnchor != DoB.Verifiers()[0] {
		t.Error("Anchors and Verifiers should be the same when there is only one Verifier")
	}

	if resultAnchor.Type() != anchor.AnchorTypeVerifier {
		t.Errorf(
			"Parsed anchor type is incorrect. Expected: %q, actual: %q",
			anchor.AnchorTypeVerifier,
			resultAnchor.Type())
	}

	expectedDate := time.Date(2018, time.April, 11, 12, 13, 4, 0, time.UTC)
	actualDate := resultAnchor.SignedTimestamp().Timestamp().UTC()
	if actualDate != expectedDate {
		t.Errorf(
			"Parsed anchor SignedTimestamp is incorrect. Expected: %q, actual: %q",
			expectedDate,
			actualDate)
	}

	expectedSubType := ""
	if resultAnchor.SubType() != expectedSubType {
		t.Errorf(
			"Parsed anchor SubType is incorrect. Expected: %q, actual: %q",
			expectedSubType,
			resultAnchor.SubType())
	}

	expectedValue := "YOTI_ADMIN"
	if resultAnchor.Value()[0] != expectedValue {
		t.Errorf(
			"Parsed anchor Value is incorrect. Expected: %q, actual: %q",
			expectedValue,
			resultAnchor.Value()[0])
	}

	actualSerialNo := resultAnchor.OriginServerCerts()[0].SerialNumber
	assertServerCertSerialNo(t, "256616937783084706710155170893983549581", actualSerialNo)
}

func TestAnchors_None(t *testing.T) {
	anchorSlice := []*anchor.Anchor{}

	sources := anchor.GetSources(anchorSlice)
	if len(sources) > 0 {
		t.Error("GetSources should not return anything with empty anchors")
	}

	verifiers := anchor.GetVerifiers(anchorSlice)
	if len(verifiers) > 0 {
		t.Error("GetVerifiers should not return anything with empty anchors")
	}
}

func TestDateOfBirthAttribute(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "testattributedateofbirth.txt")

	dateOfBirthAttribute, err := attribute.NewTime(protoAttribute)

	if err != nil {
		t.Errorf("error creating time attribute: %q", err)
	}

	expectedDateOfBirth := time.Date(1970, time.December, 01, 0, 0, 0, 0, time.UTC)
	actualDateOfBirth := dateOfBirthAttribute.Value()

	if !actualDateOfBirth.Equal(expectedDateOfBirth) {
		t.Errorf(
			"Parsed attribute Date of Birth is incorrect. Expected: %q, actual: %q",
			expectedDateOfBirth,
			actualDateOfBirth)
	}
}

func TestNewImageSlice(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "testattributemultivalue.txt")

	documentImagesAttribute, err := attribute.NewImageSlice(protoAttribute)

	if err != nil {
		t.Errorf("error creating image slice attribute: %q", err)
	}

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttribute.Value(), documentImagesAttribute.Anchors()[0])
}

func TestNewMultiValue(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "testattributemultivalue.txt")

	multiValueAttribute, err := attribute.NewMultiValue(protoAttribute)

	if err != nil {
		t.Errorf("error creating multi value attribute: %q", err)
	}

	var documentImagesAttributeItems []*attribute.Image = attribute.CreateImageSlice(multiValueAttribute.Value())

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttributeItems, multiValueAttribute.Anchors()[0])
}

func TestNestedMultiValue(t *testing.T) {
	var innerMultiValueProtoValue []byte = createAttributeFromTestFile(t, "testattributemultivalue.txt").Value

	var intMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_STRING,
		Data:        []byte("string"),
	}

	var multiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Data:        innerMultiValueProtoValue,
	}

	var multiValueItemSlice = []*yotiprotoattr.MultiValue_Value{intMultiValueItem, multiValueItem}
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

	multiValueAttribute, err := attribute.NewMultiValue(protoAttribute)

	if err != nil {
		t.Errorf("error creating multi value attribute: %q", err)
	}

	for key, value := range multiValueAttribute.Value() {
		switch key {
		case 0:
			value0 := value.GetValue()
			if value0.(string) != "string" {
				t.Errorf("Unexpected Value: %q", value0)
			}
		case 1:
			value1 := value.GetValue()

			innerItems, ok := value1.([]*attribute.Item)
			if !ok {
				t.Errorf("Unexpected Value: %q", value1)
			}

			for innerKey, item := range innerItems {
				switch innerKey {
				case 0:
					assertIsExpectedImage(t, parseImage(t, item.GetValue()), "jpeg", "vWgD//2Q==")

				case 1:
					assertIsExpectedImage(t, parseImage(t, item.GetValue()), "jpeg", "38TVEH/9k=")
				}
			}
		}
	}
}

func TestMultiValueGenericGetter(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "testattributemultivalue.txt")
	profile := createProfileWithSingleAttribute(protoAttribute)

	multiValueAttribute := profile.GetAttribute(attrConstDocumentImages)

	// We need to cast, since GetAttribute always returns generic attributes
	multiValueAttributeValue := multiValueAttribute.Value().([]*attribute.Item)
	imageSlice := attribute.CreateImageSlice(multiValueAttributeValue)

	assertIsExpectedDocumentImagesAttribute(t, imageSlice, multiValueAttribute.Anchors()[0])
}

func parseImage(t *testing.T, innerImageInterface interface{}) *attribute.Image {
	innerImageBytes, ok := innerImageInterface.([]byte)
	if !ok {
		t.Errorf("Unexpected Value: %q", innerImageInterface)
	}

	innerImage, err := attribute.ParseImageValue(yotiprotoattr.ContentType_JPEG, innerImageBytes)
	if err != nil {
		t.Errorf("error parsing image: %q", err)
	}

	return innerImage
}

func assertIsExpectedDocumentImagesAttribute(t *testing.T, actualDocumentImages []*attribute.Image, anchor *anchor.Anchor) {
	if len(actualDocumentImages) != 2 {
		t.Error("This Document Images attribute should have two images")
	}

	assertIsExpectedImage(t, actualDocumentImages[0], "jpeg", "vWgD//2Q==")
	assertIsExpectedImage(t, actualDocumentImages[1], "jpeg", "38TVEH/9k=")

	expectedValue := "NATIONAL_ID"
	if anchor.Value()[0] != expectedValue {
		t.Errorf(
			"Parsed anchor Value is incorrect. Expected: %q, actual: %q",
			expectedValue,
			anchor.Value()[0])
	}

	expectedSubType := "STATE_ID"
	if anchor.SubType() != expectedSubType {
		t.Errorf(
			"Parsed anchor SubType is incorrect. Expected: %q, actual: %q",
			expectedSubType,
			anchor.SubType())
	}
}

func assertIsExpectedImage(t *testing.T, image *attribute.Image, imageType string, expectedBase64URLLast10 string) {
	if image.Type != imageType {
		t.Errorf(
			"Incorrect image type. Expected: %q, actual: %q",
			imageType,
			image.Type)
	}

	actualBase64URL := image.Base64URL()

	ActualBase64URLLast10Chars := actualBase64URL[len(actualBase64URL)-10:]

	if ActualBase64URLLast10Chars != expectedBase64URLLast10 {
		t.Errorf("Base64URL does not match. Expected: %q, actual: %q", expectedBase64URLLast10, ActualBase64URLLast10Chars)
	}
}

func marshallMultiValue(t *testing.T, multiValue *yotiprotoattr.MultiValue) []byte {
	marshalled, err := proto.Marshal(multiValue)

	if err != nil {
		t.Errorf("Unable to marshall MULTI_VALUE value. Error: %q", err)
	}

	return marshalled
}

func assertServerCertSerialNo(t *testing.T, expectedSerialNo string, actualSerialNo *big.Int) {
	expectedSerialNoBigInt := new(big.Int)
	expectedSerialNoBigInt, ok := expectedSerialNoBigInt.SetString(expectedSerialNo, 10)
	if !ok {
		t.Error("Unexpected error when setting string as big int")
	}

	if expectedSerialNoBigInt.Cmp(actualSerialNo) != 0 { //0 == equivalent
		t.Errorf(
			"Parsed anchor OriginServerCerts is incorrect. Expected: %q, actual: %q",
			expectedSerialNo,
			actualSerialNo)
	}
}

func createAttributeFromTestFile(t *testing.T, filename string) *yotiprotoattr.Attribute {
	attributeBytes, err := decodeTestFile(t, filename)

	if err != nil {
		t.Errorf("error decoding test file: %q", err)
	}

	attributeStruct := &yotiprotoattr.Attribute{}

	if err := proto.Unmarshal(attributeBytes, attributeStruct); err != nil {
		t.Errorf("Unable to parse MULTI_VALUE value: %q. Error: %q", string(attributeBytes), err)
	}

	return attributeStruct
}

func createAnchorSliceFromTestFile(t *testing.T, filename string) []*yotiprotoattr.Anchor {
	anchorBytes, err := decodeTestFile(t, filename)

	if err != nil {
		t.Errorf("error decoding test file: %q", err)
	}

	protoAnchor := &yotiprotoattr.Anchor{}
	if err := proto.Unmarshal(anchorBytes, protoAnchor); err != nil {
		t.Errorf("Error converting test anchor bytes into a Protobuf anchor. Error: %q", err)
	}

	protoAnchors := append([]*yotiprotoattr.Anchor{}, protoAnchor)

	return protoAnchors
}

func decodeTestFile(t *testing.T, filename string) (result []byte, err error) {
	base64Bytes := readTestFile(t, filename)
	base64String := string(base64Bytes)
	filebytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return filebytes, nil
}

func createProfileWithSingleAttribute(attr *yotiprotoattr.Attribute) Profile {
	var attributeSlice []*yotiprotoattr.Attribute
	attributeSlice = append(attributeSlice, attr)

	return Profile{
		attributeSlice: attributeSlice,
	}
}

func readTestFile(t *testing.T, filename string) (result []byte) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}

	return b
}

func createTestHeaders() (result map[string]string) {
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
