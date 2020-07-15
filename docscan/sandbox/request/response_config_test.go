package request

import (
	"encoding/json"
	"fmt"
)

func Example_responseConfigBuilder_WithCheckReports() {
	checkReports, err := NewCheckReportsBuilder().Build()
	if err != nil {
		return
	}

	responseConfig, err := NewResponseConfigBuilder().
		WithCheckReports(checkReports).
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(responseConfig)
	fmt.Println(string(data))
	// Output: {"check_reports":{"ID_DOCUMENT_AUTHENTICITY":[],"ID_DOCUMENT_TEXT_DATA_CHECK":[],"ID_DOCUMENT_FACE_MATCH_CHECK":[],"LIVENESS":[]}}
}

func Example_responseConfigBuilder_WithCheckReports_WithTaskResults() {
	taskResults, err := NewTaskResultsBuilder().Build()
	if err != nil {
		return
	}

	checkReports, err := NewCheckReportsBuilder().Build()
	if err != nil {
		return
	}

	responseConfig, err := NewResponseConfigBuilder().
		WithTaskResults(taskResults).
		WithCheckReports(checkReports).
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(responseConfig)
	fmt.Println(string(data))
	// Output: {"task_results":{"ID_DOCUMENT_TEXT_DATA_EXTRACTION":[]},"check_reports":{"ID_DOCUMENT_AUTHENTICITY":[],"ID_DOCUMENT_TEXT_DATA_CHECK":[],"ID_DOCUMENT_FACE_MATCH_CHECK":[],"LIVENESS":[]}}
}
