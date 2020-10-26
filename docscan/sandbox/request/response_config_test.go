package request

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestResponseConfigBuilder_Build_ShouldRequireCheckReports(t *testing.T) {
	_, err := NewResponseConfigBuilder().Build()

	assert.Error(t, err, "Check Reports must be provided")
}

func ExampleResponseConfigBuilder() {
	taskResults, err := NewTaskResultsBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	checkReports, err := NewCheckReportsBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	responseConfig, err := NewResponseConfigBuilder().
		WithTaskResults(taskResults).
		WithCheckReports(checkReports).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(responseConfig)
	fmt.Println(string(data))
	// Output: {"task_results":{"ID_DOCUMENT_TEXT_DATA_EXTRACTION":[],"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_EXTRACTION":[]},"check_reports":{"ID_DOCUMENT_AUTHENTICITY":[],"ID_DOCUMENT_TEXT_DATA_CHECK":[],"ID_DOCUMENT_FACE_MATCH":[],"LIVENESS":[],"ID_DOCUMENT_COMPARISON":[],"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_CHECK":[]}}
}

func ExampleResponseConfigBuilder_minimal() {
	checkReports, err := NewCheckReportsBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	responseConfig, err := NewResponseConfigBuilder().
		WithCheckReports(checkReports).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(responseConfig)
	fmt.Println(string(data))
	// Output: {"check_reports":{"ID_DOCUMENT_AUTHENTICITY":[],"ID_DOCUMENT_TEXT_DATA_CHECK":[],"ID_DOCUMENT_FACE_MATCH":[],"LIVENESS":[],"ID_DOCUMENT_COMPARISON":[],"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_CHECK":[]}}
}
