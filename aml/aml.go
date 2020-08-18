package aml

import (
	"encoding/json"
)

// Address Address for Anti Money Laundering (AML) purposes
type Address struct {
	Country  string `json:"country"`
	Postcode string `json:"post_code"`
}

// Profile User profile for Anti Money Laundering (AML) checks
type Profile struct {
	GivenNames string  `json:"given_names"`
	FamilyName string  `json:"family_name"`
	Address    Address `json:"address"`
	SSN        string  `json:"ssn"`
}

// Result Result of Anti Money Laundering (AML) check for a particular user
type Result struct {
	OnFraudList bool `json:"on_fraud_list"`
	OnPEPList   bool `json:"on_pep_list"`
	OnWatchList bool `json:"on_watch_list"`
}

// GetResult Parses AML result from response
func GetResult(response []byte) (Result, error) {
	var result Result
	err := json.Unmarshal(response, &result)

	if err != nil {
		return result, err
	}

	return result, err
}
