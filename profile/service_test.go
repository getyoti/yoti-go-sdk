package profile

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotocom"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoshare"
	"github.com/golang/protobuf/proto"

	"github.com/getyoti/yoti-go-sdk/v3/test"

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

func TestProfileService_ParseIsAgeVerifiedValue_True(t *testing.T) {
	trueValue := []byte("true")

	isAgeVerified, err := parseIsAgeVerifiedValue(trueValue)
	assert.NilError(t, err, "Failed to parse IsAgeVerified value")
	assert.Check(t, *isAgeVerified)
}

func TestProfileService_ParseIsAgeVerifiedValue_False(t *testing.T) {
	falseValue := []byte("false")

	isAgeVerified, err := parseIsAgeVerifiedValue(falseValue)
	assert.NilError(t, err, "Failed to parse IsAgeVerified value")
	assert.Check(t, !*isAgeVerified)

}
func TestProfileService_ParseIsAgeVerifiedValue_InvalidValueThrowsError(t *testing.T) {
	invalidValue := []byte("invalidBool")

	_, err := parseIsAgeVerifiedValue(invalidValue)

	assert.Assert(t, err != nil)
}

func TestProfileService_ErrIsThrownForInvalidToken(t *testing.T) {
	_, err := GetActivityDetails(nil, "invalidToken", "clientSdkId", "apiUrl", getValidKey())

	assert.ErrorContains(t, err, "unable to decrypt token")
}

func TestProfileService_RequestErrIsReturned(t *testing.T) {
	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
			}, nil
		},
	}
	_, err := GetActivityDetails(client, test.EncryptedToken, "clientSdkId", "https://apiUrl", getValidKey())

	assert.ErrorContains(t, err, "404: Profile not found")
}

func TestProfileService_InvalidToken(t *testing.T) {
	_, err := GetActivityDetails(nil, "", "sdkId", "https://apiurl", getValidKey())
	assert.ErrorContains(t, err, "invalid Token")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestProfileService_ParseExtraData_ErrorDecrypting(t *testing.T) {
	receipt := &receiptDO{
		ExtraDataContent: "invalidExtraData",
	}
	_, err := parseExtraData(receipt, getValidKey(), nil)

	assert.ErrorContains(t, err, "Unable to decrypt ExtraData from the receipt.")
}

func TestProfileService_GetActivityDetails(t *testing.T) {
	key := getValidKey()

	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	rememberMeID := "remember_me_id0123456789"

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + test.WrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS", "timestamp":"2006-01-02T15:04:05.999999Z"}}`)),
			}, nil
		},
	}

	activityDetails, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)
	assert.NilError(t, err)

	profile := activityDetails.UserProfile

	assert.Equal(t, activityDetails.RememberMeID(), rememberMeID)
	assert.Assert(t, is.Nil(activityDetails.ExtraData().AttributeIssuanceDetails()))

	expectedSelfieValue := "selfie0123456789"

	assert.DeepEqual(t, profile.Selfie().Value().Data(), []byte(expectedSelfieValue))
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
	assert.NilError(t, err)

	assert.Assert(t, actualDoB != nil)
	assert.DeepEqual(t, actualDoB.Value(), &expectedDoB)
}

func TestProfileService_SharingFailure_ReturnsFailure(t *testing.T) {
	key := getValidKey()

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"session_data":"session_data","receipt":{"receipt_id": null,"other_party_profile_content": null,"policy_uri":null,"personal_key":null,"remember_me_id":null, "sharing_outcome":"FAILURE","timestamp":"2016-09-23T13:04:11Z"}}`)),
			}, nil
		},
	}
	_, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)
	assert.ErrorContains(t, err, "sharing failure")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestProfileService_TokenDecodedSuccessfully(t *testing.T) {
	key := getValidKey()

	expectedPath := "/profile/" + test.Token

	client := &mockHTTPClient{
		do: func(request *http.Request) (*http.Response, error) {
			parsed, err := url.Parse(request.URL.String())
			assert.NilError(t, err, "Yoti API did not generate a valid URI.")
			assert.Equal(t, parsed.Path, expectedPath, "Yoti API did not generate a valid URL path.")

			return &http.Response{
				StatusCode: 500,
			}, nil
		},
	}

	_, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)
	assert.ErrorContains(t, err, "unknown HTTP error")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func TestProfileService_ParentRememberMeID(t *testing.T) {
	key := getValidKey()

	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	parentRememberMeID := "parent_remember_me_id0123456789"

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + test.WrappedReceiptKey +
					`","other_party_profile_content": "` + otherPartyProfileContent +
					`","parent_remember_me_id":"` + parentRememberMeID +
					`", "sharing_outcome":"SUCCESS", "timestamp":"2006-01-02T15:04:05.999999Z"}}`)),
			}, nil
		},
	}

	activityDetails, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)

	assert.NilError(t, err)
	assert.Equal(t, activityDetails.ParentRememberMeID(), parentRememberMeID)
}
func TestProfileService_ParseWithoutProfile_Success(t *testing.T) {
	key := getValidKey()

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

		client := &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + test.WrappedReceiptKey + `",` +
						otherPartyProfileContent + `"remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS", "timestamp":"` + timestampString + `", "receipt_id":"` + receiptID + `"}}`)),
				}, nil
			},
		}

		activityDetails, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)

		assert.NilError(t, err)
		assert.Equal(t, activityDetails.RememberMeID(), rememberMeID)
		assert.Equal(t, activityDetails.Timestamp(), timestamp)
		assert.Equal(t, activityDetails.ReceiptID(), receiptID)
	}
}

