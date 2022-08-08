package retrieve

// IdentityProfileResponse contains the SubjectId, teh Result/FailureReasonResponse, verified identity details,
// and the verification report that certifies how the identity was verified and how the verification level was achieved.
type IdentityProfileResponse struct {
	SubjectId             string                 `json:"subject_id"`
	Result                string                 `json:"result"`
	FailureReasonResponse FailureReasonResponse  `json:"failure_reason"`
	Report                map[string]interface{} `json:"identity_profile_report"`
}
