package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ShareCodeResourceResponse represents a Share Code resource for a given session
type ShareCodeResourceResponse struct {
	*ResourceResponse
	// Source is the source of the share code
	Source string `json:"source"`
	// CreatedAt is the time when the share code was created
	CreatedAt string `json:"created_at"`
	// LastUpdated is the time when the share code was last updated
	LastUpdated string `json:"last_updated"`
	// LookupProfile contains the lookup profile with media attributes
	LookupProfile *ProfileResponse `json:"lookup_profile"`
	// ReturnedProfile contains the returned profile with media attributes
	ReturnedProfile *ProfileResponse `json:"returned_profile"`
	// IDPhoto contains the ID photo with media attributes
	IDPhoto *PhotoResponse `json:"id_photo"`
	// File contains the file with media attributes
	File *FileResponse `json:"file"`
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
			s.verifyShareCodeTasks = append(s.verifyShareCodeTasks, &VerifyShareCodeTaskResponse{TaskResponse: task})
		}
	}

	return nil
}

// ProfileResponse represents a profile with media attributes
type ProfileResponse struct {
	Media *MediaResponse `json:"media"`
}

// PhotoResponse represents a photo with media attributes
type PhotoResponse struct {
	Media *MediaResponse `json:"media"`
}
