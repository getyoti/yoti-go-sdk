package retrieve

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// ShareCodeMediaResponse wraps a MediaResponse inside a share code resource field.
// The API returns media objects nested under a "media" key, e.g.
//
//	{"id_photo": {"media": {"id": "...", "type": "IMAGE", ...}}}
type ShareCodeMediaResponse struct {
	Media *MediaResponse `json:"media"`
}

// ShareCodeResourceResponse represents a share code resource for a given session
type ShareCodeResourceResponse struct {
	*ResourceResponse
	Source               string                  `json:"source"`
	CreatedAt            string                  `json:"created_at"`
	LastUpdated          string                  `json:"last_updated"`
	LookupProfile        *ShareCodeMediaResponse `json:"lookup_profile,omitempty"`
	ReturnedProfile      *ShareCodeMediaResponse `json:"returned_profile,omitempty"`
	IDPhoto              *ShareCodeMediaResponse `json:"id_photo,omitempty"`
	File                 *ShareCodeMediaResponse `json:"file,omitempty"`
	verifyShareCodeTasks []*VerifyShareCodeTaskResponse
}

// VerifyShareCodeTasks returns a slice of verify share code tasks associated with the share code
func (s *ShareCodeResourceResponse) VerifyShareCodeTasks() []*VerifyShareCodeTaskResponse {
	return s.verifyShareCodeTasks
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (s *ShareCodeResourceResponse) UnmarshalJSON(data []byte) error {
	type sourceObject struct {
		Type string `json:"type"`
	}
	type wire struct {
		ID              string                  `json:"id"`
		Source          json.RawMessage         `json:"source"`
		CreatedAt       string                  `json:"created_at"`
		LastUpdated     string                  `json:"last_updated"`
		LookupProfile   *ShareCodeMediaResponse `json:"lookup_profile,omitempty"`
		ReturnedProfile *ShareCodeMediaResponse `json:"returned_profile,omitempty"`
		IDPhoto         *ShareCodeMediaResponse `json:"id_photo,omitempty"`
		File            *ShareCodeMediaResponse `json:"file,omitempty"`
		Tasks           []*TaskResponse         `json:"tasks"`
	}

	var w wire
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}

	s.ResourceResponse = &ResourceResponse{ID: w.ID, Tasks: w.Tasks}
	s.CreatedAt = w.CreatedAt
	s.LastUpdated = w.LastUpdated
	s.LookupProfile = w.LookupProfile
	s.ReturnedProfile = w.ReturnedProfile
	s.IDPhoto = w.IDPhoto
	s.File = w.File

	// API may return source as either a string (legacy) or an object like {"type":"END_USER"}.
	if len(w.Source) != 0 && string(w.Source) != "null" {
		var sourceString string
		if err := json.Unmarshal(w.Source, &sourceString); err == nil {
			s.Source = sourceString
		} else {
			var sourceObj sourceObject
			if err := json.Unmarshal(w.Source, &sourceObj); err == nil {
				if sourceObj.Type != "" {
					s.Source = sourceObj.Type
				} else {
					s.Source = string(w.Source)
				}
			} else {
				// Be lenient to avoid breaking when API adds new shapes.
				s.Source = string(w.Source)
			}
		}
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
