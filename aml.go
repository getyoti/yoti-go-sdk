package yoti

import (
	"encoding/json"
)

//AmlAddress Address for Anti Money Laundering (AML) purposes
type AmlAddress struct {
	Country  string `json:"country"`
	Postcode string `json:"post_code"`
}

//AmlProfile User profile for Anti Money Laundering (AML) checks
type AmlProfile struct {
	GivenNames string     `json:"given_names"`
	FamilyName string     `json:"family_name"`
	Address    AmlAddress `json:"address"`
	SSN        string     `json:"ssn"`
}

//AmlResult Result of Anti Money Laundering (AML) check for a particular user
type AmlResult struct {
	OnFraudList bool `json:"on_fraud_list"`
	OnPEPList   bool `json:"on_pep_list"`
	OnWatchList bool `json:"on_watch_list"`
}

//GetAmlResultFromResponse Parses AML result from response
func GetAmlResultFromResponse(amlResponse []byte) AmlResult {
	var amlResult AmlResult
	json.Unmarshal(amlResponse, &amlResult)

	return amlResult
}
