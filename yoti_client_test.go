package yoti

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/getyoti/yoti-go-sdk/v3/aml"
	"github.com/getyoti/yoti-go-sdk/v3/attribute"
	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/profile"
	"github.com/getyoti/yoti-go-sdk/v3/test"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
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

func TestYotiClient_KeyLoad_Failure(t *testing.T) {
	key, _ := ioutil.ReadFile("test/test-key-invalid-format.pem")
	_, err := NewClient("", key)

	assert.ErrorContains(t, err, "Invalid Key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_HttpFailure_ReturnsFailure(t *testing.T) {
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.GetActivityDetails(test.EncryptedToken)

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
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 404,
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.GetActivityDetails(test.EncryptedToken)

	assert.ErrorContains(t, err, "Profile Not Found")
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
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

func TestClient_OverrideAPIURL_ShouldSetAPIURL(t *testing.T) {
	client := &Client{}
	expectedURL := "expectedurl.com"
	client.OverrideAPIURL(expectedURL)
	assert.Equal(t, client.getAPIURL(), expectedURL)
}

func TestYotiClient_GetAPIURLUsesOverriddenBaseUrlOverEnvVariable(t *testing.T) {
	client := Client{}
	client.OverrideAPIURL("overridenBaseUrl")

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	result := client.getAPIURL()

	assert.Equal(t, "overridenBaseUrl", result)
}

func TestYotiClient_GetAPIURLUsesEnvVariable(t *testing.T) {
	client := Client{}

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	result := client.getAPIURL()

	assert.Equal(t, "envBaseUrl", result)
}

func TestYotiClient_GetAPIURLUsesDefaultUrlAsFallbackWithEmptyEnvValue(t *testing.T) {
	client := Client{}

	os.Setenv("YOTI_API_URL", "")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/api/v1", result)
}

func TestYotiClient_GetAPIURLUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	client := Client{}

	os.Unsetenv("YOTI_API_URL")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/api/v1", result)
}

func createStandardAmlProfile() (result aml.AmlProfile) {
	var amlAddress = aml.AmlAddress{
		Country: "GBR"}

	var amlProfile = aml.AmlProfile{
		GivenNames: "Edward Richard George",
		FamilyName: "Heath",
		Address:    amlAddress}

	return amlProfile
}

func TestYotiClient_PerformAmlCheck_WithInvalidJSON(t *testing.T) {
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("Not a JSON document")),
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.PerformAmlCheck(createStandardAmlProfile())
	assert.Check(t, strings.Contains(err.Error(), "invalid character"))
}

func TestYotiClient_PerformAmlCheck_Success(t *testing.T) {
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`{"on_fraud_list":true,"on_pep_list":true,"on_watch_list":true}`)),
				}, nil
			},
		},
		Key: key,
	}

	result, err := client.PerformAmlCheck(createStandardAmlProfile())

	assert.Assert(t, is.Nil(err))

	assert.Check(t, result.OnFraudList)
	assert.Check(t, result.OnPEPList)
	assert.Check(t, result.OnWatchList)

}

func TestYotiClient_PerformAmlCheck_Unsuccessful(t *testing.T) {
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 503,
					Body:       ioutil.NopCloser(strings.NewReader(`SERVICE UNAVAILABLE - Unable to reach the Integrity Service`)),
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.PerformAmlCheck(createStandardAmlProfile())
	assert.ErrorContains(t, err, "AML Check was unsuccessful")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func TestAttributeImage_Image_Png(t *testing.T) {
	attributeName := consts.AttrSelfie
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
	attributeName := consts.AttrSelfie
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
	attributeName := consts.AttrSelfie
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
	attributeName := consts.AttrSelfie
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
	attributeName := consts.AttrSelfie
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

func createProfileWithSingleAttribute(attr *yotiprotoattr.Attribute) profile.UserProfile {
	var attributeSlice []*yotiprotoattr.Attribute
	attributeSlice = append(attributeSlice, attr)

	attributeList := &yotiprotoattr.AttributeList{
		Attributes: attributeSlice,
	}

	return profile.NewUserProfile(attributeList)
}

func getValidKey() *rsa.PrivateKey {
	return test.GetValidKey("test/test-key.pem")
}
