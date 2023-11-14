package retrieve

// AdvancedIdentityProfileResponse contains the SubjectId, the Result/FailureReasonResponse, verified identity details,
// and the verification reports that certifies how the identity was verified and how the verification levels were achieved.
type AdvancedIdentityProfileResponse struct {
	SubjectId             string                 `json:"subject_id"`
	Result                string                 `json:"result"`
	FailureReasonResponse FailureReasonResponse  `json:"failure_reason"`
	Report                map[string]interface{} `json:"identity_profile_report"`
}
