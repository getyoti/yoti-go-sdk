package retrieve

import (
	"encoding/json"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// GetSessionResult contains the information about a created session
type GetSessionResult struct {
	ClientSessionTokenTTL               int                      `json:"client_session_token_ttl"`
	ClientSessionToken                  string                   `json:"client_session_token"`
	SessionID                           string                   `json:"session_id"`
	UserTrackingID                      string                   `json:"user_tracking_id"`
	State                               string                   `json:"state"`
	Checks                              []*CheckResponse         `json:"checks"`
	Resources                           *ResourceContainer       `json:"resources"`
	BiometricConsentTimestamp           *time.Time               `json:"biometric_consent"`
	IdentityProfileResponse             *IdentityProfileResponse `json:"identity_profile"`
	IdentityProfilePreview              *IdentityProfilePreview  `json:"identity_profile_preview"`
	ImportTokenResponse                 *ImportTokenResponse     `json:"import_token"`
	authenticityChecks                  []*AuthenticityCheckResponse
	faceMatchChecks                     []*FaceMatchCheckResponse
	textDataChecks                      []*TextDataCheckResponse
	livenessChecks                      []*LivenessCheckResponse
	thirdPartyIdentityChecks            []*ThirdPartyIdentityCheckResponse
	idDocumentComparisonChecks          []*IDDocumentComparisonCheckResponse
	supplementaryDocumentTextDataChecks []*SupplementaryDocumentTextDataCheckResponse
	watchlistScreeningChecks            []*WatchlistScreeningCheckResponse
	watchlistAdvancedCAChecks           []*WatchlistAdvancedCACheckResponse
}

// AuthenticityChecks filters the checks, returning only document authenticity checks
func (g *GetSessionResult) AuthenticityChecks() []*AuthenticityCheckResponse {
	return g.authenticityChecks
}

// FaceMatchChecks filters the checks, returning only FaceMatch checks
func (g *GetSessionResult) FaceMatchChecks() []*FaceMatchCheckResponse {
	return g.faceMatchChecks
}

// TextDataChecks filters the checks, returning only ID Document Text Data checks
// Deprecated: replaced by IDDocumentTextDataChecks()
func (g *GetSessionResult) TextDataChecks() []*TextDataCheckResponse {
	return g.IDDocumentTextDataChecks()
}

// ThirdPartyIdentityChecks filters the checks, returning only external credit reference agency checks
func (g *GetSessionResult) ThirdPartyIdentityChecks() []*ThirdPartyIdentityCheckResponse {
	return g.thirdPartyIdentityChecks
}

// IDDocumentTextDataChecks filters the checks, returning only ID Document Text Data checks
func (g *GetSessionResult) IDDocumentTextDataChecks() []*TextDataCheckResponse {
	return g.textDataChecks
}

// LivenessChecks filters the checks, returning only Liveness checks
func (g *GetSessionResult) LivenessChecks() []*LivenessCheckResponse {
	return g.livenessChecks
}

// IDDocumentComparisonChecks filters the checks, returning only the identity document comparison checks
func (g *GetSessionResult) IDDocumentComparisonChecks() []*IDDocumentComparisonCheckResponse {
	return g.idDocumentComparisonChecks
}

// SupplementaryDocumentTextDataChecks filters the checks, returning only the supplementary document text data checks
func (g *GetSessionResult) SupplementaryDocumentTextDataChecks() []*SupplementaryDocumentTextDataCheckResponse {
	return g.supplementaryDocumentTextDataChecks
}

// WatchlistScreeningChecks filters the checks, returning only the Watchlist Screening checks
func (g *GetSessionResult) WatchlistScreeningChecks() []*WatchlistScreeningCheckResponse {
	return g.watchlistScreeningChecks
}

// WatchlistAdvancedCAChecks filters the checks, returning only the Watchlist Advanced CA Screening checks
func (g *GetSessionResult) WatchlistAdvancedCAChecks() []*WatchlistAdvancedCACheckResponse {
	return g.watchlistAdvancedCAChecks
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

		case constants.IDDocumentComparison:
			g.idDocumentComparisonChecks = append(g.idDocumentComparisonChecks, &IDDocumentComparisonCheckResponse{CheckResponse: check})

		case constants.ThirdPartyIdentityCheck:
			g.thirdPartyIdentityChecks = append(
				g.thirdPartyIdentityChecks,
				&ThirdPartyIdentityCheckResponse{
					CheckResponse: check,
				})

		case constants.SupplementaryDocumentTextDataCheck:
			g.supplementaryDocumentTextDataChecks = append(
				g.supplementaryDocumentTextDataChecks,
				&SupplementaryDocumentTextDataCheckResponse{
					CheckResponse: check,
				},
			)

		case constants.WatchlistScreening:
			g.watchlistScreeningChecks = append(
				g.watchlistScreeningChecks,
				&WatchlistScreeningCheckResponse{
					CheckResponse: check,
				},
			)

		case constants.WatchlistAdvancedCA:
			g.watchlistAdvancedCAChecks = append(
				g.watchlistAdvancedCAChecks,
				&WatchlistAdvancedCACheckResponse{
					CheckResponse: check,
				},
			)
		}
	}

	return nil
}
