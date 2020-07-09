package yoti

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/consts"
	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotocom"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
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

func createExtraDataContent(t *testing.T, pemBytes []byte, protoExtraData *yotiprotoshare.ExtraData, wrappedReceiptKey string) string {
	outBytes, err := proto.Marshal(protoExtraData)
	assert.NilError(t, err)

	keyBytes, _ := pem.Decode(pemBytes)
	key, err := x509.ParsePKCS1PrivateKey(keyBytes.Bytes)
	assert.NilError(t, err)
	cipherBytes, err := base64.StdEncoding.DecodeString(wrappedReceiptKey)
	assert.NilError(t, err)
	unwrappedKey, err := rsa.DecryptPKCS1v15(rand.Reader, key, cipherBytes)
	assert.NilError(t, err)
	cipherBlock, err := aes.NewCipher(unwrappedKey)
	assert.NilError(t, err)

	padLength := cipherBlock.BlockSize() - len(outBytes)%cipherBlock.BlockSize()
	outBytes = append(outBytes, bytes.Repeat([]byte{byte(padLength)}, padLength)...)

	iv := make([]byte, cipherBlock.BlockSize())
	encrypter := cipher.NewCBCEncrypter(cipherBlock, iv)
	encrypter.CryptBlocks(outBytes, outBytes)

	outProto := &yotiprotocom.EncryptedData{
		CipherText: outBytes,
		Iv:         iv,
	}
	outBytes, err = proto.Marshal(outProto)
	assert.NilError(t, err)

	return base64.StdEncoding.EncodeToString(outBytes)
}

func TestYotiClient_DefaultHTTPClientShouldTimeout(t *testing.T) {
	client := Client{}
	defer func() {
		_ = recover()
		assert.Assert(t, client.HTTPClient.(*http.Client).Timeout == 10*time.Second)
	}()
	_, _ = client.doRequest(nil)
}

func TestYotiClient_SetHTTPClientTimeout(t *testing.T) {
	client := Client{}
	client.HTTPClient = &http.Client{Timeout: 12 * time.Minute}
	defer func() {
		_ = recover()
		assert.Assert(t, client.HTTPClient.(*http.Client).Timeout == 12*time.Minute)
	}()
	_, _ = client.doRequest(nil)
}

