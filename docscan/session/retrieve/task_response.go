package retrieve

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// TaskResponse represents the attributes of a task, for any given session
type TaskResponse struct {
	ID                      string                    `json:"id"`
	Type                    string                    `json:"type"`
	State                   string                    `json:"state"`
	Created                 *time.Time                `json:"created"`
	LastUpdated             *time.Time                `json:"last_updated"`
	GeneratedChecks         []*GeneratedCheckResponse `json:"generated_checks"`
	GeneratedMedia          []*GeneratedMedia         `json:"generated_media"`
	generatedTextDataChecks []*GeneratedTextDataCheckResponse
}

// GeneratedTextDataChecks  filters the checks, returning only text data checks
func (t *TaskResponse) GeneratedTextDataChecks() []*GeneratedTextDataCheckResponse {
	return t.generatedTextDataChecks
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (t *TaskResponse) UnmarshalJSON(data []byte) error {
	type result TaskResponse // declared as "type" to prevent recursive unmarshalling
	if err := json.Unmarshal(data, (*result)(t)); err != nil {
		return err
	}

	for _, check := range t.GeneratedChecks {
		switch check.Type {
		case constants.IDDocumentTextDataCheck:
			t.generatedTextDataChecks = append(t.generatedTextDataChecks, &GeneratedTextDataCheckResponse{GeneratedCheckResponse: check})

		default:
			fmt.Printf("Unrecognized check type: `%s`", check.Type)
		}
	}

	return nil
}
