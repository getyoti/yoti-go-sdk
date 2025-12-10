package filter

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRequiredShareCodeBuilder(t *testing.T) {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithIssuer("test-issuer").
		WithScheme("test-scheme").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, shareCode.Type(), "SHARE_CODE")
	assert.Equal(t, shareCode.Issuer, "test-issuer")
	assert.Equal(t, shareCode.Scheme, "test-scheme")
}

func TestRequiredShareCodeMarshalJSON(t *testing.T) {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithIssuer("test-issuer").
		WithScheme("test-scheme").
		Build()

	assert.NilError(t, err)

	data, err := json.Marshal(shareCode)
	assert.NilError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NilError(t, err)

	assert.Equal(t, result["type"], "SHARE_CODE")
	assert.Equal(t, result["issuer"], "test-issuer")
	assert.Equal(t, result["scheme"], "test-scheme")
}

func TestRequiredShareCodeEmptyFields(t *testing.T) {
	shareCode, err := NewRequiredShareCodeBuilder().Build()

	assert.NilError(t, err)

	data, err := json.Marshal(shareCode)
	assert.NilError(t, err)

	expected := `{"type":"SHARE_CODE"}`
	assert.Equal(t, string(data), expected)
}
