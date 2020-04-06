package profilesandbox

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2"
	"github.com/getyoti/yoti-go-sdk/v2/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v2/file"
	"github.com/getyoti/yoti-go-sdk/v2/profile/sandbox"

	_ "github.com/joho/godotenv/autoload"
	"gotest.tools/v3/assert"
)

func TestAnchorExample(t *testing.T) {
	sandboxClientSdkId := os.Getenv("YOTI_SANDBOX_CLIENT_SDK_ID")
	pemFileBytes, err := file.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))
	assert.NilError(t, err)

	privateKey, err := cryptoutil.ParseRSAKey(pemFileBytes)
	assert.NilError(t, err)

	sandboxClient := sandbox.Client{
		ClientSdkID: sandboxClientSdkId,
		Key:         privateKey,
	}

	sourceAnchor := sandbox.SourceAnchor("NFC", time.Now().UTC(), "PASSPORT")
	verifierAnchor := sandbox.VerifierAnchor("", time.Now().UTC(), "YOTI_ADMIN")

	tokenRequest := sandbox.TokenRequest{}.
		WithFamilyName("Smith", []sandbox.Anchor{sourceAnchor, verifierAnchor})

	sandboxToken, err := sandboxClient.SetupSharingProfile(tokenRequest)
	assert.NilError(t, err)

	yotiClient := yoti.Client{
		Key:   pemFileBytes,
		SdkID: sandboxClientSdkId,
	}
	yotiClient.OverrideAPIURL("https://api.yoti.com/sandbox/v1")

	activityDetails, errStrings := yotiClient.GetActivityDetails(sandboxToken)
	if len(errStrings) > 0 {
		log.Fatalf("%v", errStrings)
	}

	// Test your application's logic here
	assert.Equal(t, "NFC", activityDetails.UserProfile.FamilyName().Sources()[0].SubType())
}