func TestYotiClient_KeyLoad_Failure(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key-invalid-format.pem")
	_, err := NewClient("", key)
	assert.Check(t, err != nil)
	assert.Check(t, strings.HasPrefix(err.Error(), "Invalid Key"))
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestNewYotiClient_InvalidToken(t *testing.T) {
	var err error
	key, _ := ioutil.ReadFile("test-key.pem")

	client, err := NewClient("sdkId", key)
	assert.NilError(t, err)

	_, err = client.getActivityDetails("")

	assert.Check(t, err != nil)
	assert.Check(t, strings.HasPrefix(err.Error(), "Invalid Token"))
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_HttpFailure_ReturnsFailure(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	_, err = client.getActivityDetails(encryptedToken)

	assert.Check(t, err != nil)
	assert.ErrorContains(t, err, "Unknown HTTP Error")
	tempError, temporary := err.(interface {
		Temporary() bool
		Unwrap() error
	})
	assert.Check(t, temporary)
	assert.Check(t, tempError.Temporary())
	assert.ErrorContains(t, tempError.Unwrap(), "Unknown HTTP Error")
}

func TestYotiClient_HttpFailure_ReturnsProfileNotFound(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 404,
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	_, err = client.getActivityDetails(encryptedToken)

	assert.Check(t, err != nil)
	assert.Check(t, strings.HasPrefix(err.Error(), "Profile Not Found"))
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_SharingFailure_ReturnsFailure(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`{"session_data":"session_data","receipt":{"receipt_id": null,"other_party_profile_content": null,"policy_uri":null,"personal_key":null,"remember_me_id":null, "sharing_outcome":"FAILURE","timestamp":"2016-09-23T13:04:11Z"}}`)),
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	_, err = client.getActivityDetails(encryptedToken)

	assert.Check(t, err != nil)
	assert.Check(t, strings.HasPrefix(err.Error(), ErrSharingFailure.Error()))
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_TokenDecodedSuccessfully(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	expectedAbsoluteURL := "/api/v1/profile/" + encryptedToken

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(request *http.Request) (*http.Response, error) {
				parsed, err := url.Parse(request.URL.String())
				assert.Assert(t, is.Nil(err), "Yoti API did not generate a valid URI.")
				assert.Equal(t, parsed.Path, expectedAbsoluteURL, "Yoti API did not generate a valid URL path.")

				return &http.Response{
					StatusCode: 500,
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	_, err = client.getActivityDetails(encryptedToken)

	assert.Check(t, err != nil)
	assert.Check(t, strings.HasPrefix(err.Error(), "Unknown HTTP Error"))
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func TestYotiClient_ParseProfile_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	rememberMeID := "remember_me_id0123456789"

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`)),
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	activityDetails, errorStrings := client.GetActivityDetails(encryptedToken)

	assert.Assert(t, is.Nil(errorStrings))

	profile := activityDetails.UserProfile

	assert.Equal(t, activityDetails.RememberMeID(), rememberMeID)

	assert.Assert(t, is.Nil(activityDetails.ExtraData().AttributeIssuanceDetails()))

	expectedSelfieValue := "selfie0123456789"

	assert.DeepEqual(t, profile.Selfie().Value().Data, []byte(expectedSelfieValue))
	assert.Equal(t, profile.MobileNumber().Value(), "phone_number0123456789")

	assert.Equal(
		t,
		profile.GetAttribute("phone_number").Value(),
		"phone_number0123456789",
	)

	assert.Check(t,
		profile.GetImageAttribute("doesnt_exist") == nil,
	)

	assert.Check(t, profile.GivenNames() == nil)
	assert.Check(t, profile.FamilyName() == nil)
	assert.Check(t, profile.FullName() == nil)
	assert.Check(t, profile.EmailAddress() == nil)
	images, _ := profile.DocumentImages()
	assert.Check(t, images == nil)
	documentDetails, _ := profile.DocumentDetails()
	assert.Check(t, documentDetails == nil)

	expectedDoB := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)

	actualDoB, err := profile.DateOfBirth()
	assert.Assert(t, is.Nil(err))

	assert.Assert(t, actualDoB != nil)
	assert.DeepEqual(t, actualDoB.Value(), &expectedDoB)
}

func TestYotiClient_ParentRememberMeID(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")
	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	parentRememberMeID := "parent_remember_me_id0123456789"

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey +
						`","other_party_profile_content": "` + otherPartyProfileContent +
						`","parent_remember_me_id":"` + parentRememberMeID + `", "sharing_outcome":"SUCCESS"}}`)),
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	activityDetails, errorStrings := client.getActivityDetails(encryptedToken)

	assert.Assert(t, is.Nil(errorStrings))
	assert.Equal(t, activityDetails.ParentRememberMeID(), parentRememberMeID)
}
func TestYotiClient_ParseWithoutProfile_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")
	rememberMeID := "remember_me_id0123456789"
	timestamp := time.Date(1973, 11, 29, 9, 33, 9, 0, time.UTC)
	timestampString := func(a []byte, _ error) string {
		return string(a)
	}(timestamp.MarshalText())
	receiptID := "receipt_id123"

	var otherPartyProfileContents = []string{
		`"other_party_profile_content": null,`,
		`"other_party_profile_content": "",`,
		``}

	for _, otherPartyProfileContent := range otherPartyProfileContents {

		client := Client{
			HTTPClient: &mockHTTPClient{
				do: func(*http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: 200,
						Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `",` +
							otherPartyProfileContent + `"remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS", "timestamp":"` + timestampString + `", "receipt_id":"` + receiptID + `"}}`)),
					}, nil
				},
			},
		}
		var err error
		client.Key, err = loadRsaKey(key)
		assert.NilError(t, err)

		activityDetails, errStrings := client.getActivityDetails(encryptedToken)

		assert.Assert(t, is.Nil(errStrings))
		assert.Equal(t, activityDetails.RememberMeID(), rememberMeID)
		assert.Equal(t, activityDetails.Timestamp(), timestamp)
		assert.Equal(t, activityDetails.ReceiptID(), receiptID)
	}
}
func TestYotiClient_ShouldParseAndDecryptExtraDataContent(t *testing.T) {
	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	rememberMeID := "remember_me_id0123456789"

	pemBytes, err := ioutil.ReadFile("test-key.pem")
	assert.NilError(t, err)

	attributeName := "attributeName"
	dataEntries := make([]*yotiprotoshare.DataEntry, 0)
	expiryDate := time.Now().UTC().AddDate(0, 0, 1)
	thirdPartyAttributeDataEntry := test.CreateThirdPartyAttributeDataEntry(t, &expiryDate, []string{attributeName}, "tokenValue")

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	extraDataContent := createExtraDataContent(t, pemBytes, protoExtraData, wrappedReceiptKey)

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` +
						wrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","extra_data_content": "` +
						extraDataContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`)),
				}, nil
			},
		},
	}
	client.Key, err = loadRsaKey(pemBytes)
	assert.NilError(t, err)

	activityDetails, err := client.getActivityDetails(encryptedToken)
	assert.NilError(t, err)

	assert.Equal(t, rememberMeID, activityDetails.RememberMeID())
	assert.Assert(t, activityDetails.ExtraData().AttributeIssuanceDetails() != nil)
	assert.Equal(t, activityDetails.UserProfile.MobileNumber().Value(), "phone_number0123456789")
}

func TestYotiClient_ShouldCarryOnProcessingIfIssuanceTokenIsNotPresent(t *testing.T) {
	var attributeName = "attributeName"
	dataEntries := make([]*yotiprotoshare.DataEntry, 0)
	expiryDate := time.Now().UTC().AddDate(0, 0, 1)
	thirdPartyAttributeDataEntry := test.CreateThirdPartyAttributeDataEntry(t, &expiryDate, []string{attributeName}, "")

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	pemBytes, err := ioutil.ReadFile("test-key.pem")
	assert.NilError(t, err)

	extraDataContent := createExtraDataContent(t, pemBytes, protoExtraData, wrappedReceiptKey)

	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"

	rememberMeID := "remember_me_id0123456789"

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` +
						wrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","extra_data_content": "` +
						extraDataContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`)),
				}, nil
			},
		},
	}
	client.Key, err = loadRsaKey(pemBytes)
	assert.NilError(t, err)

	activityDetails, err := client.getActivityDetails(encryptedToken)

	assert.Check(t, err != nil)
	assert.Check(t, strings.Contains(err.Error(), "Issuance Token is invalid"))

	assert.Equal(t, rememberMeID, activityDetails.RememberMeID())
	assert.Assert(t, is.Nil(activityDetails.ExtraData().AttributeIssuanceDetails()))
	assert.Equal(t, activityDetails.UserProfile.MobileNumber().Value(), "phone_number0123456789")
}
func TestYotiClient_ParseWithoutRememberMeID_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	var otherPartyProfileContents = []string{
		`"other_party_profile_content": null,`,
		`"other_party_profile_content": "",`}

	for _, otherPartyProfileContent := range otherPartyProfileContents {

		client := Client{
			HTTPClient: &mockHTTPClient{
				do: func(*http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: 200,
						Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + wrappedReceiptKey + `",` +
							otherPartyProfileContent + `"sharing_outcome":"SUCCESS"}}`)),
					}, nil
				},
			},
		}
		var err error
		client.Key, err = loadRsaKey(key)
		assert.NilError(t, err)

		_, errStrings := client.getActivityDetails(encryptedToken)

		assert.Assert(t, is.Nil(errStrings))
	}
}

