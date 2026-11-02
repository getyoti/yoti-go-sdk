package retrieve_test

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/retrieve"
	"gotest.tools/v3/assert"
)

func TestGetSessionResult_UnmarshalJSON_WithCompanyProfileAndOrganisationName(t *testing.T) {
	raw := []byte(`{
		"session_id": "test-session-id",
		"organisation_name": "Acme Organisation",
		"company_profile": {
			"company_name": "Acme Corp"
		}
	}`)

	var result retrieve.GetSessionResult
	err := result.UnmarshalJSON(raw)
	assert.NilError(t, err)

	assert.Equal(t, "Acme Organisation", result.GetOrganisationName())

	cp := result.GetCompanyProfile()
	assert.Assert(t, cp != nil)
	assert.Equal(t, "Acme Corp", cp.GetCompanyName())
}

func TestGetSessionResult_UnmarshalJSON_WithNullCompanyProfile(t *testing.T) {
	raw := []byte(`{
		"session_id": "test-session-id",
		"company_profile": null
	}`)

	var result retrieve.GetSessionResult
	err := result.UnmarshalJSON(raw)
	assert.NilError(t, err)

	assert.Assert(t, result.GetCompanyProfile() == nil)
}

func TestGetSessionResult_UnmarshalJSON_WithAbsentCompanyProfileAndOrganisationName(t *testing.T) {
	raw := []byte(`{
		"session_id": "test-session-id"
	}`)

	var result retrieve.GetSessionResult
	err := result.UnmarshalJSON(raw)
	assert.NilError(t, err)

	assert.Assert(t, result.GetCompanyProfile() == nil)
	assert.Equal(t, "", result.GetOrganisationName())
}

func TestCompanyProfileResponse_GetCompanyName(t *testing.T) {
	data := []byte(`{"company_name":"Test Company Ltd"}`)

	var cp retrieve.CompanyProfileResponse
	err := json.Unmarshal(data, &cp)
	assert.NilError(t, err)

	assert.Equal(t, "Test Company Ltd", cp.GetCompanyName())
}
