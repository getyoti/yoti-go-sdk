package sandbox

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestBreakdownBuilder_WithSubCheck(t *testing.T) {
	breakdown := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").Build()

	assert.Equal(t, breakdown.SubCheck, "some_sub_check")
}

func TestBreakdownBuilder_WithResult(t *testing.T) {
	breakdown := NewBreakdownBuilder().
		WithResult("some_result").Build()

	assert.Equal(t, breakdown.Result, "some_result")
}

func TestBreakdownBuilder_WithDetail(t *testing.T) {
	breakdown := NewBreakdownBuilder().
		WithDetail("some_name", "some_value").
		Build()

	assert.Equal(t, breakdown.Details[0].Name, "some_name")
	assert.Equal(t, breakdown.Details[0].Value, "some_value")
}

func ExampleNewBreakdownBuilder() {
	breakdown := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").
		WithResult("some_result").
		WithDetail("some_name", "some_value").
		Build()

	data, _ := json.Marshal(breakdown)
	fmt.Println(string(data))
	// Output: {"sub_check":"some_sub_check","result":"some_result","details":[{"name":"some_name","value":"some_value"}]}
}