func TestProfileService_ShouldParseAndDecryptExtraDataContent(t *testing.T) {
	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	rememberMeID := "remember_me_id0123456789"

	pemBytes, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	dataEntries := make([]*yotiprotoshare.DataEntry, 0)
	expiryDate := time.Now().UTC().AddDate(0, 0, 1)
	thirdPartyAttributeDataEntry := test.CreateThirdPartyAttributeDataEntry(t, &expiryDate, []string{attributeName}, "tokenValue")

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	extraDataContent := createExtraDataContent(t, protoExtraData)

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` +
					test.WrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","extra_data_content": "` +
					extraDataContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS", "timestamp":"2006-01-02T15:04:05.999999Z"}}`)),
			}, nil
		},
	}
	key, err := cryptoutil.ParseRSAKey(pemBytes)
	assert.NilError(t, err)

	activityDetails, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)
	assert.NilError(t, err)

	assert.Equal(t, rememberMeID, activityDetails.RememberMeID())
	assert.Assert(t, activityDetails.ExtraData().AttributeIssuanceDetails() != nil)
	assert.Equal(t, activityDetails.UserProfile.MobileNumber().Value(), "phone_number0123456789")
}

func TestProfileService_ShouldCarryOnProcessingIfIssuanceTokenIsNotPresent(t *testing.T) {
	dataEntries := make([]*yotiprotoshare.DataEntry, 0)
	expiryDate := time.Now().UTC().AddDate(0, 0, 1)
	thirdPartyAttributeDataEntry := test.CreateThirdPartyAttributeDataEntry(t, &expiryDate, []string{attributeName}, "")

	dataEntries = append(dataEntries, &thirdPartyAttributeDataEntry)
	protoExtraData := &yotiprotoshare.ExtraData{
		List: dataEntries,
	}

	pemBytes, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	extraDataContent := createExtraDataContent(t, protoExtraData)

	otherPartyProfileContent := "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"

	rememberMeID := "remember_me_id0123456789"

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` +
					test.WrappedReceiptKey + `","other_party_profile_content": "` + otherPartyProfileContent + `","extra_data_content": "` +
					extraDataContent + `","remember_me_id":"` + rememberMeID + `", "sharing_outcome":"SUCCESS"}}`)),
			}, nil
		},
	}

	key, err := cryptoutil.ParseRSAKey(pemBytes)
	assert.NilError(t, err)

	activityDetails, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", key)

	assert.Check(t, err != nil)
	assert.Check(t, strings.Contains(err.Error(), "Issuance Token is invalid"))

	assert.Equal(t, rememberMeID, activityDetails.RememberMeID())
	assert.Assert(t, is.Nil(activityDetails.ExtraData().AttributeIssuanceDetails()))
	assert.Equal(t, activityDetails.UserProfile.MobileNumber().Value(), "phone_number0123456789")
}
func TestProfileService_ParseWithoutRememberMeID_Success(t *testing.T) {
	var otherPartyProfileContents = []string{
		`"other_party_profile_content": null,`,
		`"other_party_profile_content": "",`}

	for _, otherPartyProfileContent := range otherPartyProfileContents {

		client := &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{"receipt":{"wrapped_receipt_key": "` + test.WrappedReceiptKey + `",` +
						otherPartyProfileContent + `"sharing_outcome":"SUCCESS", "timestamp":"2006-01-02T15:04:05.999999Z"}}`)),
				}, nil
			},
		}
		_, err := GetActivityDetails(client, test.EncryptedToken, "sdkId", "https://apiurl", getValidKey())

		assert.NilError(t, err)
	}
}

func getValidKey() *rsa.PrivateKey {
	return test.GetValidKey("../test/test-key.pem")
}

func createExtraDataContent(t *testing.T, protoExtraData *yotiprotoshare.ExtraData) string {
	outBytes, err := proto.Marshal(protoExtraData)
	assert.NilError(t, err)

	key := getValidKey()
	assert.NilError(t, err)

	cipherBytes, err := base64.StdEncoding.DecodeString(test.WrappedReceiptKey)
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
