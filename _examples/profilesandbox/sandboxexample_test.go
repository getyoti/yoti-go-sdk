package profilesandbox

import (
	"encoding/base64"
	"os"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3"
	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/file"
	"github.com/getyoti/yoti-go-sdk/v3/profile/sandbox"

	_ "github.com/joho/godotenv/autoload"
	"gotest.tools/v3/assert"
)

func TestExample(t *testing.T) {
	sandboxClientSdkId := os.Getenv("YOTI_SANDBOX_CLIENT_SDK_ID")
	pemFileBytes, err := file.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	assert.NilError(t, err)

	privateKey, err := cryptoutil.ParseRSAKey(pemFileBytes)
	assert.NilError(t, err)

	sandboxClient := sandbox.Client{
		ClientSdkID: sandboxClientSdkId,
		Key:         privateKey,
	}

	var dateOfBirthUnder18 = time.Now().AddDate(-10, 0, 0)
	var documentImages = sandbox.DocumentImages{}.WithJpegImage([]byte("some_image_value"))

	tokenRequest := (&sandbox.TokenRequest{}).
		WithRememberMeID("remember_me_id_12345").
		WithAgeVerification(dateOfBirthUnder18, sandbox.Derivation{}.AgeUnder(18), nil).
		WithGivenNames("some given names", nil).
		WithFamilyName("some family name", nil).
		WithFullName("some full name", nil).
		WithDateOfBirth(dateOfBirthUnder18, nil).
		WithGender("some gender", nil).
		WithPhoneNumber("some phone number", nil).
		WithNationality("some nationality", nil).
		WithStructuredPostalAddress(
			map[string]interface{}{
				"building_number": "1",
				"address_line1":   "some street name",
			}, nil).
		WithBase64Selfie(base64.StdEncoding.EncodeToString([]byte("some_image_value")), nil).
		WithEmailAddress("some@email", nil).
		WithDocumentDetails("PASSPORT USA 1234abc", nil).
		WithDocumentImages(documentImages, nil)

	sandboxToken, err := sandboxClient.SetupSharingProfile(tokenRequest)
	assert.NilError(t, err)

	yotiClient, err := yoti.NewClient(
		sandboxClientSdkId,
		pemFileBytes,
	)
	assert.NilError(t, err)
	yotiClient.OverrideAPIURL("https://api.yoti.com/sandbox/v1")

	activityDetails, err := yotiClient.GetActivityDetails(sandboxToken)
	assert.NilError(t, err)

	// Test your application's logic here
	assert.Equal(t, "some@email", activityDetails.UserProfile.EmailAddress().Value())
}
