package yoti

import (
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
	"time"
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

	_, err := getActivityDetails(requester, encryptedToken, sdkID, key)

	if err == nil {
		t.Error("Expected failure")
	} else if !strings.HasPrefix(err.Error(), "Invalid Key") {
		t.Errorf("expected outcome type starting with %q instead received %q", "Invalid Key", err.Error())
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

	_, err := getActivityDetails(requester, encryptedToken, sdkID, key)
	if err == nil {
		t.Error("Expected failure")
	} else if err != ErrFailure {
		t.Errorf("expected outcome type %q instead received %q", ErrFailure.Error(), err.Error())
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

	_, err := getActivityDetails(requester, encryptedToken, sdkID, key)
	if err == nil {
		t.Error("Expected failure")
	} else if err != ErrProfileNotFound {
		t.Errorf("expected outcome type %q instead received %q", ErrProfileNotFound.Error(), err.Error())
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

	_, err := getActivityDetails(requester, encryptedToken, sdkID, key)
	if err == nil {
		t.Error("Expected failure")
	} else if err != ErrSharingFailure {
		t.Errorf("expected outcome type %q instead received %q", ErrSharingFailure.Error(), err.Error())
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

	_, err := getActivityDetails(requester, encryptedToken, sdkID, key)
	if err == nil {
		t.Error("Expected failure")
	} else if err != ErrFailure {
		t.Errorf("expected outcome type %q instead received %q", ErrFailure.Error(), err.Error())
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

	profile, err := getActivityDetails(requester, encryptedToken, sdkID, key)

	if err != nil {
		t.Error(err)
	}

	if profile.ID != rememberMeID {
		t.Errorf("expected id %q instead received %q", rememberMeID, profile.ID)
	}

	if profile.Selfie == nil {
		t.Error(`expected user selfie but it was not present in the returned profile`)
	} else if string(profile.Selfie.Data) != "selfie0123456789" {
		t.Errorf("expected user selfie %q instead received %q", "selfie0123456789", string(profile.Selfie.Data))
	}

	if profile.MobileNumber != "phone_number0123456789" {
		t.Errorf("expected user mobile %q instead received %q", "phone_number0123456789", profile.MobileNumber)
	}

	dob := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)
	if profile.DateOfBirth == nil {
		t.Error(`expected date of birth but it was not present in the returned profile`)
	} else if !profile.DateOfBirth.Equal(dob) {
		t.Errorf("expected date of birth %q instead received %q", profile.DateOfBirth.Format(time.UnixDate), dob.Format(time.UnixDate))
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

		profile, err := getActivityDetails(requester, encryptedToken, sdkID, key)

		if err != nil {
			t.Error(err)
		}

		if profile.ID != rememberMeID {
			t.Errorf("expected id %q instead received %q", rememberMeID, profile.ID)
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
		CreateStandardAmlProfile(),
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
		CreateStandardAmlProfile(),
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

	_, err := unmarshallJSON(invalidStructuredAddress)

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

	parsedStructuredAddress, err := unmarshallJSON(structuredAddress)

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

	structuredAddress, err := unmarshallJSON(structuredAddressBytes)
	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var result = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 ""}

	address, err := getFormattedAddressIfAddressIsMissing(result)
	if err != nil {
		t.Errorf("Failed to add formatted address to address, error was %q", err.Error())
	}

	escapedFormattedAddressText := strings.Replace(formattedAddressText, `\n`, "\n", -1)

	if address != escapedFormattedAddressText {
		t.Errorf("Address does not equal the expected formatted address. address: %q, formatted address: %q", address, formattedAddressText)
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
	structuredAddress, err := unmarshallJSON(structuredAddressBytes)

	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var result = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 addressText}

	newFormattedAddress, err := getFormattedAddressIfAddressIsMissing(result)

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

	structuredAddress, err := unmarshallJSON(structuredAddressBytes)

	if err != nil {
		t.Errorf("Failed to parse structured address, error was %q", err.Error())
	}

	var result = UserProfile{
		ID:                      "remember_me_id0123456789",
		OtherAttributes:         make(map[string]AttributeValue),
		StructuredPostalAddress: structuredAddress,
		Address:                 ""}

	address, err := getFormattedAddressIfAddressIsMissing(result)

	if err != nil {
		t.Errorf("Failed to add formatted address to address, error was %q", err.Error())
	}

	if address != "" {
		t.Errorf("Formatted address missing, but address was still changed to: %q", address)
	}
}

func CreateHeaders() (result map[string]string) {

	headers := make(map[string]string)

	headers["Header1"] = "test"

	return headers
}

func CreateStandardAmlProfile() (result AmlProfile) {
	var amlAddress = AmlAddress{
		Country: "GBR"}

	var amlProfile = AmlProfile{
		GivenNames: "Edward Richard George",
		FamilyName: "Heath",
		Address:    amlAddress}

	return amlProfile
}
