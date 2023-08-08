package yoti

import (
	"fmt"
	"os"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
	"gotest.tools/v3/assert"
)

func TestDigitalIDClient(t *testing.T) {
	key, err := os.ReadFile("./test/test-key.pem")
	assert.NilError(t, err)

	_, err = NewDigitalIdentityClient("some-sdk-id", key)
	assert.NilError(t, err)
}

func TestDigitalIDClient_KeyLoad_Failure(t *testing.T) {
	key, err := os.ReadFile("test/test-key-invalid-format.pem")
	assert.NilError(t, err)

	_, err = NewDigitalIdentityClient("", key)

	assert.ErrorContains(t, err, "invalid key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestDigitalIDClient_CreateShareURL(t *testing.T) {

	policy, err := (&digitalidentity.PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	assert.NilError(t, err)

	session, err := (&digitalidentity.ShareSessionBuilder{}).WithPolicy(policy).Build()
	assert.NilError(t, err)
	fmt.Println(session)

}
