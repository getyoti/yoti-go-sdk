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
