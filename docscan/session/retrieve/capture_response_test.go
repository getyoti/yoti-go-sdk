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
			{"type": "LIVENESS", "id": "id3", "state": "pending"},
			{"type": "FACE_CAPTURE", "id": "id4", "state": "pending"},
			{"type": "UNKNOWN_TYPE", "id": "id5", "state": "pending"}
		]
	}`)

	var c CaptureResponse
	err := json.Unmarshal(jsonData, &c)
	assert.NilError(t, err)

	assert.Equal(t, "given", c.BiometricConsent)
	assert.Equal(t, 5, len(c.RequiredResources))

	_, ok := c.RequiredResources[0].(*RequiredIdDocumentResourceResponse)
	assert.Assert(t, ok)
	assert.Equal(t, "ID_DOCUMENT", c.RequiredResources[0].GetType())
	assert.Equal(t, "id1", c.RequiredResources[0].(*RequiredIdDocumentResourceResponse).ID)

	_, ok = c.RequiredResources[1].(*RequiredSupplementaryDocumentResourceResponse)
	assert.Assert(t, ok)
	assert.Equal(t, "SUPPLEMENTARY_DOCUMENT", c.RequiredResources[1].GetType())

	_, ok = c.RequiredResources[2].(*RequiredZoomLivenessResourceResponse)
	assert.Assert(t, ok)
	assert.Equal(t, "LIVENESS", c.RequiredResources[2].GetType())

	_, ok = c.RequiredResources[3].(*RequiredFaceCaptureResourceResponse)
	assert.Assert(t, ok)
	assert.Equal(t, "FACE_CAPTURE", c.RequiredResources[3].GetType())

	unknownRes, ok := c.RequiredResources[4].(*UnknownRequiredResourceResponse)
	assert.Assert(t, ok)
	assert.Equal(t, "UNKNOWN_TYPE", unknownRes.GetType())
	assert.Assert(t, unknownRes.String() != "")

}

func TestCaptureResponse_Getters(t *testing.T) {
	c := CaptureResponse{
		RequiredResources: []RequiredResourceResponse{
			&RequiredIdDocumentResourceResponse{BaseRequiredResource{Type: "ID_DOCUMENT", ID: "id1"}},
			&RequiredSupplementaryDocumentResourceResponse{BaseRequiredResource{Type: "SUPPLEMENTARY_DOCUMENT", ID: "id2"}},
			&RequiredZoomLivenessResourceResponse{BaseRequiredResource{Type: "LIVENESS", ID: "id3"}},
			&RequiredFaceCaptureResourceResponse{BaseRequiredResource{Type: "FACE_CAPTURE", ID: "id4"}},
			&RequiredLivenessResourceResponse{BaseRequiredResource{Type: "LIVENESS", ID: "id5"}},
		},
	}

	docResources := c.GetDocumentResourceRequirements()
	assert.Equal(t, 2, len(docResources))
	types := []string{docResources[0].GetType(), docResources[1].GetType()}
	assert.Assert(t, (types[0] == "ID_DOCUMENT" && types[1] == "SUPPLEMENTARY_DOCUMENT") || (types[1] == "ID_DOCUMENT" && types[0] == "SUPPLEMENTARY_DOCUMENT"))

	idDocs := c.GetIdDocumentResourceRequirements()
	assert.Equal(t, 1, len(idDocs))
	assert.Equal(t, "id1", idDocs[0].ID)

	suppDocs := c.GetSupplementaryResourceRequirements()
	assert.Equal(t, 1, len(suppDocs))
	assert.Equal(t, "id2", suppDocs[0].ID)

	zoomLiveness := c.GetZoomLivenessResourceRequirements()
	assert.Equal(t, 1, len(zoomLiveness))
	assert.Equal(t, "id3", zoomLiveness[0].ID)

	faceCapture := c.GetFaceCaptureResourceRequirements()
	assert.Equal(t, 1, len(faceCapture))
	assert.Equal(t, "id4", faceCapture[0].ID)
}

func TestCaptureResponse_EmptyResources(t *testing.T) {
	jsonData := []byte(`{"biometric_consent": "none", "required_resources": []}`)

	var c CaptureResponse
	err := json.Unmarshal(jsonData, &c)
	assert.NilError(t, err)
	assert.Equal(t, "none", c.BiometricConsent)
	assert.Equal(t, 0, len(c.RequiredResources))
}
