package retrieve

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/file"
	"gotest.tools/v3/assert"
)

func TestLivenessResourceResponse_UnmarshalJSON(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	assert.Equal(t, 2, len(result.LivenessCapture))
	assert.Equal(t, "ZOOM", result.LivenessCapture[0].LivenessType)
	assert.Equal(t, "OTHER_LIVENESS_TYPE", result.LivenessCapture[1].LivenessType)

	assert.Equal(t, "IMAGE", result.ZoomLivenessResources()[0].Frames[0].Media.Type)
	assert.Equal(t, "BINARY", result.ZoomLivenessResources()[0].FaceMap.Media.Type)
}

func TestStaticLivenessResourceResponse_UnmarshalJSON(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container-static.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	assert.Equal(t, 3, len(result.LivenessCapture))
	assert.Equal(t, "STATIC", result.LivenessCapture[0].LivenessType)
}

func TestLivenessResourceResponse_UnmarshalJSON_Invalid(t *testing.T) {
	var result ResourceContainer
	err := result.UnmarshalJSON([]byte("some-invalid-json"))
	assert.ErrorContains(t, err, "invalid character")
}

func TestResourceContainer_filterForCheck_Zoom(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	check := &CheckResponse{ResourcesUsed: []string{"a831bc40-e3c2-11ea-87d0-0242ac130003"}}
	filtered := result.filterForCheck(check)

	assert.Equal(t, 1, len(filtered.LivenessCapture))
	assert.Equal(t, 1, len(filtered.RawLivenessCapture))
	assert.Equal(t, 1, len(filtered.ZoomLivenessResources()))
	assert.Equal(t, 0, len(filtered.StaticLivenessResources()))

	// The original container must be untouched.
	assert.Equal(t, 2, len(result.LivenessCapture))
}

func TestResourceContainer_filterForCheck_Static(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container-static.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	check := &CheckResponse{ResourcesUsed: []string{"1bf432f3-63fd-4fa7-b243-84adc305dd5f"}}
	filtered := result.filterForCheck(check)

	assert.Equal(t, 1, len(filtered.LivenessCapture))
	assert.Equal(t, 1, len(filtered.StaticLivenessResources()))
	assert.Equal(t, 0, len(filtered.ZoomLivenessResources()))

	// The original container must be untouched.
	assert.Equal(t, 3, len(result.LivenessCapture))
}

func TestResourceContainer_filterForCheck_NoMatches(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	check := &CheckResponse{ResourcesUsed: []string{"unknown-id"}}
	filtered := result.filterForCheck(check)

	assert.Equal(t, 0, len(filtered.LivenessCapture))
	assert.Equal(t, 0, len(filtered.RawLivenessCapture))
	assert.Equal(t, 0, len(filtered.ZoomLivenessResources()))
	assert.Equal(t, 0, len(filtered.StaticLivenessResources()))
}

func TestResourceContainer_filterForCheck_NilReceiverAndNilCheck(t *testing.T) {
	var result *ResourceContainer
	filtered := result.filterForCheck(nil)
	assert.Assert(t, filtered != nil)
	assert.Equal(t, 0, len(filtered.IDDocuments))

	result = &ResourceContainer{IDDocuments: []*IDDocumentResourceResponse{{ResourceResponse: &ResourceResponse{ID: "some-id"}}}}
	filtered = result.filterForCheck(nil)
	assert.Assert(t, filtered != nil)
	assert.Equal(t, 0, len(filtered.IDDocuments))
}

func TestResourceContainer_filterForCheck_DocumentsAndShareCodes(t *testing.T) {
	result := &ResourceContainer{
		IDDocuments: []*IDDocumentResourceResponse{
			{ResourceResponse: &ResourceResponse{ID: "id-doc-1"}},
			{ResourceResponse: &ResourceResponse{ID: "id-doc-2"}},
			{ResourceResponse: nil},
		},
		SupplementaryDocuments: []*SupplementaryDocumentResourceResponse{
			{ResourceResponse: &ResourceResponse{ID: "supp-doc-1"}},
		},
		ShareCodes: []*ShareCodeResourceResponse{
			{ResourceResponse: &ResourceResponse{ID: "share-code-1"}},
		},
	}

	check := &CheckResponse{ResourcesUsed: []string{"id-doc-1", "supp-doc-1", "share-code-1"}}
	filtered := result.filterForCheck(check)

	assert.Equal(t, 1, len(filtered.IDDocuments))
	assert.Equal(t, "id-doc-1", filtered.IDDocuments[0].GetID())
	assert.Equal(t, 1, len(filtered.SupplementaryDocuments))
	assert.Equal(t, 1, len(filtered.ShareCodes))
}
