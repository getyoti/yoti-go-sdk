package retrieve

// GetSessionResult contains the information about a created session
type GetSessionResult struct {
	ClientSessionTokenTTL int               `json:"client_session_token_ttl"`
	ClientSessionToken    string            `json:"client_session_token"`
	SessionID             string            `json:"session_id"`
	UserTrackingID        string            `json:"user_tracking_id"`
	State                 string            `json:"state"`
	Checks                []CheckResponse   `json:"checks"`
	Resources             ResourceContainer `json:"resources"`
}
