package sandbox

import (
	"encoding/json"
	"fmt"
)

func Example_responseConfigBuilder() {
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
	// Output: {"task_results":{"ID_DOCUMENT_TEXT_DATA_EXTRACTION":null},"check_reports":{"ID_DOCUMENT_AUTHENTICITY":null,"ID_DOCUMENT_TEXT_DATA_CHECK":null,"ID_DOCUMENT_FACE_MATCH_CHECK":null,"LIVENESS":null,"async_report_delay":0}}
}
