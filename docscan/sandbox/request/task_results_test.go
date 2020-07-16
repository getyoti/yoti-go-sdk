package request

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/task"
)

func ExampleNewTaskResultsBuilder() {
	textDataExtractionTask, err := task.NewDocumentTextDataExtractionTaskBuilder().
		Build()
	if err != nil {
		return
	}

	taskResults, err := NewTaskResultsBuilder().
		WithDocumentTextDataExtractionTask(textDataExtractionTask).
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(taskResults)
	fmt.Println(string(data))
	// Output: {"ID_DOCUMENT_TEXT_DATA_EXTRACTION":[{"result":{}}]}
}
