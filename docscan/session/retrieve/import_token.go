package retrieve

// ImportTokenResponse contains info about the media needed to retrieve the ImportToken.
type ImportTokenResponse struct {
	FailureReason string         `json:"failure_reason"`
	Media         *MediaResponse `json:"media"`
}
