package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestShareCodeResourceResponse(t *testing.T) {
	jsonData := `{
		"id": "share-code-123",
		"source": "test-source",
		"created_at": "2023-01-01T00:00:00Z",
		"last_updated": "2023-01-02T00:00:00Z",
		"lookup_profile": {
			"media": {
				"id": "media-1",
				"type": "JSON",
				"created": "2023-01-01T00:00:00Z",
				"last_updated": "2023-01-01T00:00:00Z"
			}
		},
		"returned_profile": {
			"media": {
				"id": "media-2",
				"type": "JSON",
				"created": "2023-01-01T00:00:00Z",
				"last_updated": "2023-01-01T00:00:00Z"
			}
		},
		"id_photo": {
			"media": {
				"id": "media-3",
				"type": "IMAGE",
				"created": "2023-01-01T00:00:00Z",
				"last_updated": "2023-01-01T00:00:00Z"
			}
		},
		"file": {
			"media": {
				"id": "media-4",
				"type": "PDF",
				"created": "2023-01-01T00:00:00Z",
				"last_updated": "2023-01-01T00:00:00Z"
			}
		},
		"tasks": [
			{
				"type": "VERIFY_SHARE_CODE_TASK",
				"id": "task-123",
				"state": "DONE",
				"created": "2023-01-01T00:00:00Z",
				"last_updated": "2023-01-02T00:00:00Z",
				"generated_media": [
					{
						"id": "gen-media-1",
						"type": "JSON"
					}
				]
			}
		]
	}`

	var shareCode ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &shareCode)
	assert.NilError(t, err)

	assert.Equal(t, shareCode.ID, "share-code-123")
	assert.Equal(t, shareCode.Source, "test-source")
	assert.Equal(t, shareCode.CreatedAt, "2023-01-01T00:00:00Z")
	assert.Equal(t, shareCode.LastUpdated, "2023-01-02T00:00:00Z")

	assert.Assert(t, shareCode.LookupProfile != nil)
	assert.Assert(t, shareCode.LookupProfile.Media != nil)
	assert.Equal(t, shareCode.LookupProfile.Media.ID, "media-1")

	assert.Assert(t, shareCode.ReturnedProfile != nil)
	assert.Assert(t, shareCode.ReturnedProfile.Media != nil)
	assert.Equal(t, shareCode.ReturnedProfile.Media.ID, "media-2")

	assert.Assert(t, shareCode.IDPhoto != nil)
	assert.Assert(t, shareCode.IDPhoto.Media != nil)
	assert.Equal(t, shareCode.IDPhoto.Media.ID, "media-3")

	assert.Assert(t, shareCode.File != nil)
	assert.Assert(t, shareCode.File.Media != nil)
	assert.Equal(t, shareCode.File.Media.ID, "media-4")

	assert.Equal(t, len(shareCode.Tasks), 1)
	assert.Equal(t, shareCode.Tasks[0].Type, "VERIFY_SHARE_CODE_TASK")
	assert.Equal(t, shareCode.Tasks[0].ID, "task-123")

	verifyTasks := shareCode.VerifyShareCodeTasks()
	assert.Equal(t, len(verifyTasks), 1)
	assert.Equal(t, verifyTasks[0].ID, "task-123")
}
