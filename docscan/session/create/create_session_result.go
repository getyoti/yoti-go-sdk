package create

// SessionResult contains the information about a created session
type SessionResult struct {
	ClientSessionTokenTTL int    `json:"client_session_token_ttl"`
	ClientSessionToken    string `json:"client_session_token"`
	SessionID             string `json:"session_id"`
}
