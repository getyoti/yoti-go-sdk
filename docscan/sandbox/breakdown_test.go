package sandbox

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestBreakdown_WithSubCheck(t *testing.T) {
	breakdown := Breakdown{}.WithSubCheck("some_sub_check")

	assert.Equal(t, breakdown.SubCheck, "some_sub_check")
}

func TestBreakdown_WithResult(t *testing.T) {
	breakdown := Breakdown{}.WithResult("some_result")

	assert.Equal(t, breakdown.Result, "some_result")
}

func TestBreakdown_WithDetail(t *testing.T) {
	breakdown := Breakdown{}.WithDetail("some_name", "some_value")

	assert.Equal(t, breakdown.Details[0].Name, "some_name")
	assert.Equal(t, breakdown.Details[0].Value, "some_value")
}

func ExampleBreakdown() {
	breakdown := Breakdown{}.WithSubCheck("some_sub_check").WithResult("some_result").WithDetail("some_name", "some_value")

	data, _ := json.Marshal(breakdown)
	fmt.Println(string(data))
	// Output: {"sub_check":"some_sub_check","result":"some_result","details":[{"name":"some_name","value":"some_value"}]}
}
