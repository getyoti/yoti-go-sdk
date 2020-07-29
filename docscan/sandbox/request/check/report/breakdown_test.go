package report

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func Test_BreakdownBuilder(t *testing.T) {
	breakdown, err := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").
		WithResult("some_result").
		WithDetail("some_name", "some_value").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, breakdown.SubCheck, "some_sub_check")
	assert.Equal(t, breakdown.Result, "some_result")
	assert.Equal(t, breakdown.Details[0].Name, "some_name")
	assert.Equal(t, breakdown.Details[0].Value, "some_value")
}

func Test_BreakdownBuilder_ShouldRequireSubCheck(t *testing.T) {
	_, err := NewBreakdownBuilder().
		WithResult("some_result").
		Build()

	assert.Error(t, err, "Sub Check cannot be empty")
}

func Test_BreakdownBuilder_ShouldRequireResult(t *testing.T) {
	_, err := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").
		Build()

	assert.Error(t, err, "Result cannot be empty")
}

func ExampleBreakdownBuilder() {
	breakdown, err := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").
		WithResult("some_result").
		WithDetail("some_name", "some_value").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(breakdown)
	fmt.Println(string(data))
	// Output: {"sub_check":"some_sub_check","result":"some_result","details":[{"name":"some_name","value":"some_value"}]}
}

func ExampleBreakdownBuilder_minimal() {
	breakdown, err := NewBreakdownBuilder().
		WithSubCheck("some_sub_check").
		WithResult("some_result").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(breakdown)
	fmt.Println(string(data))
	// Output: {"sub_check":"some_sub_check","result":"some_result","details":[]}
}
