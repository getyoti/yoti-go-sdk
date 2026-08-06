package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCaptureResponse_UnmarshalJSON(t *testing.T) {
	jsonData := []byte(`{
		"biometric_consent": "given",
		"required_resources": [
			{"type": "ID_DOCUMENT", "id": "id1", "state": "pending"},
			{"type": "SUPPLEMENTARY_DOCUMENT", "id": "id2", "state": "pending"},
			{"type": "LIVENESS", "id": "id3", "state": "pending", "liveness_type": "ZOOM"},
			{"type": "LIVENESS", "id": "id4", "state": "pending", "liveness_type": "STATIC"},
			{"type": "FACE_CAPTURE", "id": "id5", "state": "pending"},
			{"type": "UNKNOWN_TYPE", "id": "id6", "state": "pending"}
		]
	}`)

	var c CaptureResponse
	err := json.Unmarshal(jsonData, &c)
	assert.NilError(t, err)
	assert.Equal(t, "given", c.BiometricConsent)
	assert.Equal(t, 6, len(c.RequiredResources))

	_, ok := c.RequiredResources[0].(*RequiredIdDocumentResourceResponse)
	assert.Assert(t, ok)

	_, ok = c.RequiredResources[1].(*RequiredSupplementaryDocumentResourceResponse)
	assert.Assert(t, ok)

	_, ok = c.RequiredResources[2].(*RequiredZoomLivenessResourceResponse)
	assert.Assert(t, ok)

	_, ok = c.RequiredResources[3].(*RequiredStaticLivenessResourceResponse)
	assert.Assert(t, ok)

	_, ok = c.RequiredResources[4].(*RequiredFaceCaptureResourceResponse)
	assert.Assert(t, ok)

	_, ok = c.RequiredResources[5].(*UnknownRequiredResourceResponse)
	assert.Assert(t, ok)
}

func TestCaptureResponse_UnmarshalJSON_SupportedCountries(t *testing.T) {
	jsonData := []byte(`{
		"biometric_consent": "given",
		"required_resources": [
			{
				"type": "ID_DOCUMENT",
				"id": "id1",
				"state": "pending",
				"supported_countries": [
					{
						"code": "IND",
						"supported_documents": [
							{"type": "AADHAAR", "providers": ["DIGILOCKER"]}
						]
					}
				]
			}
		]
	}`)

	var c CaptureResponse
	err := json.Unmarshal(jsonData, &c)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(c.RequiredResources))

	idDocument, ok := c.RequiredResources[0].(*RequiredIdDocumentResourceResponse)
	assert.Assert(t, ok)

	assert.Equal(t, 1, len(idDocument.SupportedCountries))
	assert.Equal(t, "IND", idDocument.SupportedCountries[0].Code)
	assert.Equal(t, 1, len(idDocument.SupportedCountries[0].SupportedDocuments))
	assert.Equal(t, "AADHAAR", idDocument.SupportedCountries[0].SupportedDocuments[0].Type)
	assert.Equal(t, 1, len(idDocument.SupportedCountries[0].SupportedDocuments[0].Providers))
	assert.Equal(t, "DIGILOCKER", idDocument.SupportedCountries[0].SupportedDocuments[0].Providers[0])
}

func TestCaptureResponse_UnmarshalJSON_SupportedCountriesOmitted(t *testing.T) {
	jsonData := []byte(`{
		"biometric_consent": "given",
		"required_resources": [
			{"type": "ID_DOCUMENT", "id": "id1", "state": "pending"}
		]
	}`)

	var c CaptureResponse
	err := json.Unmarshal(jsonData, &c)
	assert.NilError(t, err)

	idDocument, ok := c.RequiredResources[0].(*RequiredIdDocumentResourceResponse)
	assert.Assert(t, ok)
	assert.Equal(t, 0, len(idDocument.SupportedCountries))
}

func TestCaptureResponse_Getters(t *testing.T) {
	c := CaptureResponse{
		RequiredResources: []RequiredResourceResponse{
			&RequiredIdDocumentResourceResponse{BaseRequiredResource: BaseRequiredResource{Type: "ID_DOCUMENT", ID: "id1"}},
			&RequiredSupplementaryDocumentResourceResponse{BaseRequiredResource{Type: "SUPPLEMENTARY_DOCUMENT", ID: "id2"}},
			&RequiredZoomLivenessResourceResponse{BaseRequiredResource{Type: "LIVENESS", ID: "id3", LivenessType: "ZOOM"}},
			&RequiredStaticLivenessResourceResponse{BaseRequiredResource{Type: "LIVENESS", ID: "id4", LivenessType: "STATIC"}},
			&RequiredFaceCaptureResourceResponse{BaseRequiredResource{Type: "FACE_CAPTURE", ID: "id5"}},
		},
	}

	assert.Equal(t, 2, len(c.GetDocumentResourceRequirements()))
	assert.Equal(t, 1, len(c.GetIdDocumentResourceRequirements()))
	assert.Equal(t, 1, len(c.GetSupplementaryResourceRequirements()))
	assert.Equal(t, 1, len(c.GetZoomLivenessResourceRequirements()))
	assert.Equal(t, 1, len(c.GetStaticLivenessResourceRequirements()))
	assert.Equal(t, 1, len(c.GetFaceCaptureResourceRequirements()))
}

func TestCaptureResponse_EmptyResources(t *testing.T) {
	jsonData := []byte(`{"biometric_consent": "none", "required_resources": []}`)

	var c CaptureResponse
	err := json.Unmarshal(jsonData, &c)
	assert.NilError(t, err)
	assert.Equal(t, "none", c.BiometricConsent)
	assert.Equal(t, 0, len(c.RequiredResources))
}
