package yoti

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

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

func TestYotiClient_GetAPIURLUsesDefaultUrlAsFallback(t *testing.T) {
	client := Client{}

	os.Setenv("YOTI_API_URL", "")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/api/v1", result)
}
