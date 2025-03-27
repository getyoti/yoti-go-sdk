package retrieve

type FailureReasonResponse struct {
	ReasonCode                string                     `json:"reason_code"`
	RequirementsNotMetDetails []RequirementsNotMetDetail `json:"requirements_not_met_details"`
}

type RequirementsNotMetDetail struct {
	FailureType            string `json:"failure_type"`
	DocumentType           string `json:"document_type"`
	DocumentCountryIsoCode string `json:"document_country_iso_code"`
	AuditId                string `json:"audit_id"`
	Details                string `json:"details"`
}
