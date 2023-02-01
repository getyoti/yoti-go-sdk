package retrieve

// ImportTokenResponse contains info about the media needed to retrieve the ImportToken.
type ImportTokenResponse struct {
	FailureReasonResponse FailureReasonResponse `json:"failure_reason"`
	Media                 *MediaResponse        `json:"media"`
}
