package profile

import (
	"net/http"
	"testing"

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

	assert.Assert(t, is.Nil(err), "Failed to parse IsAgeVerified value")
	assert.Check(t, *isAgeVerified)
}

func TestProfileService_ParseIsAgeVerifiedValue_False(t *testing.T) {
	falseValue := []byte("false")

	isAgeVerified, err := parseIsAgeVerifiedValue(falseValue)

	assert.Assert(t, is.Nil(err), "Failed to parse IsAgeVerified value")
	assert.Check(t, !*isAgeVerified)

}
func TestProfileService_ParseIsAgeVerifiedValue_InvalidValueThrowsError(t *testing.T) {
	invalidValue := []byte("invalidBool")

	_, err := parseIsAgeVerifiedValue(invalidValue)

	assert.Assert(t, err != nil)
}

func TestProfileService_ErrIsThrownForInvalidToken(t *testing.T) {
	_, err := GetActivityDetails(nil, "invalidToken", "clientSdkId", "apiUrl", test.GetValidKey("../test/test-key.pem"))

	assert.ErrorContains(t, err, "Unable to decrypt token")
}

func TestProfileService_RequestErrIsReturned(t *testing.T) {
	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
			}, nil
		},
	}
	_, err := GetActivityDetails(client, test.EncryptedToken, "clientSdkId", "https://apiUrl", test.GetValidKey("../test/test-key.pem"))

	assert.ErrorContains(t, err, "404: Profile Not Found")
}

func TestProfileService_ParseExtraData_ErrorDecrypting(t *testing.T) {
	receipt := &receiptDO{
		ExtraDataContent: "invalidExtraData",
	}
	_, err := parseExtraData(receipt, test.GetValidKey("../test/test-key.pem"), nil)

	assert.ErrorContains(t, err, "Unable to decrypt ExtraData from the receipt.")
}