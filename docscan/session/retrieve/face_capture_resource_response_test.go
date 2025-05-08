package retrieve

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/file"
	"gotest.tools/v3/assert"
)

func TestResourceContainer_UnmarshalJSON_FaceCapture(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container-face.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(result.FaceCapture))
	assert.Equal(t, "string", result.FaceCapture[0].ID)
	assert.Equal(t, "string", result.FaceCapture[0].Source.Type)
	assert.Equal(t, "string", result.FaceCapture[0].Image.Media.ID)
	assert.Equal(t, "string", result.FaceCapture[0].Image.Media.Type)

	expectedCreated, err := time.Parse(time.RFC3339, "2025-05-08T12:00:00Z")
	assert.NilError(t, err)
	expectedLastUpdated, err := time.Parse(time.RFC3339, "2025-05-08T12:05:00Z")
	assert.NilError(t, err)

	assert.Assert(t, result.FaceCapture[0].Image.Media.Created.Equal(expectedCreated), "Created timestamps are not equal")
	assert.Assert(t, result.FaceCapture[0].Image.Media.LastUpdated.Equal(expectedLastUpdated), "LastUpdated timestamps are not equal")
	assert.DeepEqual(t, []json.RawMessage{}, result.FaceCapture[0].Tasks)
}
func TestResourceContainer_UnmarshalJSON_FaceCapture_Multiple(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container-face-multiple.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	assert.Equal(t, 2, len(result.FaceCapture))

	// First Face Capture
	assert.Equal(t, "face_id_1", result.FaceCapture[0].ID)
	assert.Equal(t, "upload", result.FaceCapture[0].Source.Type)
	assert.Equal(t, "media_id_1", result.FaceCapture[0].Image.Media.ID)
	assert.Equal(t, "image/jpeg", result.FaceCapture[0].Image.Media.Type)

	expectedCreated1, err := time.Parse(time.RFC3339, "2025-05-06T12:00:00Z")
	assert.NilError(t, err)
	expectedLastUpdated1, err := time.Parse(time.RFC3339, "2025-05-06T12:05:00Z")
	assert.NilError(t, err)

	assert.Assert(t, result.FaceCapture[0].Image.Media.Created.Equal(expectedCreated1), "First Face Capture Created timestamps are not equal")
	assert.Assert(t, result.FaceCapture[0].Image.Media.LastUpdated.Equal(expectedLastUpdated1), "First Face Capture LastUpdated timestamps are not equal")
	assert.DeepEqual(t, []json.RawMessage{json.RawMessage(`"verification"`), json.RawMessage(`"quality_check"`)}, result.FaceCapture[0].Tasks)

	// Second Face Capture
	assert.Equal(t, "face_id_2", result.FaceCapture[1].ID)
	assert.Equal(t, "camera", result.FaceCapture[1].Source.Type)
	assert.Equal(t, "media_id_2", result.FaceCapture[1].Image.Media.ID)
	assert.Equal(t, "image/png", result.FaceCapture[1].Image.Media.Type)

	expectedCreated2, err := time.Parse(time.RFC3339, "2025-05-05T12:10:00Z")
	assert.NilError(t, err)
	expectedLastUpdated2, err := time.Parse(time.RFC3339, "2025-05-05T12:15:00Z")
	assert.NilError(t, err)

	assert.Assert(t, result.FaceCapture[1].Image.Media.Created.Equal(expectedCreated2), "Second Face Capture Created timestamps are not equal")
	assert.Assert(t, result.FaceCapture[1].Image.Media.LastUpdated.Equal(expectedLastUpdated2), "Second Face Capture LastUpdated timestamps are not equal")
	assert.DeepEqual(t, []json.RawMessage{}, result.FaceCapture[1].Tasks)
}

func TestResourceContainer_UnmarshalJSON_FaceCapture_EmptyArray(t *testing.T) {
	bytes, err := file.ReadFile("../../../test/fixtures/resource-container-face-empty.json")
	assert.NilError(t, err)

	var result ResourceContainer
	err = json.Unmarshal(bytes, &result)
	assert.NilError(t, err)

	assert.Equal(t, 0, len(result.FaceCapture))
}
