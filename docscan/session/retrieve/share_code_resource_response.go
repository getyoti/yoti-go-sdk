package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ShareCodeResourceResponse represents a share code resource for a given session
type ShareCodeResourceResponse struct {
	*ResourceResponse
	Source               string         `json:"source"`
	CreatedAt            string         `json:"created_at"`
	LastUpdated          string         `json:"last_updated"`
	LookupProfile        *MediaResponse `json:"lookup_profile,omitempty"`
	ReturnedProfile      *MediaResponse `json:"returned_profile,omitempty"`
	IDPhoto              *MediaResponse `json:"id_photo,omitempty"`
	File                 *MediaResponse `json:"file,omitempty"`
	verifyShareCodeTasks []*VerifyShareCodeTaskResponse
}

// VerifyShareCodeTasks returns a slice of verify share code tasks associated with the share code
func (s *ShareCodeResourceResponse) VerifyShareCodeTasks() []*VerifyShareCodeTaskResponse {
	return s.verifyShareCodeTasks
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (s *ShareCodeResourceResponse) UnmarshalJSON(data []byte) error {
	type result ShareCodeResourceResponse // declared as "type" to prevent recursive unmarshalling
	if err := json.Unmarshal(data, (*result)(s)); err != nil {
		return err
	}

	for _, task := range s.Tasks {
		switch task.Type {
		case constants.VerifyShareCodeTask:
			s.verifyShareCodeTasks = append(
				s.verifyShareCodeTasks,
				&VerifyShareCodeTaskResponse{
					TaskResponse: task,
				},
			)
		}
	}

	return nil
}
