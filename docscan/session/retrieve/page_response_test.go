package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPageResponse_ExtractionImageIds_WithValues(t *testing.T) {
	jsonData := []byte(`{
		"capture_method": "UPLOAD",
		"extraction_image_ids": ["uuid-1", "uuid-2"]
	}`)

	var page PageResponse
	err := json.Unmarshal(jsonData, &page)
	assert.NilError(t, err)

	assert.Equal(t, 2, len(page.ExtractionImageIds))
	assert.Equal(t, "uuid-1", page.ExtractionImageIds[0])
	assert.Equal(t, "uuid-2", page.ExtractionImageIds[1])
}

func TestPageResponse_ExtractionImageIds_EmptyArray(t *testing.T) {
	jsonData := []byte(`{
		"capture_method": "UPLOAD",
		"extraction_image_ids": []
	}`)

	var page PageResponse
	err := json.Unmarshal(jsonData, &page)
	assert.NilError(t, err)

	assert.Equal(t, 0, len(page.ExtractionImageIds))
}

func TestPageResponse_ExtractionImageIds_NullValue(t *testing.T) {
	jsonData := []byte(`{
		"capture_method": "UPLOAD",
		"extraction_image_ids": null
	}`)

	var page PageResponse
	err := json.Unmarshal(jsonData, &page)
	assert.NilError(t, err)

	assert.Equal(t, 0, len(page.ExtractionImageIds))
}

func TestPageResponse_ExtractionImageIds_FieldAbsent(t *testing.T) {
	jsonData := []byte(`{
		"capture_method": "UPLOAD"
	}`)

	var page PageResponse
	err := json.Unmarshal(jsonData, &page)
	assert.NilError(t, err)

	assert.Equal(t, 0, len(page.ExtractionImageIds))
}

func TestPageResponse_ExtractionImageIds_RoundTrip(t *testing.T) {
	ids := []string{"uuid-aaa", "uuid-bbb"}
	page := PageResponse{
		CaptureMethod:      "CAMERA",
		ExtractionImageIds: ids,
	}

	data, err := json.Marshal(page)
	assert.NilError(t, err)

	var result PageResponse
	err = json.Unmarshal(data, &result)
	assert.NilError(t, err)

	assert.Equal(t, 2, len(result.ExtractionImageIds))
	assert.Equal(t, "uuid-aaa", result.ExtractionImageIds[0])
	assert.Equal(t, "uuid-bbb", result.ExtractionImageIds[1])
}
