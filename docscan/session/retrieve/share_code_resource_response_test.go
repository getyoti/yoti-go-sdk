package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestShareCodeResourceResponse_ShouldUnmarshalCorrectly(t *testing.T) {
	jsonData := `{
		"id": "share-code-123",
		"source": "test-source",
		"created_at": "2026-01-14T10:00:00Z",
		"last_updated": "2026-01-14T11:00:00Z",
		"tasks": [
			{
				"type": "VERIFY_SHARE_CODE_TASK",
				"id": "task-123",
				"state": "DONE"
			}
		]
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, "share-code-123", response.ID)
	assert.Equal(t, "test-source", response.Source)
	assert.Equal(t, "2026-01-14T10:00:00Z", response.CreatedAt)
	assert.Equal(t, "2026-01-14T11:00:00Z", response.LastUpdated)
	assert.Equal(t, 1, len(response.VerifyShareCodeTasks()))
	assert.Equal(t, "VERIFY_SHARE_CODE_TASK", response.VerifyShareCodeTasks()[0].Type)
	assert.Equal(t, "task-123", response.VerifyShareCodeTasks()[0].ID)
	assert.Equal(t, "DONE", response.VerifyShareCodeTasks()[0].State)
}

func TestShareCodeResourceResponse_WithMultipleTasks(t *testing.T) {
	jsonData := `{
		"id": "share-code-456",
		"source": "test-source-2",
		"created_at": "2026-01-14T10:00:00Z",
		"last_updated": "2026-01-14T12:00:00Z",
		"tasks": [
			{
				"type": "VERIFY_SHARE_CODE_TASK",
				"id": "task-1",
				"state": "PENDING"
			},
			{
				"type": "VERIFY_SHARE_CODE_TASK",
				"id": "task-2",
				"state": "DONE"
			}
		]
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, 2, len(response.VerifyShareCodeTasks()))
	assert.Equal(t, "task-1", response.VerifyShareCodeTasks()[0].ID)
	assert.Equal(t, "task-2", response.VerifyShareCodeTasks()[1].ID)
}

func TestShareCodeResourceResponse_WithMediaFields(t *testing.T) {
	jsonData := `{
		"id": "share-code-789",
		"source": "test-source-3",
		"created_at": "2026-01-14T10:00:00Z",
		"last_updated": "2026-01-14T13:00:00Z",
		"lookup_profile": {
			"id": "media-1",
			"type": "JSON"
		},
		"returned_profile": {
			"id": "media-2",
			"type": "JSON"
		},
		"id_photo": {
			"id": "media-3",
			"type": "IMAGE"
		},
		"file": {
			"id": "media-4",
			"type": "PDF"
		},
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, "share-code-789", response.ID)
	assert.Assert(t, response.LookupProfile != nil)
	assert.Equal(t, "media-1", response.LookupProfile.ID)
	assert.Equal(t, "JSON", response.LookupProfile.Type)
	assert.Assert(t, response.ReturnedProfile != nil)
	assert.Equal(t, "media-2", response.ReturnedProfile.ID)
	assert.Assert(t, response.IDPhoto != nil)
	assert.Equal(t, "media-3", response.IDPhoto.ID)
	assert.Equal(t, "IMAGE", response.IDPhoto.Type)
	assert.Assert(t, response.File != nil)
	assert.Equal(t, "media-4", response.File.ID)
	assert.Equal(t, "PDF", response.File.Type)
}

func TestShareCodeResourceResponse_WithNoTasks(t *testing.T) {
	jsonData := `{
		"id": "share-code-000",
		"source": "test-source-0",
		"created_at": "2026-01-14T10:00:00Z",
		"last_updated": "2026-01-14T10:30:00Z",
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, 0, len(response.VerifyShareCodeTasks()))
}

func TestShareCodeResourceResponse_WithMixedTaskTypes(t *testing.T) {
	jsonData := `{
		"id": "share-code-mixed",
		"source": "test-source-mixed",
		"created_at": "2026-01-14T10:00:00Z",
		"last_updated": "2026-01-14T14:00:00Z",
		"tasks": [
			{
				"type": "VERIFY_SHARE_CODE_TASK",
				"id": "task-verify",
				"state": "DONE"
			},
			{
				"type": "OTHER_TASK_TYPE",
				"id": "task-other",
				"state": "PENDING"
			}
		]
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, 2, len(response.Tasks))
	assert.Equal(t, 1, len(response.VerifyShareCodeTasks()))
	assert.Equal(t, "task-verify", response.VerifyShareCodeTasks()[0].ID)
}
