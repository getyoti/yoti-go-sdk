package retrieve

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// GetSessionResult contains the information about a created session
type GetSessionResult struct {
	ClientSessionTokenTTL int                `json:"client_session_token_ttl"`
	ClientSessionToken    string             `json:"client_session_token"`
	SessionID             string             `json:"session_id"`
	UserTrackingID        string             `json:"user_tracking_id"`
	State                 string             `json:"state"`
	Checks                []*CheckResponse   `json:"checks"`
	Resources             *ResourceContainer `json:"resources"`
	authenticityChecks    []*AuthenticityCheckResponse
	faceMatchChecks       []*FaceMatchCheckResponse
	textDataChecks        []*TextDataCheckResponse
	livenessChecks        []*LivenessCheckResponse
}

// AuthenticityChecks filters the checks, returning only document authenticity checks
func (g *GetSessionResult) AuthenticityChecks() []*AuthenticityCheckResponse {
	return g.authenticityChecks
}

// FaceMatchChecks filters the checks, returning only FaceMatch checks
func (g *GetSessionResult) FaceMatchChecks() []*FaceMatchCheckResponse {
	return g.faceMatchChecks
}

// TextDataChecks filters the checks, returning only Text Data checks
func (g *GetSessionResult) TextDataChecks() []*TextDataCheckResponse {
	return g.textDataChecks
}

// LivenessChecks filters the checks, returning only Liveness checks
func (g *GetSessionResult) LivenessChecks() []*LivenessCheckResponse {
	return g.livenessChecks
}

// UnmarshalJSON handles the custom JSON unmarshalling
func (g *GetSessionResult) UnmarshalJSON(data []byte) error {
	type result GetSessionResult // declared as "type" to prevent recursive unmarshalling
	if err := json.Unmarshal(data, (*result)(g)); err != nil {
		return err
	}

	for _, check := range g.Checks {
		switch check.Type {
		case constants.IDDocumentAuthenticity:
			g.authenticityChecks = append(g.authenticityChecks, &AuthenticityCheckResponse{CheckResponse: check})

		case constants.IDDocumentFaceMatch:
			g.faceMatchChecks = append(g.faceMatchChecks, &FaceMatchCheckResponse{CheckResponse: check})

		case constants.IDDocumentTextDataCheck:
			g.textDataChecks = append(g.textDataChecks, &TextDataCheckResponse{CheckResponse: check})

		case constants.Liveness:
			g.livenessChecks = append(g.livenessChecks, &LivenessCheckResponse{CheckResponse: check})

		default:
			fmt.Printf("Unrecognized check type: `%s`", check.Type)
		}
	}

	return nil
}
