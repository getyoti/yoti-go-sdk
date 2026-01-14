package filter

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRequiredShareCode_ShouldBuildWithIssuerAndScheme(t *testing.T) {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithIssuer("test-issuer").
		WithScheme("test-scheme").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, "test-issuer", shareCode.Issuer)
	assert.Equal(t, "test-scheme", shareCode.Scheme)
}

func TestRequiredShareCode_ShouldBuildWithOnlyIssuer(t *testing.T) {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithIssuer("test-issuer").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, "test-issuer", shareCode.Issuer)
	assert.Equal(t, "", shareCode.Scheme)
}

func TestRequiredShareCode_ShouldBuildWithOnlyScheme(t *testing.T) {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithScheme("test-scheme").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, "", shareCode.Issuer)
	assert.Equal(t, "test-scheme", shareCode.Scheme)
}

func TestRequiredShareCode_ShouldMarshalCorrectly(t *testing.T) {
	shareCode, _ := NewRequiredShareCodeBuilder().
		WithIssuer("test-issuer").
		WithScheme("test-scheme").
		Build()

	data, err := json.Marshal(shareCode)
	assert.NilError(t, err)

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	assert.Equal(t, "test-issuer", result["issuer"])
	assert.Equal(t, "test-scheme", result["scheme"])
}

func TestRequiredShareCode_Type(t *testing.T) {
	shareCode, _ := NewRequiredShareCodeBuilder().Build()
	assert.Equal(t, "SHARE_CODE", shareCode.Type())
}

func ExampleNewRequiredShareCodeBuilder() {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithIssuer("yoti").
		WithScheme("DBS").
		Build()

	if err != nil {
		// handle error
	}

	_ = shareCode
	// Output:
}

func ExampleRequiredShareCodeBuilder_Build() {
	shareCode, err := NewRequiredShareCodeBuilder().
		WithIssuer("yoti").
		WithScheme("DBS").
		Build()

	if err != nil {
		// handle error
	}

	_ = shareCode
	// Output:
}
