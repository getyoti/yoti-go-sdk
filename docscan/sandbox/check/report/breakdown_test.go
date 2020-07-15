package report

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_breakdownBuilder_WithSubCheck(t *testing.T) {
	breakdown, err := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").Build()

	assert.NilError(t, err)
	assert.Equal(t, breakdown.SubCheck, "some_sub_check")
}

func Test_breakdownBuilder_WithResult(t *testing.T) {
	breakdown, err := NewBreakdownBuilder().
		WithResult("some_result").Build()

	assert.NilError(t, err)
	assert.Equal(t, breakdown.Result, "some_result")
}

func Test_breakdownBuilder_WithDetail(t *testing.T) {
	breakdown, err := NewBreakdownBuilder().
		WithDetail("some_name", "some_value").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, breakdown.Details[0].Name, "some_name")
	assert.Equal(t, breakdown.Details[0].Value, "some_value")
}

func Example_breakdownBuilder() {
	breakdown, err := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").
		WithResult("some_result").
		WithDetail("some_name", "some_value").
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(breakdown)
	fmt.Println(string(data))
	// Output: {"sub_check":"some_sub_check","result":"some_result","details":[{"name":"some_name","value":"some_value"}]}
}
