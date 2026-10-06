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

func TestShareCodeResourceResponse_ShouldUnmarshalWithObjectSource(t *testing.T) {
	jsonData := `{
		"id": "share-code-obj",
		"source": {"type": "END_USER"},
		"created_at": "2026-01-14T10:00:00Z",
		"last_updated": "2026-01-14T11:00:00Z",
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, "share-code-obj", response.ID)
	assert.Equal(t, "END_USER", response.Source)
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
			"media": { "id": "media-1", "type": "JSON" }
		},
		"returned_profile": {
			"media": { "id": "media-2", "type": "JSON" }
		},
		"id_photo": {
			"media": { "id": "media-3", "type": "IMAGE" }
		},
		"file": {
			"media": { "id": "media-4", "type": "PDF" }
		},
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, "share-code-789", response.ID)
	assert.Assert(t, response.LookupProfile != nil)
	assert.Assert(t, response.LookupProfile.Media != nil)
	assert.Equal(t, "media-1", response.LookupProfile.Media.ID)
	assert.Equal(t, "JSON", response.LookupProfile.Media.Type)
	assert.Assert(t, response.ReturnedProfile != nil)
	assert.Assert(t, response.ReturnedProfile.Media != nil)
	assert.Equal(t, "media-2", response.ReturnedProfile.Media.ID)
	assert.Assert(t, response.IDPhoto != nil)
	assert.Assert(t, response.IDPhoto.Media != nil)
	assert.Equal(t, "media-3", response.IDPhoto.Media.ID)
	assert.Equal(t, "IMAGE", response.IDPhoto.Media.Type)
	assert.Assert(t, response.File != nil)
	assert.Assert(t, response.File.Media != nil)
	assert.Equal(t, "media-4", response.File.Media.ID)
	assert.Equal(t, "PDF", response.File.Media.Type)
}

func TestShareCodeResourceResponse_WithNullMediaFields(t *testing.T) {
	jsonData := `{
		"id": "share-code-null-media",
		"source": {"type": "END_USER"},
		"created_at": "2026-02-05T11:33:46Z",
		"last_updated": "2026-02-05T11:33:50Z",
		"lookup_profile": null,
		"returned_profile": null,
		"id_photo": null,
		"file": null,
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Assert(t, response.LookupProfile == nil)
	assert.Assert(t, response.ReturnedProfile == nil)
	assert.Assert(t, response.IDPhoto == nil)
	assert.Assert(t, response.File == nil)
}

func TestShareCodeResourceResponse_WithMissingMediaFields(t *testing.T) {
	jsonData := `{
		"id": "share-code-no-media",
		"source": {"type": "END_USER"},
		"created_at": "2026-02-05T11:33:46Z",
		"last_updated": "2026-02-05T11:33:50Z",
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Assert(t, response.LookupProfile == nil)
	assert.Assert(t, response.ReturnedProfile == nil)
	assert.Assert(t, response.IDPhoto == nil)
	assert.Assert(t, response.File == nil)
}

func TestShareCodeResourceResponse_WithEmptyMediaObject(t *testing.T) {
	jsonData := `{
		"id": "share-code-empty-media",
		"source": {"type": "END_USER"},
		"created_at": "2026-02-05T11:33:46Z",
		"last_updated": "2026-02-05T11:33:50Z",
		"id_photo": {},
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Assert(t, response.IDPhoto != nil)
	assert.Assert(t, response.IDPhoto.Media == nil)
}

func TestShareCodeResourceResponse_WithPartialMediaFields(t *testing.T) {
	jsonData := `{
		"id": "share-code-partial",
		"source": {"type": "END_USER"},
		"created_at": "2026-02-05T11:33:46Z",
		"last_updated": "2026-02-05T11:33:50Z",
		"lookup_profile": {
			"media": { "id": "media-1", "type": "JSON" }
		},
		"tasks": []
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Assert(t, response.LookupProfile != nil)
	assert.Assert(t, response.LookupProfile.Media != nil)
	assert.Equal(t, "media-1", response.LookupProfile.Media.ID)
	assert.Equal(t, "JSON", response.LookupProfile.Media.Type)
	assert.Assert(t, response.ReturnedProfile == nil)
	assert.Assert(t, response.IDPhoto == nil)
	assert.Assert(t, response.File == nil)
}

func TestShareCodeResourceResponse_FullRealisticPayload(t *testing.T) {
	jsonData := `{
		"id": "abc12345-6789-abcd-ef01-234567890abc",
		"source": {"type": "END_USER"},
		"created_at": "2026-02-05T11:33:46Z",
		"last_updated": "2026-02-05T11:33:50Z",
		"lookup_profile": {
			"media": { "id": "df419a66-0449-41cf-a795-6dfaa993d1f6", "type": "JSON", "created": "2026-02-05T11:33:46Z", "last_updated": "2026-02-05T11:33:50Z" }
		},
		"returned_profile": {
			"media": { "id": "f2152059-2868-47c9-8f5f-64966c1b66b0", "type": "JSON", "created": "2026-02-05T11:33:46Z", "last_updated": "2026-02-05T11:33:50Z" }
		},
		"id_photo": {
			"media": { "id": "45e4ee9d-a77b-4007-afe9-ab7067687aff", "type": "IMAGE", "created": "2026-02-05T11:33:46Z", "last_updated": "2026-02-05T11:33:50Z" }
		},
		"file": {
			"media": { "id": "c83a9f12-1234-5678-9abc-def012345678", "type": "PDF", "created": "2026-02-05T11:33:46Z", "last_updated": "2026-02-05T11:33:50Z" }
		},
		"tasks": [
			{
				"type": "VERIFY_SHARE_CODE_TASK",
				"id": "73141aa9-a01f-4de9-9281-1b11cda7ab75",
				"state": "DONE",
				"created": "2026-02-05T11:33:46Z",
				"last_updated": "2026-02-05T11:33:50Z",
				"generated_media": [
					{ "id": "df419a66-0449-41cf-a795-6dfaa993d1f6", "type": "PDF" },
					{ "id": "45e4ee9d-a77b-4007-afe9-ab7067687aff", "type": "IMAGE" },
					{ "id": "f2152059-2868-47c9-8f5f-64966c1b66b0", "type": "JSON" }
				]
			}
		]
	}`

	var response ShareCodeResourceResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NilError(t, err)
	assert.Equal(t, "abc12345-6789-abcd-ef01-234567890abc", response.ID)
	assert.Equal(t, "END_USER", response.Source)

	assert.Assert(t, response.LookupProfile != nil)
	assert.Assert(t, response.LookupProfile.Media != nil)
	assert.Equal(t, "df419a66-0449-41cf-a795-6dfaa993d1f6", response.LookupProfile.Media.ID)
	assert.Equal(t, "JSON", response.LookupProfile.Media.Type)
	assert.Assert(t, response.LookupProfile.Media.Created != nil)
	assert.Assert(t, response.LookupProfile.Media.LastUpdated != nil)

	assert.Assert(t, response.ReturnedProfile != nil)
	assert.Assert(t, response.ReturnedProfile.Media != nil)
	assert.Equal(t, "f2152059-2868-47c9-8f5f-64966c1b66b0", response.ReturnedProfile.Media.ID)

	assert.Assert(t, response.IDPhoto != nil)
	assert.Assert(t, response.IDPhoto.Media != nil)
	assert.Equal(t, "45e4ee9d-a77b-4007-afe9-ab7067687aff", response.IDPhoto.Media.ID)
	assert.Equal(t, "IMAGE", response.IDPhoto.Media.Type)

	assert.Assert(t, response.File != nil)
	assert.Assert(t, response.File.Media != nil)
	assert.Equal(t, "c83a9f12-1234-5678-9abc-def012345678", response.File.Media.ID)
	assert.Equal(t, "PDF", response.File.Media.Type)

	assert.Equal(t, 1, len(response.VerifyShareCodeTasks()))
	assert.Equal(t, "DONE", response.VerifyShareCodeTasks()[0].State)
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