func TestYotiClient_ParseIsAgeVerifiedValue_True(t *testing.T) {
	trueValue := []byte("true")

	isAgeVerified, err := parseIsAgeVerifiedValue(trueValue)

	assert.Assert(t, is.Nil(err), "Failed to parse IsAgeVerified value")
	assert.Check(t, *isAgeVerified)
}

func TestYotiClient_ParseIsAgeVerifiedValue_False(t *testing.T) {
	falseValue := []byte("false")

	isAgeVerified, err := parseIsAgeVerifiedValue(falseValue)

	assert.Assert(t, is.Nil(err), "Failed to parse IsAgeVerified value")
	assert.Check(t, !*isAgeVerified)

}
func TestYotiClient_ParseIsAgeVerifiedValue_InvalidValueThrowsError(t *testing.T) {
	invalidValue := []byte("invalidBool")

	_, err := parseIsAgeVerifiedValue(invalidValue)

	assert.Assert(t, err != nil)
}
func TestYotiClient_UnmarshallJSONValue_InvalidValueThrowsError(t *testing.T) {
	invalidStructuredAddress := []byte("invalidBool")

	_, err := attribute.UnmarshallJSON(invalidStructuredAddress)

	assert.Assert(t, err != nil)
}

func TestYotiClient_UnmarshallJSONValue_ValidValue(t *testing.T) {
	const (
		countryIso  = "IND"
		nestedValue = "NestedValue"
	)

	var structuredAddress = []byte(`
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
	`)

	parsedStructuredAddress, err := attribute.UnmarshallJSON(structuredAddress)

	assert.Assert(t, is.Nil(err), "Failed to parse structured address")

	actualCountryIso := parsedStructuredAddress["country_iso"]

	assert.Equal(t, countryIso, actualCountryIso)
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

func TestClient_OverrideAPIURL_ShouldSetAPIURL(t *testing.T) {
	client := &Client{}
	expectedURL := "expectedurl.com"
	client.OverrideAPIURL(expectedURL)
	assert.Equal(t, client.getAPIURL(), expectedURL)
}
