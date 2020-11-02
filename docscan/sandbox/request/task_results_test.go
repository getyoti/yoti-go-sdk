package request

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/task"
)

func ExampleTaskResultsBuilder() {
	textDataExtractionTask, err := task.NewDocumentTextDataExtractionTaskBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	supplementaryTextDataExtractionTask, err := task.NewSupplementaryDocumentTextDataExtractionTaskBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	taskResults, err := NewTaskResultsBuilder().
		WithDocumentTextDataExtractionTask(textDataExtractionTask).
		WithSupplementaryDocumentTextDataExtractionTask(supplementaryTextDataExtractionTask).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(taskResults)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"ID_DOCUMENT_TEXT_DATA_EXTRACTION":[{"result":{}}],"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_EXTRACTION":[{"result":{}}]}
}
