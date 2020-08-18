package retrieve

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
	"gotest.tools/v3/assert"
)

func TestResourceContainer_UnmarshalJSON(t *testing.T) {
	zoomLivenessResourceResponse := &LivenessResourceResponse{
		ResourceResponse: ResourceResponse{ID: "id1234"},
		LivenessType:     constants.Zoom,
	}

	var livenessResources []*LivenessResourceResponse
	livenessResources = append(livenessResources, zoomLivenessResourceResponse)
	livenessResources = append(livenessResources, &LivenessResourceResponse{
		ResourceResponse: ResourceResponse{ID: "other_id1234"},
		LivenessType:     "OTHER_TYPE",
	})

	resourceContainer := ResourceContainer{
		LivenessCapture: livenessResources,
	}
	marshalled, err := json.Marshal(&resourceContainer)
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(marshalled, &result)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(result.ZoomLivenessResources()))
	assert.Equal(t, "id1234", result.ZoomLivenessResources()[0].ID)
}
